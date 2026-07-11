package dashboard

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/utils"
	"github.com/jmoiron/sqlx"
)

var ErrDashboardNotFound = errors.New("dashboard not found")

type Repository interface {
	Create(ctx context.Context, payload AddDashboardPayload, media *utils.MediaPayload) error
	List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error)
	FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error)
	Update(ctx context.Context, payload EditDashboardPayload, media *utils.MediaPayload) error
	Delete(ctx context.Context, payload DashboardPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddDashboardPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var mediaID *uint
	if media != nil {
		mediaID, err = utils.InsertMedia(ctx, tx, media)
		if err != nil {
			return err
		}
	}

	query := `
		INSERT INTO galleries (created_by, category_id, title, description, img_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.ExecContext(ctx, query, payload.CreatedBy, payload.CategoryID, payload.Title, payload.Description, mediaID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error) {
	var results []DashboardResponse

	query := `
		SELECT g.id, g.created_by, COALESCE(u.name, '') AS creator_name, g.category_id, COALESCE(c.name, '') AS category_name, g.title, g.description, g.created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		ORDER BY g.id DESC
		LIMIT $1 OFFSET $2
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error) {
	var result DashboardResponse

	query := `
		SELECT g.id, g.created_by, COALESCE(u.name, '') AS creator_name, g.category_id, COALESCE(c.name, '') AS category_name, g.title, g.description, g.created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		WHERE g.id = $1
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDashboardNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditDashboardPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if media != nil {
		mediaID, err := utils.InsertMedia(ctx, tx, media)
		if err != nil {
			return err
		}

		query := `
			UPDATE galleries
			SET category_id = $2,
				title = $3,
				description = $4,
				img_id = $5
			WHERE id = $1
		`
		result, err := tx.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Description, *mediaID)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ErrDashboardNotFound
		}
	} else {
		query := `
			UPDATE galleries
			SET category_id = $2,
				title = $3,
				description = $4
			WHERE id = $1
		`
		result, err := tx.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Description)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ErrDashboardNotFound
		}
	}

	return tx.Commit()
}

func (r *repository) Delete(ctx context.Context, payload DashboardPayload) error {
	query := `
		DELETE FROM galleries
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
		return ErrDashboardNotFound
	}

	return nil
}
