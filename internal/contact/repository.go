package contact

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrContactNotFound = errors.New("contact not found")

type Repository interface {
	Create(ctx context.Context, payload AddContactPayload) error
	Detail(ctx context.Context) (*ContactResponse, error)
	Update(ctx context.Context, payload EditContactPayload) error
	Delete(ctx context.Context, payload ContactPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddContactPayload) error {
	query := `
		INSERT INTO villages (
			name, province, regency, district, postal_code, address, phone, email, website,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := r.db.ExecContext(ctx, query,
		payload.Name, payload.Province, payload.Regency, payload.District, payload.PostalCode,
		payload.Address, payload.Phone, payload.Email, payload.Website,
	)
	return err
}

func (r *repository) Detail(ctx context.Context) (*ContactResponse, error) {
	var result ContactResponse
	query := `
		SELECT name, province, regency, district, postal_code, address, phone, email, website
		FROM villages
		ORDER BY id DESC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ContactResponse{}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditContactPayload) error {
	query := `
		UPDATE villages
		SET name = $2,
			province = $3,
			regency = $4,
			district = $5,
			postal_code = $6,
			address = $7,
			phone = $8,
			email = $9,
			website = $10,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query,
		payload.ID, payload.Name, payload.Province, payload.Regency, payload.District,
		payload.PostalCode, payload.Address, payload.Phone, payload.Email, payload.Website,
	)
	if err != nil {
		return err
	}

	return ensureContactAffected(result)
}

func (r *repository) Delete(ctx context.Context, payload ContactPayload) error {
	query := `
		DELETE FROM villages
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return err
	}

	return ensureContactAffected(result)
}

func ensureContactAffected(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrContactNotFound
	}
	return nil
}
