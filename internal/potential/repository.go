package potential

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrPotentialNotFound = errors.New("potential not found")

type Repository interface {
	Create(ctx context.Context, payload AddPotentialPayload, media *potentialMediaPayload) error
	List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error)
	FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error)
	Update(ctx context.Context, payload EditPotentialPayload, media *potentialMediaPayload) error
	Delete(ctx context.Context, payload PotentialPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddPotentialPayload, media *potentialMediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var mediaID *uint
	if media != nil {
		mediaID, err = insertPotentialMedia(ctx, tx, media)
		if err != nil {
			return err
		}
	}

	locationQuery := `
		INSERT INTO locations (latitude, longitude, creator_id, title, description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var locationID uint
	err = tx.QueryRowContext(ctx, locationQuery, payload.Latitude, payload.Longitude, payload.UploadedBy, payload.Title, payload.Description).Scan(&locationID)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO potentials (category_id, title, subtitle, slug, description, location_id, owner_name, owner_msisdn, media_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = tx.ExecContext(ctx, query, payload.CategoryID, payload.Title, payload.Subtitle, payload.Slug, payload.Description, locationID, payload.OwnerName, payload.OwnerMsisdn, mediaID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error) {
	var results []PotentialResponse

	query := `
		SELECT id, category_id, title, subtitle, slug, description, location_id, owner_name, owner_msisdn, media_id, created_at
		FROM potentials
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error) {
	var result PotentialResponse

	query := `
		SELECT id, category_id, title, subtitle, slug, description, location_id, owner_name, owner_msisdn, media_id, created_at
		FROM potentials
		WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPotentialNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditPotentialPayload, media *potentialMediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var locationID uint
	err = tx.QueryRowContext(ctx, "SELECT location_id FROM potentials WHERE id = $1", payload.ID).Scan(&locationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrPotentialNotFound
		}
		return err
	}

	updateLocationQuery := `
		UPDATE locations
		SET latitude = $2, longitude = $3, title = $4, description = $5
		WHERE id = $1
	`
	_, err = tx.ExecContext(ctx, updateLocationQuery, locationID, payload.Latitude, payload.Longitude, payload.Title, payload.Description)
	if err != nil {
		return err
	}

	if media != nil {
		mediaID, err := insertPotentialMedia(ctx, tx, media)
		if err != nil {
			return err
		}

		query := `
			UPDATE potentials
			SET category_id = $2,
				title = $3,
				subtitle = $4,
				slug = $5,
				description = $6,
				owner_name = $7,
				owner_msisdn = $8,
				media_id = $9
			WHERE id = $1
		`
		result, err := tx.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Subtitle, payload.Slug, payload.Description, payload.OwnerName, payload.OwnerMsisdn, *mediaID)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ErrPotentialNotFound
		}
	} else {
		query := `
			UPDATE potentials
			SET category_id = $2,
				title = $3,
				subtitle = $4,
				slug = $5,
				description = $6,
				owner_name = $7,
				owner_msisdn = $8
			WHERE id = $1
		`
		result, err := tx.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Subtitle, payload.Slug, payload.Description, payload.OwnerName, payload.OwnerMsisdn)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ErrPotentialNotFound
		}
	}

	return tx.Commit()
}

func insertPotentialMedia(ctx context.Context, tx *sqlx.Tx, media *potentialMediaPayload) (*uint, error) {
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

func (r *repository) Delete(ctx context.Context, payload PotentialPayload) error {
	query := `
		DELETE FROM potentials
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
		return ErrPotentialNotFound
	}

	return nil
}
