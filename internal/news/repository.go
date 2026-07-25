package news

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"cipicung.id/be/utils"
	"github.com/jmoiron/sqlx"
)

var ErrNewsNotFound = errors.New("news not found")

type Repository interface {
	Create(ctx context.Context, payload AddNewsPayload, media *utils.MediaPayload) error
	List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error)
	FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsResponse, error)
	Update(ctx context.Context, payload EditNewsPayload, media *utils.MediaPayload) error
	Delete(ctx context.Context, payload NewsPayload) error
	FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsResponse, error)
	FindHeader(ctx context.Context, payload NewsByIdPayload) (NewsHeaderResponse, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddNewsPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin database transaction: %w", err)
	}
	defer tx.Rollback()

	documentQuery := `
		INSERT INTO documents (category_id, uploaded_by, title, description, media_id)
		VALUES ($1, $2, $3, $4, NULL)
		RETURNING id
	`
	var documentID uint
	if err := tx.QueryRowContext(ctx, documentQuery, payload.CategoryID, payload.UploadedBy, payload.Title, payload.Description).Scan(&documentID); err != nil {
		return fmt.Errorf("insert document: %w", err)
	}

	if _, err := utils.AttachMediaToEntityColumn(ctx, tx, media, "news", documentID, "image", "documents", "media_id"); err != nil {
		return fmt.Errorf("attach news media: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit database transaction: %w", err)
	}
	return nil
}

func (r *repository) List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error) {
	var results []NewsResponse

	query := `
		SELECT d.id, COALESCE(d.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(d.uploaded_by, 0) AS uploaded_by, COALESCE(u.name, '') AS uploader_name,
			COALESCE(d.title, '') AS title, COALESCE(d.description, '') AS description,
			COALESCE(m.file_path, '') AS media, COALESCE(d.created_at, NOW()) AS created_at
		FROM documents d
		LEFT JOIN users u ON d.uploaded_by = u.id
		LEFT JOIN categories c ON d.category_id = c.id
		LEFT JOIN media m ON d.media_id = m.id
		ORDER BY d.created_at DESC
		LIMIT $1 OFFSET $2
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsResponse, error) {
	var result NewsResponse

	query := `
		SELECT d.id, COALESCE(d.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(d.uploaded_by, 0) AS uploaded_by, COALESCE(u.name, '') AS uploader_name,
			COALESCE(d.title, '') AS title, COALESCE(d.description, '') AS description,
			COALESCE(m.file_path, '') AS media, COALESCE(d.created_at, NOW()) AS created_at
		FROM documents d
		LEFT JOIN users u ON d.uploaded_by = u.id
		LEFT JOIN categories c ON d.category_id = c.id
		LEFT JOIN media m ON d.media_id = m.id
		WHERE d.id = $1
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNewsNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditNewsPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		UPDATE documents
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
		return ErrNewsNotFound
	}
	_, oldMedia, err := utils.ReplaceMediaOnEntityColumn(ctx, tx, media, "news", payload.ID, "image", "documents", "media_id")
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return utils.RemoveMediaFile(oldMedia)
}

func (r *repository) Delete(ctx context.Context, payload NewsPayload) error {
	query := `
		DELETE FROM documents
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
		return ErrNewsNotFound
	}

	return nil
}

func (r *repository) FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsResponse, error) {
	var results []NewsResponse

	query := `
		SELECT d.id, COALESCE(d.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(d.uploaded_by, 0) AS uploaded_by, COALESCE(u.name, '') AS uploader_name,
			COALESCE(d.title, '') AS title, COALESCE(d.description, '') AS description,
			COALESCE(m.file_path, '') AS media, COALESCE(d.created_at, NOW()) AS created_at
		FROM documents d
		LEFT JOIN users u ON d.uploaded_by = u.id
		LEFT JOIN categories c ON d.category_id = c.id
		LEFT JOIN media m ON d.media_id = m.id
		WHERE d.created_at::date = $1
		ORDER BY d.created_at DESC
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Date); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) FindHeader(ctx context.Context, payload NewsByIdPayload) (NewsHeaderResponse, error) {
	var result NewsHeaderResponse

	query := `
		SELECT d.id, COALESCE(d.title, '') AS title
		FROM documents d
		WHERE d.id = $1
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NewsHeaderResponse{}, ErrNewsNotFound
		}
		return NewsHeaderResponse{}, err
	}

	return result, nil
}
