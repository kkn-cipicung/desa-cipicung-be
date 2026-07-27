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
	FindOverview(ctx context.Context) (*DashboardOverviewOutput, error)
	CreateOverview(ctx context.Context, payload AddDashboardOverviewPayload, media *utils.MediaPayload) (*DashboardOverviewOutput, error)
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
		WHERE id = $2
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
		SELECT g.id, COALESCE(g.created_by, 0) AS created_by, COALESCE(u.name, '') AS creator_name,
			COALESCE(g.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(g.title, '') AS title, COALESCE(g.description, '') AS description,
			COALESCE(m.file_path, '') AS media, COALESCE(g.is_active, FALSE) AS is_active,
			COALESCE(g.created_at, NOW()) AS created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
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
		SELECT g.id, COALESCE(g.created_by, 0) AS created_by, COALESCE(u.name, '') AS creator_name,
			COALESCE(g.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(g.title, '') AS title, COALESCE(g.description, '') AS description,
			COALESCE(m.file_path, '') AS media, COALESCE(g.is_active, FALSE) AS is_active,
			COALESCE(g.created_at, NOW()) AS created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
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
		SELECT g.id, COALESCE(g.created_by, 0) AS created_by, COALESCE(u.name, '') AS creator_name,
			COALESCE(g.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
			COALESCE(g.title, '') AS title, COALESCE(g.description, '') AS description,
			COALESCE(m.file_path, '') AS media, COALESCE(g.is_active, FALSE) AS is_active,
			COALESCE(g.created_at, NOW()) AS created_at
		FROM galleries g
		LEFT JOIN users u ON g.created_by = u.id
		LEFT JOIN categories c ON g.category_id = c.id
		LEFT JOIN media m ON g.media_id = m.id
		WHERE g.is_active = TRUE
		ORDER BY g.id DESC
		LIMIT 1
	`

	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			var fallbackResult DashboardResponse
			fallbackQuery := `
				SELECT g.id, COALESCE(g.created_by, 0) AS created_by, COALESCE(u.name, '') AS creator_name,
					COALESCE(g.category_id, 0) AS category_id, COALESCE(c.name, '') AS category_name,
					COALESCE(g.title, '') AS title, COALESCE(g.description, '') AS description,
					COALESCE(m.file_path, '') AS media, COALESCE(g.is_active, FALSE) AS is_active,
					COALESCE(g.created_at, NOW()) AS created_at
				FROM galleries g
				LEFT JOIN users u ON g.created_by = u.id
				LEFT JOIN categories c ON g.category_id = c.id
				LEFT JOIN media m ON g.media_id = m.id
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

func (r *repository) FindOverview(ctx context.Context) (*DashboardOverviewOutput, error) {
	var result DashboardOverviewOutput

	query := `
		SELECT 
			COALESCE((SELECT title FROM villages WHERE title <> '' ORDER BY id DESC LIMIT 1), '') AS title,
			COALESCE((SELECT description FROM villages WHERE description <> '' ORDER BY id DESC LIMIT 1), '') AS description,
			COALESCE((SELECT m.file_path FROM villages v JOIN media m ON v.id = m.entity_id AND m.entity_type = 'village' ORDER BY v.id DESC LIMIT 1), '') AS media,
			COALESCE((SELECT area FROM villages WHERE area <> '' ORDER BY id DESC LIMIT 1), '') AS area,
			COALESCE((SELECT NULLIF(regexp_replace(COALESCE(population, ''), '[^0-9]', '', 'g'), '')::BIGINT FROM villages ORDER BY id DESC LIMIT 1), (SELECT total_population FROM villages ORDER BY id DESC LIMIT 1), 0) AS population,
			COALESCE((SELECT total_family FROM villages ORDER BY id DESC LIMIT 1), 0) AS total_family,
			COALESCE((SELECT CASE WHEN hamlet_one IS NOT NULL AND hamlet_two IS NOT NULL THEN 2 ELSE 1 END FROM villages ORDER BY id DESC LIMIT 1), 0) AS total_hamlet,
			COALESCE((SELECT COUNT(*) FROM documents), 0) AS total_news,
			COALESCE((SELECT COUNT(*) FROM potentials), 0) AS total_potential
	`

	if err := r.db.GetContext(ctx, &result, query); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) CreateOverview(ctx context.Context, payload AddDashboardOverviewPayload, media *utils.MediaPayload) (*DashboardOverviewOutput, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var existingID uint
	_ = tx.GetContext(ctx, &existingID, `SELECT id FROM villages ORDER BY id DESC LIMIT 1`)

	var villageID uint
	var oldMedia *utils.ReplacedMediaPayload

	if existingID > 0 {
		villageID = existingID
		updateQuery := `
			UPDATE villages
			SET title = $1,
				description = $2,
				area = $3,
				total_family = $4,
				population = CASE WHEN $5 > 0 THEN $5::text ELSE population END,
				total_population = CASE WHEN $5 > 0 THEN $5 ELSE total_population END,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $6
		`
		if _, err := tx.ExecContext(ctx, updateQuery,
			payload.Title, payload.Description, payload.Area, payload.TotalFamily, payload.Population, villageID,
		); err != nil {
			return nil, err
		}

		if media != nil {
			var oldMediaFilePath string
			_ = tx.GetContext(ctx, &oldMediaFilePath, `
				SELECT file_path FROM media WHERE entity_type = 'village' AND entity_id = $1 ORDER BY id DESC LIMIT 1
			`, villageID)
			if oldMediaFilePath != "" {
				oldMedia = &utils.ReplacedMediaPayload{FilePath: oldMediaFilePath}
				_, _ = tx.ExecContext(ctx, `DELETE FROM media WHERE entity_type = 'village' AND entity_id = $1`, villageID)
			}
			if _, err := utils.AttachMediaToEntity(ctx, tx, media, "village", villageID, "image"); err != nil {
				return nil, err
			}
		}
	} else {
		insertQuery := `
			INSERT INTO villages (
				name, title, description, area, total_family, population, total_population, is_active, created_at, updated_at
			) VALUES (
				$1, $1, $2, $3, $4, $5::text, $5, TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
			)
			RETURNING id
		`
		if err := tx.QueryRowContext(ctx, insertQuery,
			payload.Title, payload.Description, payload.Area, payload.TotalFamily, payload.Population,
		).Scan(&villageID); err != nil {
			return nil, err
		}

		if _, err := utils.AttachMediaToEntity(ctx, tx, media, "village", villageID, "image"); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	if oldMedia != nil {
		_ = utils.RemoveMediaFile(oldMedia)
	}

	return r.FindOverview(ctx)
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
			AND EXISTS (SELECT 1 FROM categories c WHERE c.id = $2)
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
			WHERE g.id = $1
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
	`, activeID)
	return err
}

func (r *repository) Delete(ctx context.Context, payload DashboardPayload) error {
	query := `
		DELETE FROM galleries
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
