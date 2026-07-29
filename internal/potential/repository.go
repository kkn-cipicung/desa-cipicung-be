package potential

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/utils"
	"github.com/jmoiron/sqlx"
)

var ErrPotentialNotFound = errors.New("potential not found")

type Repository interface {
	Create(ctx context.Context, payload AddPotentialPayload, media *utils.MediaPayload) error
	List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error)
	FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error)
	Update(ctx context.Context, payload EditPotentialPayload, media *utils.MediaPayload) error
	Delete(ctx context.Context, payload PotentialPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddPotentialPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	locationID, err := resolveLocationID(ctx, tx, payload.LocationID, payload.Location, &payload.UploadedBy)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO potentials (category_id, title, subtitle, slug, description, location_id, media_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`

	result, err := tx.ExecContext(ctx, query, payload.CategoryID, payload.Title, payload.Subtitle, payload.Slug, payload.Description, locationID)
	if err != nil {
		return err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	potentialID := uint(insertID)

	if _, err := utils.AttachMediaToEntityColumn(ctx, tx, media, "potential", potentialID, "image", "potentials", "media_id"); err != nil {
		return err
	}

	return tx.Commit()
}

func resolveLocationID(ctx context.Context, tx *sqlx.Tx, locationID *uint, location *PotentialLocationInput, createdByID *uint) (*uint, error) {
	if location != nil && location.ID != nil {
		return location.ID, nil
	}
	if location != nil {
		newLocationID, err := utils.InsertLocation(ctx, tx, &utils.LocationPayload{
			Latitude:    location.Latitude,
			Longitude:   location.Longitude,
			CreatedByID: createdByID,
			Title:       location.Title,
			Description: location.Description,
		})
		if err != nil {
			return nil, err
		}
		locationID = newLocationID
	}

	return locationID, nil
}

func (r *repository) List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error) {
	var results []PotentialResponse

	query := `
		SELECT p.id, COALESCE(p.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(p.title, '') AS title, COALESCE(p.subtitle, '') AS subtitle,
			COALESCE(p.slug, '') AS slug, COALESCE(p.description, '') AS description,
			p.location_id, COALESCE(m.file_path, '') AS media, p.created_at
		FROM potentials p
		LEFT JOIN categories c ON p.category_id = c.id
		LEFT JOIN media m ON p.media_id = m.id
		ORDER BY p.id DESC
		LIMIT ? OFFSET ?
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error) {
	var result PotentialResponse

	query := `
		SELECT p.id, COALESCE(p.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(p.title, '') AS title, COALESCE(p.subtitle, '') AS subtitle,
			COALESCE(p.slug, '') AS slug, COALESCE(p.description, '') AS description,
			p.location_id, COALESCE(m.file_path, '') AS media, p.created_at
		FROM potentials p
		LEFT JOIN categories c ON p.category_id = c.id
		LEFT JOIN media m ON p.media_id = m.id
		WHERE p.id = ?
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPotentialNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditPotentialPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	locationID, err := resolveLocationID(ctx, tx, payload.LocationID, payload.Location, nil)
	if err != nil {
		return err
	}

	query := `
		UPDATE potentials
		SET category_id = ?,
			title = ?,
			subtitle = ?,
			slug = ?,
			description = ?,
			location_id = COALESCE(?, location_id)
		WHERE id = ?
	`
	result, err := tx.ExecContext(ctx, query, payload.CategoryID, payload.Title, payload.Subtitle, payload.Slug, payload.Description, locationID, payload.ID)
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
	_, oldMedia, err := utils.ReplaceMediaOnEntityColumn(ctx, tx, media, "potential", payload.ID, "image", "potentials", "media_id")
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return utils.RemoveMediaFile(oldMedia)
}

func (r *repository) Delete(ctx context.Context, payload PotentialPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var locationID *uint
	if err := tx.GetContext(ctx, &locationID, `SELECT location_id FROM potentials WHERE id = ?`, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrPotentialNotFound
		}
		return err
	}

	result, err := tx.ExecContext(ctx, `DELETE FROM potentials WHERE id = ?`, payload.ID)
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

	if locationID != nil {
		_, err = tx.ExecContext(ctx, `
			DELETE FROM locations
			WHERE id = ?
				AND NOT EXISTS (SELECT 1 FROM potentials WHERE location_id = ?)
				AND NOT EXISTS (SELECT 1 FROM businesses WHERE location_id = ?)
		`, *locationID, *locationID, *locationID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
