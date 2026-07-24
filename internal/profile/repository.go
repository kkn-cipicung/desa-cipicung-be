package profile

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var ErrProfileNotFound = errors.New("profile not found")
var ErrOfficialNotFound = errors.New("official not found")

const headmanPosition = "kepala-desa"

type Repository interface {
	Create(ctx context.Context, payload AddProfilePayload) error
	Detail(ctx context.Context) (*ProfileResponse, error)
	FindActive(ctx context.Context) (*ProfileResponse, error)
	FindFirst(ctx context.Context) (*ProfileResponse, error)
	FindHeadmen(ctx context.Context, villageID uint) ([]ProfileOfficialOutput, error)
	FindRegionBoundary(ctx context.Context) (*ProfileRegionBoundaryResponse, error)
	FindVisionMission(ctx context.Context) (*ProfileVisionMissionResponse, error)
	FindGovernmentStructure(ctx context.Context) ([]GovernmentStructureResponse, error)
	FindResourcePotential(ctx context.Context) (ResourcePotentialResponse, error)
	Update(ctx context.Context, payload EditProfilePayload) error
	Activate(ctx context.Context, payload ProfilePayload) error
	Delete(ctx context.Context, payload ProfilePayload) error
	CreateOfficial(ctx context.Context, payload AddOfficialPayload) error
	ListOfficials(ctx context.Context, payload ListOfficialPayload) ([]OfficialResponse, error)
	FindOfficialByID(ctx context.Context, payload OfficialPayload) (*OfficialResponse, error)
	UpdateOfficial(ctx context.Context, payload EditOfficialPayload) error
	DeleteOfficial(ctx context.Context, payload OfficialPayload) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payload AddProfilePayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO villages (
			name, province, regency, district, postal_code, address, phone, email, website,
			latitude, longitude, vision, mission, history, description, region, hamlet_one,
			hamlet_two, north_border, east_border, south_border, west_border, area, population,
			created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15, $16, $17, $18,
			$19, $20, $21, $22, $23, $24,
			CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		)
		RETURNING id
	`
	var villageID uint
	if err := tx.QueryRowContext(ctx, query,
		payload.Name, payload.Province, payload.Regency, payload.District, payload.PostalCode,
		payload.Address, payload.Phone, payload.Email, payload.Website, payload.Latitude,
		payload.Longitude, payload.Vision, pq.Array(payload.Mission), payload.History, payload.Description,
		payload.Region, payload.HamletOne, payload.HamletTwo, payload.NorthBorder,
		payload.EastBorder, payload.SouthBorder, payload.WestBorder, payload.Area, payload.Population,
	).Scan(&villageID); err != nil {
		return err
	}

	if len(payload.Headmen) > 0 {
		for index := range payload.Headmen {
			if err := insertHeadman(ctx, tx, villageID, &payload.Headmen[index]); err != nil {
				return err
			}
		}
	} else if payload.Headman != nil {
		if err := insertHeadman(ctx, tx, villageID, payload.Headman); err != nil {
			return err
		}
	}
	if err := insertGovernmentOfficials(ctx, tx, villageID, payload.Officials); err != nil {
		return err
	}
	if err := replaceResourcePotential(ctx, tx, payload.ResourcePotential); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) FindHeadmen(ctx context.Context, villageID uint) ([]ProfileOfficialOutput, error) {
	var results []ProfileOfficialOutput
	query := `
		SELECT id, name, position, COALESCE(phone, '') AS phone, COALESCE(email, '') AS email,
			COALESCE(description, '') AS description, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, false) AS is_active,
			COALESCE(TO_CHAR(start_date, 'YYYY-MM-DD'), '') AS start_date,
			CASE WHEN finish_date IS NULL THEN NULL ELSE TO_CHAR(finish_date, 'YYYY-MM-DD') END AS finish_date
		FROM officials
		WHERE village_id = $1 AND position = $2
		ORDER BY start_date DESC NULLS LAST, id DESC
	`
	if err := r.db.SelectContext(ctx, &results, query, villageID, headmanPosition); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) Detail(ctx context.Context) (*ProfileResponse, error) {
	var result ProfileResponse
	query := profileSelectQuery() + ` ORDER BY v.id ASC LIMIT 1 `
	if err := r.db.GetContext(ctx, &result, query, headmanPosition); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ProfileResponse{Mission: []string{}}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindActive(ctx context.Context) (*ProfileResponse, error) {
	var result ProfileResponse
	query := profileSelectQuery() + ` WHERE v.is_active = TRUE ORDER BY v.id DESC LIMIT 1 `
	if err := r.db.GetContext(ctx, &result, query, headmanPosition); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			var fallbackResult ProfileResponse
			fallbackQuery := profileSelectQuery() + ` ORDER BY v.id ASC LIMIT 1 `
			if fallbackErr := r.db.GetContext(ctx, &fallbackResult, fallbackQuery, headmanPosition); fallbackErr == nil {
				return &fallbackResult, nil
			}
			return &ProfileResponse{Mission: []string{}}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) Activate(ctx context.Context, payload ProfilePayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var exists bool
	if err := tx.GetContext(ctx, &exists, `SELECT EXISTS (SELECT 1 FROM villages WHERE id = $1)`, payload.ID); err != nil {
		return err
	}
	if !exists {
		return ErrProfileNotFound
	}

	if _, err := tx.ExecContext(ctx, `UPDATE villages SET is_active = FALSE WHERE id <> $1 AND is_active = TRUE`, payload.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE villages SET is_active = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = $1`, payload.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) FindFirst(ctx context.Context) (*ProfileResponse, error) {
	var result ProfileResponse
	query := profileSelectQuery() + ` ORDER BY v.id ASC LIMIT 1 `
	if err := r.db.GetContext(ctx, &result, query, headmanPosition); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProfileNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindRegionBoundary(ctx context.Context) (*ProfileRegionBoundaryResponse, error) {
	var result ProfileRegionBoundaryResponse
	query := `
		SELECT COALESCE(region, '') AS region, COALESCE(hamlet_one, 0) AS hamlet_one,
			COALESCE(hamlet_two, 0) AS hamlet_two, COALESCE(north_border, '') AS north_border,
			COALESCE(east_border, '') AS east_border, COALESCE(south_border, '') AS south_border,
			COALESCE(west_border, '') AS west_border, COALESCE(area, '') AS area,
			COALESCE(population, '') AS population
		FROM villages
		ORDER BY id ASC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ProfileRegionBoundaryResponse{}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindVisionMission(ctx context.Context) (*ProfileVisionMissionResponse, error) {
	var result ProfileVisionMissionResponse
	query := `
		SELECT COALESCE(vision, '') AS vision, COALESCE(mission, ARRAY[]::varchar[]) AS mission
		FROM villages
		ORDER BY id ASC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ProfileVisionMissionResponse{Mission: []string{}}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindGovernmentStructure(ctx context.Context) ([]GovernmentStructureResponse, error) {
	var results []GovernmentStructureResponse
	query := `
		SELECT id, name, position, COALESCE(phone, '') AS phone, COALESCE(email, '') AS email,
			COALESCE(description, '') AS description, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, false) AS is_active,
			CASE WHEN start_date IS NULL THEN NULL ELSE TO_CHAR(start_date, 'YYYY-MM-DD') END AS start_date,
			CASE WHEN finish_date IS NULL THEN NULL ELSE TO_CHAR(finish_date, 'YYYY-MM-DD') END AS finish_date
		FROM officials
		WHERE position <> $1
		ORDER BY order_number ASC, id ASC
	`
	if err := r.db.SelectContext(ctx, &results, query, headmanPosition); err != nil {
		return nil, err
	}
	if results == nil {
		results = []GovernmentStructureResponse{}
	}
	return results, nil
}

func (r *repository) FindResourcePotential(ctx context.Context) (ResourcePotentialResponse, error) {
	var result ResourcePotentialResponse
	query := `
		SELECT COALESCE(title, '') AS title, COALESCE(detail, '') AS detail,
			COALESCE(description, '') AS description
		FROM potential_detail
		ORDER BY id ASC
		LIMIT 1
	`
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ResourcePotentialResponse{}, nil
		}
		return ResourcePotentialResponse{}, err
	}
	return result, nil
}

