package news

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrNewsNotFound = errors.New("news not found")

type Repository interface {
	Create(ctx context.Context, payload AddNewsPayload, media *newsMediaPayload) error
	List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error)
	FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsResponse, error)
	Update(ctx context.Context, payload EditNewsPayload, media *newsMediaPayload) error
	Delete(ctx context.Context, payload NewsPayload) error
	FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsResponse, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddNewsPayload, media *newsMediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var mediaID *uint
	if media != nil {
		mediaID, err = insertNewsMedia(ctx, tx, media)
		if err != nil {
			return err
		}
	}

	documentQuery := `
		INSERT INTO documents (category_id, uploaded_by, title, description, img_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = tx.ExecContext(ctx, documentQuery, payload.CategoryID, payload.UploadedBy, payload.Title, payload.Description, mediaID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error) {
	var results []NewsResponse

	query := `
		SELECT id, category_id, uploaded_by, title, description, created_at
		FROM documents
		ORDER BY created_at DESC
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
		SELECT id, category_id, uploaded_by, title, description, created_at
		FROM documents
		WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNewsNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditNewsPayload, media *newsMediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if media != nil {
		mediaID, err := insertNewsMedia(ctx, tx, media)
		if err != nil {
			return err
		}

		query := `
			UPDATE documents
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
			return ErrNewsNotFound
		}
	} else {
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
	}

	return tx.Commit()
}

func insertNewsMedia(ctx context.Context, tx *sqlx.Tx, media *newsMediaPayload) (*uint, error) {
	query := `
		INSERT INTO media (file_path, mime_type, uploaded_by)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id uint
	if err := tx.QueryRowContext(ctx, query, media.FilePath, media.MimeType, media.UploadedBy).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
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
		SELECT id, category_id, uploaded_by, title, description, created_at
		FROM documents
		WHERE created_at::date = $1
		ORDER BY created_at DESC
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Date); err != nil {
		return nil, err
	}

	return results, nil
}
