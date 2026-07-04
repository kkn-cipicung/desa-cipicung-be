package dashboard

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrDashboardNotFound = errors.New("dashboard not found")

type Repository interface {
	Create(ctx context.Context, payload AddDashboardPayload) error
	List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error)
	FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error)
	Update(ctx context.Context, payload EditDashboardPayload) error
	Delete(ctx context.Context, payload DashboardPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddDashboardPayload) error {
	query := `
		INSERT INTO galleries (created_by, category_id, title, description)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, payload.CreatedBy, payload.CategoryID, payload.Title, payload.Description)
	return err
}

func (r *repository) List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error) {
	var results []DashboardResponse

	query := `
		SELECT id, created_by, category_id, title, description, created_at
		FROM galleries
		ORDER BY id DESC
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
		SELECT id, created_by, category_id, title, description, created_at
		FROM galleries
		WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDashboardNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditDashboardPayload) error {
	query := `
		UPDATE galleries
		SET category_id = $2,
			title = $3,
			description = $4
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Description)
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