func (r *repository) Update(ctx context.Context, payload EditProfilePayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

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
			latitude = $11,
			longitude = $12,
			vision = $13,
			mission = $14,
			history = $15,
			description = $16,
			region = $17,
			hamlet_one = $18,
			hamlet_two = $19,
			north_border = $20,
			east_border = $21,
			south_border = $22,
			west_border = $23,
			area = $24,
			population = $25,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	result, err := tx.ExecContext(ctx, query,
		payload.ID, payload.Name, payload.Province, payload.Regency, payload.District,
		payload.PostalCode, payload.Address, payload.Phone, payload.Email, payload.Website,
		payload.Latitude, payload.Longitude, payload.Vision, pq.Array(payload.Mission), payload.History,
		payload.Description, payload.Region, payload.HamletOne, payload.HamletTwo, payload.NorthBorder,
		payload.EastBorder, payload.SouthBorder, payload.WestBorder, payload.Area, payload.Population,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrProfileNotFound
	}

	if len(payload.Headmen) > 0 {
		if _, err := tx.ExecContext(ctx, `DELETE FROM officials WHERE village_id = $1 AND position = $2`, payload.ID, headmanPosition); err != nil {
			return err
		}
		for index := range payload.Headmen {
			if err := insertHeadman(ctx, tx, payload.ID, &payload.Headmen[index]); err != nil {
				return err
			}
		}
	} else if payload.Headman != nil {
		if err := upsertHeadman(ctx, tx, payload.ID, payload.Headman); err != nil {
			return err
		}
	}
	if len(payload.Officials) > 0 {
		if _, err := tx.ExecContext(ctx, `DELETE FROM officials WHERE village_id = $1 AND position <> $2`, payload.ID, headmanPosition); err != nil {
			return err
		}
	}
	if err := insertGovernmentOfficials(ctx, tx, payload.ID, payload.Officials); err != nil {
		return err
	}
	if err := replaceResourcePotential(ctx, tx, payload.ResourcePotential); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repository) Delete(ctx context.Context, payload ProfilePayload) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `DELETE FROM officials WHERE village_id = $1`, payload.ID)
	if err != nil {
		return err
	}

	query := `DELETE FROM villages WHERE id = $1`
	result, err := tx.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrProfileNotFound
	}
	return tx.Commit()
}

