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
	FindList(ctx context.Context) ([]MapResponse, error)
	Update(ctx context.Context, payload EditMapPayload) error
	Delete(ctx context.Context, payload MapPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddMapPayload) error {
	var targetID uint
	err := r.db.GetContext(ctx, &targetID, `SELECT id FROM villages WHERE is_active = TRUE ORDER BY id DESC LIMIT 1`)
	if err != nil || targetID == 0 {
		_ = r.db.GetContext(ctx, &targetID, `SELECT id FROM villages ORDER BY id ASC LIMIT 1`)
	}

	if targetID > 0 {
		query := `
			UPDATE villages
			SET elevation = $2,
				coordinate = $3,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $1
		`
		_, updateErr := r.db.ExecContext(ctx, query, targetID, payload.Elevation, payload.Coordinate)
		return updateErr
	}

	query := `
		INSERT INTO villages (name, elevation, coordinate, is_active, created_at, updated_at)
		VALUES ('Map', $1, $2, TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, createErr := r.db.ExecContext(ctx, query, payload.Elevation, payload.Coordinate)
	return createErr
}

func (r *repository) Detail(ctx context.Context) (*MapResponse, error) {
	var result MapResponse
	query := mapSelectQuery() + `
		ORDER BY id DESC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &MapResponse{}, nil
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
			var fallbackResult MapResponse
			fallbackQuery := mapSelectQuery() + ` ORDER BY id DESC LIMIT 1 `
			if fallbackErr := r.db.GetContext(ctx, &fallbackResult, fallbackQuery); fallbackErr == nil {
				return &fallbackResult, nil
			}
			return &MapResponse{}, nil
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
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	result, err := tx.ExecContext(ctx, query,
		payload.ID, payload.Elevation, payload.Coordinate,
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

func (r *repository) FindList(ctx context.Context) ([]MapResponse, error) {
	var result []MapResponse
	query := mapSelectQuery() + `
		ORDER BY id DESC
	`
	if err := r.db.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func mapSelectQuery() string {
	return `
		SELECT id,
			COALESCE(
				NULLIF(elevation, ''),
				(SELECT elevation FROM villages WHERE is_active = TRUE AND elevation IS NOT NULL AND elevation <> '' ORDER BY id DESC LIMIT 1),
				(SELECT elevation FROM villages WHERE elevation IS NOT NULL AND elevation <> '' ORDER BY id DESC LIMIT 1),
				''
			) AS elevation,
			COALESCE(
				NULLIF(coordinate, ''),
				(SELECT coordinate FROM villages WHERE is_active = TRUE AND coordinate IS NOT NULL AND coordinate <> '' ORDER BY id DESC LIMIT 1),
				(SELECT coordinate FROM villages WHERE coordinate IS NOT NULL AND coordinate <> '' ORDER BY id DESC LIMIT 1),
				''
			) AS coordinate,
			COALESCE(
				NULLIF(hamlet_one, 0),
				(SELECT hamlet_one FROM villages WHERE is_active = TRUE AND hamlet_one IS NOT NULL AND hamlet_one > 0 ORDER BY id DESC LIMIT 1),
				(SELECT hamlet_one FROM villages WHERE hamlet_one IS NOT NULL AND hamlet_one > 0 ORDER BY id DESC LIMIT 1),
				0
			) AS hamlet_one,
			COALESCE(
				NULLIF(hamlet_two, 0),
				(SELECT hamlet_two FROM villages WHERE is_active = TRUE AND hamlet_two IS NOT NULL AND hamlet_two > 0 ORDER BY id DESC LIMIT 1),
				(SELECT hamlet_two FROM villages WHERE hamlet_two IS NOT NULL AND hamlet_two > 0 ORDER BY id DESC LIMIT 1),
				0
			) AS hamlet_two,
			COALESCE(is_active, FALSE) AS is_active
		FROM villages
	`
}

