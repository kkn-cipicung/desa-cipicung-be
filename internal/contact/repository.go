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
	FindActive(ctx context.Context) (*ContactResponse, error)
	Update(ctx context.Context, payload EditContactPayload) error
	Activate(ctx context.Context, payload ContactPayload) error
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
			ig_usn, tiktok_usn, yt_usn, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := r.db.ExecContext(ctx, query,
		payload.Name, payload.Province, payload.Regency, payload.District, payload.PostalCode,
		payload.Address, payload.Phone, payload.Email, payload.Website,
		payload.Instagram, payload.TikTok, payload.YouTube,
	)
	return err
}

func (r *repository) Detail(ctx context.Context) (*ContactResponse, error) {
	var result ContactResponse
	query := `
		SELECT id, COALESCE(name, '') AS name, COALESCE(province, '') AS province,
			COALESCE(regency, '') AS regency, COALESCE(district, '') AS district,
			COALESCE(postal_code, '') AS postal_code, COALESCE(address, '') AS address,
			COALESCE(phone, '') AS phone, COALESCE(email, '') AS email, COALESCE(website, '') AS website,
			COALESCE(ig_usn, '') AS ig_usn, COALESCE(tiktok_usn, '') AS tiktok_usn, COALESCE(yt_usn, '') AS yt_usn,
			COALESCE(is_active, FALSE) AS is_active
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

func (r *repository) FindActive(ctx context.Context) (*ContactResponse, error) {
	var result ContactResponse
	query := `
		SELECT id, COALESCE(name, '') AS name, COALESCE(province, '') AS province,
			COALESCE(regency, '') AS regency, COALESCE(district, '') AS district,
			COALESCE(postal_code, '') AS postal_code, COALESCE(address, '') AS address,
			COALESCE(phone, '') AS phone, COALESCE(email, '') AS email, COALESCE(website, '') AS website,
			COALESCE(ig_usn, '') AS ig_usn, COALESCE(tiktok_usn, '') AS tiktok_usn, COALESCE(yt_usn, '') AS yt_usn,
			COALESCE(is_active, FALSE) AS is_active
		FROM villages
		WHERE is_active = TRUE
		ORDER BY id DESC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			var fallbackResult ContactResponse
			fallbackQuery := `
				SELECT id, COALESCE(name, '') AS name, COALESCE(province, '') AS province,
					COALESCE(regency, '') AS regency, COALESCE(district, '') AS district,
					COALESCE(postal_code, '') AS postal_code, COALESCE(address, '') AS address,
					COALESCE(phone, '') AS phone, COALESCE(email, '') AS email, COALESCE(website, '') AS website,
					COALESCE(ig_usn, '') AS ig_usn, COALESCE(tiktok_usn, '') AS tiktok_usn, COALESCE(yt_usn, '') AS yt_usn,
					COALESCE(is_active, FALSE) AS is_active
				FROM villages
				ORDER BY id DESC
				LIMIT 1
			`
			if fallbackErr := r.db.GetContext(ctx, &fallbackResult, fallbackQuery); fallbackErr == nil {
				return &fallbackResult, nil
			}
			return &ContactResponse{}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Activate(ctx context.Context, payload ContactPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var exists bool
	if err := tx.GetContext(ctx, &exists, `SELECT EXISTS (SELECT 1 FROM villages WHERE id = ?)`, payload.ID); err != nil {
		return err
	}
	if !exists {
		return ErrContactNotFound
	}

	if _, err := tx.ExecContext(ctx, `UPDATE villages SET is_active = FALSE WHERE id <> ? AND is_active = TRUE`, payload.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE villages SET is_active = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, payload.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) Update(ctx context.Context, payload EditContactPayload) error {
	query := `
		UPDATE villages
		SET name = ?,
			province = ?,
			regency = ?,
			district = ?,
			postal_code = ?,
			address = ?,
			phone = ?,
			email = ?,
			website = ?,
			ig_usn = ?,
			tiktok_usn = ?,
			yt_usn = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	result, err := r.db.ExecContext(ctx, query,
		payload.Name, payload.Province, payload.Regency, payload.District,
		payload.PostalCode, payload.Address, payload.Phone, payload.Email, payload.Website,
		payload.Instagram, payload.TikTok, payload.YouTube,
		payload.ID,
	)
	if err != nil {
		return err
	}

	return ensureContactAffected(result)
}

func (r *repository) Delete(ctx context.Context, payload ContactPayload) error {
	query := `
		DELETE FROM villages
		WHERE id = ?
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
