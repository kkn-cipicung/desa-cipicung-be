package dashboard

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/utils"
	"github.com/jmoiron/sqlx"
)

var ErrDashboardNotFound = errors.New("dashboard not found")

const dashboardActivationLockID int64 = 1

type Repository interface {
	Create(ctx context.Context, payload AddDashboardPayload, media *utils.MediaPayload) error
	List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error)
	Detail(ctx context.Context) (*DashboardResponse, error)
	FindActive(ctx context.Context) (*DashboardResponse, error)
	Update(ctx context.Context, payload EditDashboardPayload, media *utils.MediaPayload) error
	Activate(ctx context.Context, payload DashboardPayload) error
	Delete(ctx context.Context, payload DashboardPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddDashboardPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO galleries (created_by, category_id, title, description, media_id)
		SELECT $1, $2, $3, $4, NULL
		FROM categories
		WHERE id = $2 AND type = 'dashboard'
		RETURNING id
	`

	var dashboardID uint
	if err := tx.QueryRowContext(ctx, query, payload.CreatedBy, payload.CategoryID, payload.Title, payload.Description).Scan(&dashboardID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrDashboardNotFound
		}
		return err
	}

	if _, err := utils.AttachMediaToEntityColumn(ctx, tx, media, "dashboard", dashboardID, "image", "galleries", "media_id"); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error) {
	var results []DashboardResponse

	query := `
		SELECT g.id, g.created_by, COALESCE(u.name, '') AS creator_name, g.category_id, COALESCE(c.name, '') AS category_name, g.title, g.description, COALESCE(m.file_path, '') AS media, g.is_active, g.created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
		WHERE c.type = 'dashboard'
		ORDER BY g.id DESC
		LIMIT $1 OFFSET $2
	`

	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) Detail(ctx context.Context) (*DashboardResponse, error) {
	var result DashboardResponse

	query := `
		SELECT g.id, g.created_by, COALESCE(u.name, '') AS creator_name, g.category_id, COALESCE(c.name, '') AS category_name, g.title, g.description, COALESCE(m.file_path, '') AS media, g.is_active, g.created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
		WHERE c.type = 'dashboard'
		ORDER BY g.id DESC
		LIMIT 1
	`

	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &DashboardResponse{}, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) FindActive(ctx context.Context) (*DashboardResponse, error) {
	var result DashboardResponse

	query := `
		SELECT g.id, g.created_by, COALESCE(u.name, '') AS creator_name, g.category_id, COALESCE(c.name, '') AS category_name, g.title, g.description, COALESCE(m.file_path, '') AS media, g.is_active, g.created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
		WHERE g.is_active = TRUE AND c.type = 'dashboard'
		ORDER BY g.id DESC
		LIMIT 1
	`

	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			var fallbackResult DashboardResponse
			fallbackQuery := `
				SELECT g.id, g.created_by, COALESCE(u.name, '') AS creator_name, g.category_id, COALESCE(c.name, '') AS category_name, g.title, g.description, COALESCE(m.file_path, '') AS media, g.is_active, g.created_at
				FROM galleries g
				LEFT JOIN users u ON g.created_by = u.id
				LEFT JOIN categories c ON g.category_id = c.id
				LEFT JOIN media m ON g.media_id = m.id
				WHERE c.type = 'dashboard'
				ORDER BY g.id DESC
				LIMIT 1
			`
			if fallbackErr := r.db.GetContext(ctx, &fallbackResult, fallbackQuery); fallbackErr == nil {
				return &fallbackResult, nil
			}
			return &DashboardResponse{}, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, payload EditDashboardPayload, media *utils.MediaPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if payload.IsActive {
		if err := lockDashboardActivation(ctx, tx); err != nil {
			return err
		}
		if err := deactivateOtherDashboards(ctx, tx, payload.ID); err != nil {
			return err
		}
	}

	query := `
		UPDATE galleries
		SET category_id = $2,
			title = $3,
			description = $4,
			is_active = $5
		WHERE id = $1
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = galleries.category_id AND c.type = 'dashboard')
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = $2 AND c.type = 'dashboard')
	`
	result, err := tx.ExecContext(ctx, query, payload.ID, payload.CategoryID, payload.Title, payload.Description, payload.IsActive)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDashboardNotFound
	}
	_, oldMedia, err := utils.ReplaceMediaOnEntityColumn(ctx, tx, media, "dashboard", payload.ID, "image", "galleries", "media_id")
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return utils.RemoveMediaFile(oldMedia)
}

func (r *repository) Activate(ctx context.Context, payload DashboardPayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := lockDashboardActivation(ctx, tx); err != nil {
		return err
	}

	var exists bool
	if err := tx.GetContext(ctx, &exists, `
		SELECT EXISTS (
			SELECT 1 FROM galleries g
			JOIN categories c ON c.id = g.category_id
			WHERE g.id = $1 AND c.type = 'dashboard'
		)
	`, payload.ID); err != nil {
		return err
	}
	if !exists {
		return ErrDashboardNotFound
	}

	if err := deactivateOtherDashboards(ctx, tx, payload.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE galleries SET is_active = TRUE WHERE id = $1`, payload.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func lockDashboardActivation(ctx context.Context, tx *sqlx.Tx) error {
	_, err := tx.ExecContext(ctx, `SELECT pg_advisory_xact_lock($1)`, dashboardActivationLockID)
	return err
}

func deactivateOtherDashboards(ctx context.Context, tx *sqlx.Tx, activeID uint) error {
	_, err := tx.ExecContext(ctx, `
		UPDATE galleries g
		SET is_active = FALSE
		WHERE g.id <> $1 AND g.is_active = TRUE
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = g.category_id AND c.type = 'dashboard')
	`, activeID)
	return err
}

func (r *repository) Delete(ctx context.Context, payload DashboardPayload) error {
	query := `
		DELETE FROM galleries
		WHERE id = $1
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = galleries.category_id AND c.type = 'dashboard')
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
		return ErrDashboardNotFound
	}

	return nil
}

func ensureDashboardAffected(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrDashboardNotFound
	}
	return nil
}
