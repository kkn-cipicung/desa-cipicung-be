package mapdata

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrMapNotFound = errors.New("map not found")

const mapActivationLockID int64 = 2

type Repository interface {
	Create(ctx context.Context, payload AddMapPayload) error
	Detail(ctx context.Context) (*MapResponse, error)
	FindActive(ctx context.Context) (*MapResponse, error)
	Update(ctx context.Context, payload EditMapPayload) error
	Activate(ctx context.Context, payload MapPayload) error
	Delete(ctx context.Context, payload MapPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddMapPayload) error {
	query := `
		INSERT INTO villages (name, elevation, coordinate, hamlet_one, hamlet_two, population, created_at, updated_at)
		VALUES ('Map', $1, $2, $3, $4, ($3::BIGINT + $4::BIGINT)::TEXT, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := r.db.ExecContext(ctx, query, payload.Elevation, payload.Coordinate, *payload.HamletOne, *payload.HamletTwo)
	return err
}

func (r *repository) Detail(ctx context.Context) (*MapResponse, error) {
	var result MapResponse
	query := mapSelectQuery() + `
		ORDER BY id DESC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrMapNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindActive(ctx context.Context) (*MapResponse, error) {
	var result MapResponse
	query := mapSelectQuery() + `
		WHERE is_active = TRUE
		ORDER BY id DESC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrMapNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditMapPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		UPDATE villages
		SET elevation = $2,
			coordinate = $3,
			hamlet_one = $4,
			hamlet_two = $5,
			population = ($4::BIGINT + $5::BIGINT)::TEXT,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	result, err := tx.ExecContext(ctx, query,
		payload.ID, payload.Elevation, payload.Coordinate, *payload.HamletOne, *payload.HamletTwo,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrMapNotFound
	}

	return tx.Commit()
}

func (r *repository) Activate(ctx context.Context, payload MapPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := lockMapActivation(ctx, tx); err != nil {
		return err
	}

	var exists bool
	if err := tx.GetContext(ctx, &exists, `SELECT EXISTS (SELECT 1 FROM villages WHERE id = $1)`, payload.ID); err != nil {
		return err
	}
	if !exists {
		return ErrMapNotFound
	}

	if err := deactivateOtherMaps(ctx, tx, payload.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE villages SET is_active = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = $1`, payload.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) Delete(ctx context.Context, payload MapPayload) error {
	query := `
		DELETE FROM villages
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrMapNotFound
	}

	return nil
}

func mapSelectQuery() string {
	return `
		SELECT COALESCE(elevation, '') AS elevation,
			COALESCE(coordinate, '') AS coordinate,
			CAST(COALESCE(NULLIF(hamlet_one, ''), '0') AS BIGINT) AS hamlet_one,
			CAST(COALESCE(NULLIF(hamlet_two, ''), '0') AS BIGINT) AS hamlet_two
		FROM villages
	`
}

func lockMapActivation(ctx context.Context, tx *sqlx.Tx) error {
	_, err := tx.ExecContext(ctx, `SELECT pg_advisory_xact_lock($1)`, mapActivationLockID)
	return err
}

func deactivateOtherMaps(ctx context.Context, tx *sqlx.Tx, activeID uint) error {
	_, err := tx.ExecContext(ctx, `UPDATE villages SET is_active = FALSE WHERE id <> $1 AND is_active = TRUE`, activeID)
	return err
}
