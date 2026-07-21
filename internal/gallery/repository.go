package gallery

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/utils"
	"github.com/jmoiron/sqlx"
)

var ErrGalleryNotFound = errors.New("gallery not found")
var ErrInvalidGalleryCategory = errors.New("category must be type gallery")

type Repository interface {
	Create(ctx context.Context, payload AddGalleryPayload, media *utils.MediaPayload) error
	List(ctx context.Context, payload ListGalleryPayload) ([]GalleryResponse, error)
	FindByID(ctx context.Context, payload GalleryPayload) (*GalleryResponse, error)
	Update(ctx context.Context, payload EditGalleryPayload, media *utils.MediaPayload) error
	Delete(ctx context.Context, payload GalleryPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddGalleryPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO galleries (created_by, category_id, title, description, media_id)
		SELECT $1, $2, $3, $4, NULL
		FROM categories
		WHERE id = $2 AND type = 'gallery'
		RETURNING id
	`
	var galleryID uint
	if err := tx.QueryRowContext(ctx, query, payload.CreatedBy, payload.CategoryID, payload.Title, payload.Description).Scan(&galleryID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrInvalidGalleryCategory
		}
		return err
	}

	if _, err := utils.AttachMediaToEntityColumn(ctx, tx, media, "gallery", galleryID, "image", "galleries", "media_id"); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) List(ctx context.Context, payload ListGalleryPayload) ([]GalleryResponse, error) {
	var results []GalleryResponse
	query := `
		SELECT g.id, COALESCE(g.title, '') AS title, COALESCE(m.file_path, '') AS image
		FROM galleries g
		JOIN categories c ON g.category_id = c.id AND c.type = 'gallery'
		LEFT JOIN media m ON g.media_id = m.id
		ORDER BY g.id DESC
		LIMIT $1 OFFSET $2
	`
	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload GalleryPayload) (*GalleryResponse, error) {
	var result GalleryResponse
	query := `
		SELECT g.id, COALESCE(g.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(g.title, '') AS title, COALESCE(g.description, '') AS description,
			COALESCE(m.file_path, '') AS image
		FROM galleries g
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
		WHERE g.id = $1 AND c.type = 'gallery'
	`
	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGalleryNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditGalleryPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		UPDATE galleries
		SET category_id = $2,
			title = $3,
			description = $4
		WHERE id = $1
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = galleries.category_id AND c.type = 'gallery')
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = $2 AND c.type = 'gallery')
	`
	result, err := tx.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Description)
	if err != nil {
		return err
	}
	if err := ensureGalleryAffected(result); err != nil {
		if errors.Is(err, ErrGalleryNotFound) {
			if err := ensureGalleryExists(ctx, tx, payload.ID); err != nil {
				return err
			}
			return ErrInvalidGalleryCategory
		}
		return err
	}
	_, oldMedia, err := utils.ReplaceMediaOnEntityColumn(ctx, tx, media, "gallery", payload.ID, "image", "galleries", "media_id")
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return utils.RemoveMediaFile(oldMedia)
}

func (r *repository) Delete(ctx context.Context, payload GalleryPayload) error {
	query := `
		DELETE FROM galleries
		WHERE id = $1
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = galleries.category_id AND c.type = 'gallery')
	`
	result, err := r.db.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return err
	}
	return ensureGalleryAffected(result)
}

func ensureGalleryAffected(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrGalleryNotFound
	}
	return nil
}

func ensureGalleryExists(ctx context.Context, tx *sqlx.Tx, id uint) error {
	var exists bool
	if err := tx.GetContext(ctx, &exists, `SELECT EXISTS (SELECT 1 FROM galleries WHERE id = $1)`, id); err != nil {
		return err
	}
	if !exists {
		return ErrGalleryNotFound
	}
	return nil
}
