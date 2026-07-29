package business

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/pkg/models"
	"github.com/jmoiron/sqlx"
)

var ErrBusinessNotFound = errors.New("business not found")

type Repository interface {
	Create(ctx context.Context, payload AddBusinessPayload) error
	List(ctx context.Context, payload ListBusinessPayload) ([]models.Business, error)
	FindByID(ctx context.Context, payload BusinessPayload) (*models.Business, error)
	Update(ctx context.Context, payload EditBusinessPayload) error
	Delete(ctx context.Context, payload BusinessPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddBusinessPayload) error {
	query := `
		INSERT INTO businesses (
			category_id,
			owner_name,
			business_name,
			description,
			phone,
			address,
			location_id,
			instagram,
			facebook
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query,
		payload.CategoryID,
		payload.OwnerName,
		payload.BusinessName,
		payload.Description,
		payload.Phone,
		payload.Address,
		payload.LocationID,
		payload.Instagram,
		payload.Facebook,
	)
	return err
}

func (r *repository) List(ctx context.Context, payload ListBusinessPayload) ([]models.Business, error) {
	var results []models.Business

	query := `
		SELECT b.id, COALESCE(b.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(b.owner_name, '') AS owner_name, COALESCE(b.business_name, '') AS business_name,
			COALESCE(b.description, '') AS description, COALESCE(b.phone, '') AS phone,
			COALESCE(b.address, '') AS address, b.location_id, b.instagram, b.facebook,
			COALESCE(b.created_at, NOW()) AS created_at
		FROM businesses b
		LEFT JOIN categories c ON b.category_id = c.id
		WHERE (? = '' OR c.slug = ?)
		ORDER BY b.id DESC
		LIMIT ? OFFSET ?
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Type, payload.Type, payload.Limit, payload.Index); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) FindByID(ctx context.Context, payload BusinessPayload) (*models.Business, error) {
	var result models.Business

	query := `
		SELECT b.id, COALESCE(b.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(b.owner_name, '') AS owner_name, COALESCE(b.business_name, '') AS business_name,
			COALESCE(b.description, '') AS description, COALESCE(b.phone, '') AS phone,
			COALESCE(b.address, '') AS address, b.location_id, b.instagram, b.facebook,
			COALESCE(b.created_at, NOW()) AS created_at
		FROM businesses b
		LEFT JOIN categories c ON b.category_id = c.id
		WHERE b.id = ?
	`

	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBusinessNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditBusinessPayload) error {
	query := `
		UPDATE businesses
		SET category_id = ?,
			owner_name = COALESCE(?, owner_name),
			business_name = COALESCE(?, business_name),
			description = COALESCE(?, description),
			phone = COALESCE(?, phone),
			address = COALESCE(?, address),
			location_id = COALESCE(?, location_id),
			instagram = COALESCE(?, instagram),
			facebook = COALESCE(?, facebook)
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query,
		payload.CategoryID,
		payload.OwnerName,
		payload.BusinessName,
		payload.Description,
		payload.Phone,
		payload.Address,
		payload.LocationID,
		payload.Instagram,
		payload.Facebook,
		payload.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrBusinessNotFound
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, payload BusinessPayload) error {
	query := `
		DELETE FROM businesses
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
		return ErrBusinessNotFound
	}
	return nil
}