func profileSelectQuery() string {
	return `
		SELECT v.id, v.name, v.province, v.regency, v.district, v.postal_code, v.address,
			COALESCE(v.phone, '') AS phone, COALESCE(v.email, '') AS email, COALESCE(v.website, '') AS website,
			COALESCE(v.latitude, 0) AS latitude,
			COALESCE(v.longitude, 0) AS longitude, COALESCE(v.vision, '') AS vision,
			COALESCE(v.mission, ARRAY[]::varchar[]) AS mission,
			COALESCE(v.history, '') AS history, COALESCE(v.description, '') AS description,
			COALESCE(v.region, '') AS region, COALESCE(v.hamlet_one, 0) AS hamlet_one,
			COALESCE(v.hamlet_two, 0) AS hamlet_two, COALESCE(v.north_border, '') AS north_border,
			COALESCE(v.east_border, '') AS east_border, COALESCE(v.south_border, '') AS south_border,
			COALESCE(v.west_border, '') AS west_border, COALESCE(v.area, '') AS area,
			COALESCE(v.population, '') AS population,
			COALESCE(v.is_active, FALSE) AS is_active,
			v.created_at, v.updated_at,
			COALESCE(o.id, 0) AS headman_id,
			COALESCE(o.name, '') AS headman_name,
			COALESCE(o.position, '') AS headman_position,
			COALESCE(o.phone, '') AS headman_phone,
			COALESCE(o.email, '') AS headman_email,
			COALESCE(o.description, '') AS headman_description,
			COALESCE(o.order_number, 0) AS headman_order_number,
			COALESCE(o.is_active, false) AS headman_is_active,
			o.start_date AS headman_start_date, o.finish_date AS headman_finish_date
		FROM villages v
		LEFT JOIN LATERAL (
			SELECT * FROM officials
			WHERE village_id = v.id AND position = $1
			ORDER BY is_active DESC, finish_date DESC NULLS FIRST, start_date DESC NULLS LAST, id DESC
			LIMIT 1
		) o ON TRUE
	`
}

