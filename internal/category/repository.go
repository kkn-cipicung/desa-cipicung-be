package category

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/pkg/models"
	"github.com/jmoiron/sqlx"
)

var ErrCategoryNotFound = errors.New("category not found")

type Repository interface {
	Create(ctx context.Context, payload AddCategoryPayload, slug string) error
	List(ctx context.Context, payload ListCategoryPayload) ([]models.Category, error)
	FindByID(ctx context.Context, payload CategoryPayload) (*models.Category, error)
	Update(ctx context.Context, payload EditCategoryPayload, slug string) error
	Delete(ctx context.Context, payload CategoryPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddCategoryPayload, slug string) error {
	query := `
		INSERT INTO categories (name, slug, type)
		VALUES (?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, payload.Name, slug, payload.Type)
	return err
}

func (r *repository) List(ctx context.Context, payload ListCategoryPayload) ([]models.Category, error) {
	var results []models.Category

	query := `
		SELECT id, COALESCE(name, '') AS name, COALESCE(slug, '') AS slug,
			COALESCE(type, '') AS type, COALESCE(created_at, NOW()) AS created_at
		FROM categories
		WHERE (? = '' OR LOWER(TRIM(type)) = ?)
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Type, payload.Type, payload.Limit, payload.Index); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload CategoryPayload) (*models.Category, error) {
	var result models.Category
	query := `
		SELECT id, COALESCE(name, '') AS name, COALESCE(slug, '') AS slug,
			COALESCE(type, '') AS type, COALESCE(created_at, NOW()) AS created_at
		FROM categories
		WHERE id = ?
	`
	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditCategoryPayload, slug string) error {
	query := `
		UPDATE categories
		SET name = ?, slug = ?, type = ?
		WHERE id = ?
	`
	result, err := r.db.ExecContext(ctx, query, payload.Name, slug, payload.Type, payload.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, payload CategoryPayload) error {
	query := `
		DELETE FROM categories
		WHERE id = ?
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
		return ErrCategoryNotFound
	}
	return nil
}
