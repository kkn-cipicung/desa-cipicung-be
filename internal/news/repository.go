package news

import (
	"context"
	"database/sql"
	"errors"

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

	documentQuery := `
		INSERT INTO documents (category_id, uploaded_by, title, description, media_id)
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
		SELECT d.id, d.category_id, COALESCE(c.name, '') AS category_name, d.uploaded_by, COALESCE(u.name, '') AS uploader_name, d.title, d.description, d.media_id, d.created_at
		FROM documents d
		LEFT JOIN users u ON d.uploaded_by = u.id
		LEFT JOIN categories c ON d.category_id = c.id
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
		SELECT d.id, d.category_id, COALESCE(c.name, '') AS category_name, d.uploaded_by, COALESCE(u.name, '') AS uploader_name, d.title, d.description, d.media_id, d.created_at
		FROM documents d
		LEFT JOIN users u ON d.uploaded_by = u.id
		LEFT JOIN categories c ON d.category_id = c.id
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

	if media != nil {
		mediaID, err := utils.InsertMedia(ctx, tx, media)
		if err != nil {
			return err
		}

		query := `
			UPDATE documents
			SET category_id = $2,
				title = $3,
				description = $4,
				media_id = $5
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
		SELECT d.id, d.category_id, COALESCE(c.name, '') AS category_name, d.uploaded_by, COALESCE(u.name, '') AS uploader_name, d.title, d.description, d.media_id, d.created_at
		FROM documents d
		LEFT JOIN users u ON d.uploaded_by = u.id
		LEFT JOIN categories c ON d.category_id = c.id
		WHERE d.created_at::date = $1
		ORDER BY d.created_at DESC
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Date); err != nil {
		return nil, err
	}

	return results, nil
}