func insertHeadman(ctx context.Context, tx *sqlx.Tx, villageID uint, headman *ProfileOfficialInput) error {
	query := `
		INSERT INTO officials (
			village_id, name, position, phone, email, description, order_number, is_active,
			start_date, finish_date, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := tx.ExecContext(ctx, query,
		villageID, headman.Name, headman.Position, headman.Phone, headman.Email,
		headman.Description, headman.OrderNumber, headman.IsActive, headman.StartDate, headman.FinishDate,
	)
	return err
}

func upsertHeadman(ctx context.Context, tx *sqlx.Tx, villageID uint, headman *ProfileOfficialInput) error {
	query := `
		UPDATE officials
		SET name = $1,
			phone = $2,
			email = $3,
			description = $4,
			order_number = $5,
			is_active = $6,
			start_date = $7,
			finish_date = $8,
			updated_at = CURRENT_TIMESTAMP
		WHERE village_id = $9 AND position = $10
	`
	result, err := tx.ExecContext(ctx, query,
		headman.Name, headman.Phone, headman.Email, headman.Description,
		headman.OrderNumber, headman.IsActive, headman.StartDate, headman.FinishDate, villageID, headmanPosition,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected > 0 {
		return nil
	}
	return insertHeadman(ctx, tx, villageID, headman)
}

func insertGovernmentOfficials(ctx context.Context, tx *sqlx.Tx, villageID uint, officials []GovernmentOfficialInput) error {
	query := `
		INSERT INTO officials (
			village_id, name, position, phone, email, description, order_number, is_active,
			start_date, finish_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	for _, official := range officials {
		if _, err := tx.ExecContext(ctx, query, villageID, official.Name, official.Position,
			official.Phone, official.Email, official.Description, official.OrderNumber,
			official.IsActive, official.StartDate, official.FinishDate); err != nil {
			return err
		}
	}
	return nil
}

func replaceResourcePotential(ctx context.Context, tx *sqlx.Tx, resource *ResourcePotentialInput) error {
	if resource == nil {
		return nil
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM potential_detail`); err != nil {
		return err
	}
	_, err := tx.ExecContext(ctx,
		`INSERT INTO potential_detail (title, detail, description) VALUES ($1, $2, $3)`,
		resource.Title, resource.Detail, resource.Description,
	)
	return err
}

func (r *repository) CreateOfficial(ctx context.Context, payload AddOfficialPayload) error {
	villageID := payload.VillageID
	if villageID == 0 {
		_ = r.db.GetContext(ctx, &villageID, `SELECT id FROM villages ORDER BY is_active DESC, id ASC LIMIT 1`)
	}
	if villageID == 0 {
		villageID = 1
	}

	query := `
		INSERT INTO officials (
			village_id, name, position, phone, email, description, order_number, is_active,
			start_date, finish_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := r.db.ExecContext(ctx, query,
		villageID, payload.Name, payload.Position, payload.Phone, payload.Email,
		payload.Description, payload.OrderNumber, payload.IsActive, payload.StartDate, payload.FinishDate,
	)
	return err
}

func (r *repository) ListOfficials(ctx context.Context, payload ListOfficialPayload) ([]OfficialResponse, error) {
	var results []OfficialResponse
	query := `
		SELECT id, village_id, name, position, COALESCE(phone, '') AS phone, COALESCE(email, '') AS email,
			COALESCE(description, '') AS description, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, false) AS is_active,
			CASE WHEN start_date IS NULL THEN NULL ELSE TO_CHAR(start_date, 'YYYY-MM-DD') END AS start_date,
			CASE WHEN finish_date IS NULL THEN NULL ELSE TO_CHAR(finish_date, 'YYYY-MM-DD') END AS finish_date,
			COALESCE(TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS created_at,
			COALESCE(TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS updated_at
		FROM officials
		WHERE ($3 = 0 OR village_id = $3)
		ORDER BY order_number ASC, id ASC
		LIMIT $1 OFFSET $2
	`
	if err := r.db.SelectContext(ctx, &results, query, payload.Limit, payload.Index, payload.VillageID); err != nil {
		return nil, err
	}
	if results == nil {
		results = []OfficialResponse{}
	}
	return results, nil
}

func (r *repository) FindOfficialByID(ctx context.Context, payload OfficialPayload) (*OfficialResponse, error) {
	var result OfficialResponse
	query := `
		SELECT id, village_id, name, position, COALESCE(phone, '') AS phone, COALESCE(email, '') AS email,
			COALESCE(description, '') AS description, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, false) AS is_active,
			CASE WHEN start_date IS NULL THEN NULL ELSE TO_CHAR(start_date, 'YYYY-MM-DD') END AS start_date,
			CASE WHEN finish_date IS NULL THEN NULL ELSE TO_CHAR(finish_date, 'YYYY-MM-DD') END AS finish_date,
			COALESCE(TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS created_at,
			COALESCE(TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS updated_at
		FROM officials
		WHERE id = $1
	`
	if err := r.db.GetContext(ctx, &result, query, payload.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrOfficialNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) UpdateOfficial(ctx context.Context, payload EditOfficialPayload) error {
	query := `
		UPDATE officials
		SET name = $2,
			position = $3,
			phone = $4,
			email = $5,
			description = $6,
			order_number = $7,
			is_active = $8,
			start_date = $9,
			finish_date = $10,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query,
		payload.ID, payload.Name, payload.Position, payload.Phone, payload.Email,
		payload.Description, payload.OrderNumber, payload.IsActive, payload.StartDate, payload.FinishDate,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrOfficialNotFound
	}
	return nil
}

func (r *repository) DeleteOfficial(ctx context.Context, payload OfficialPayload) error {
	query := `DELETE FROM officials WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrOfficialNotFound
	}
	return nil
}
