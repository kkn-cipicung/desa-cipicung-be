package profile

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/pkg/types"
	"github.com/jmoiron/sqlx"
)

var ErrProfileNotFound = errors.New("profile not found")
var ErrOfficialNotFound = errors.New("official not found")

const headmanPosition = "kepala-desa"

type Repository interface {
	Create(ctx context.Context, payload AddProfilePayload) error
	Detail(ctx context.Context) (*ProfileResponse, error)
	FindFirst(ctx context.Context) (*ProfileResponse, error)
	FindHeadmen(ctx context.Context, villageID uint) ([]ProfileOfficialOutput, error)
	FindRegionBoundary(ctx context.Context) (*ProfileRegionBoundaryResponse, error)
	FindVisionMission(ctx context.Context) (*ProfileVisionMissionResponse, error)
	FindGovernmentStructure(ctx context.Context) ([]GovernmentStructureResponse, error)
	FindResourcePotential(ctx context.Context) (ResourcePotentialResponse, error)
	Update(ctx context.Context, payload EditProfilePayload) error
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
			hamlet_two, total_family, total_rt, total_rw, rt_hamlet_one, rt_hamlet_two, rw_hamlet_one,
			rw_hamlet_two, north_border, east_border, south_border, west_border, area, population,
			total_male, total_female, demographic_religions, demographic_religion_rt,
			demographic_education, demographic_occupation, demographic_ages,
			created_at, updated_at
		)
		VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?,
			CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		)
	`
	result, err := tx.ExecContext(ctx, query, profileVillageValues(payload)...)
	if err != nil {
		return err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	villageID := uint(insertID)

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
		SELECT id, COALESCE(name, '') AS name, COALESCE(position, '') AS position,
			COALESCE(phone, '') AS phone, COALESCE(email, '') AS email,
			COALESCE(description, '') AS description, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, FALSE) AS is_active,
			COALESCE(DATE_FORMAT(start_date, '%Y-%m-%d'), '') AS start_date,
			CASE WHEN finish_date IS NULL THEN NULL ELSE DATE_FORMAT(finish_date, '%Y-%m-%d') END AS finish_date
		FROM officials
		WHERE village_id = ? AND position = ?
		ORDER BY start_date IS NULL ASC, start_date DESC, id DESC
	`
	if err := r.db.SelectContext(ctx, &results, query, villageID, headmanPosition); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) Detail(ctx context.Context) (*ProfileResponse, error) {
	var result ProfileResponse
	query := profileSelectQuery() + ` ORDER BY v.id ASC LIMIT 1 `
	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ProfileResponse{Mission: []string{}}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindFirst(ctx context.Context) (*ProfileResponse, error) {
	var result ProfileResponse
	query := profileSelectQuery() + ` ORDER BY v.id ASC LIMIT 1 `
	if err := r.db.GetContext(ctx, &result, query); err != nil {
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
			COALESCE(hamlet_two, 0) AS hamlet_two, COALESCE(total_family, 0) AS total_family,
			COALESCE(rt_hamlet_one, 0) AS rt_hamlet_one, COALESCE(rt_hamlet_two, 0) AS rt_hamlet_two,
			COALESCE(rw_hamlet_one, 0) AS rw_hamlet_one, COALESCE(rw_hamlet_two, 0) AS rw_hamlet_two,
			CASE
				WHEN COALESCE(rt_hamlet_one, 0) + COALESCE(rt_hamlet_two, 0) > 0 THEN COALESCE(rt_hamlet_one, 0) + COALESCE(rt_hamlet_two, 0)
				ELSE COALESCE(total_rt, 0)
			END AS total_rt,
			CASE
				WHEN COALESCE(rw_hamlet_one, 0) + COALESCE(rw_hamlet_two, 0) > 0 THEN COALESCE(rw_hamlet_one, 0) + COALESCE(rw_hamlet_two, 0)
				ELSE COALESCE(total_rw, 0)
			END AS total_rw,
			COALESCE(north_border, '') AS north_border,
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
		SELECT COALESCE(vision, '') AS vision, COALESCE(mission, '[]') AS mission
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
		SELECT id, COALESCE(name, '') AS name, COALESCE(position, '') AS position,
			COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, FALSE) AS is_active
		FROM officials
		WHERE is_active = TRUE
		ORDER BY
			CASE WHEN position = ? AND is_active = TRUE THEN 0 ELSE 1 END,
			order_number ASC,
			id ASC
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
		SET name = ?,
			province = ?,
			regency = ?,
			district = ?,
			postal_code = ?,
			address = ?,
			phone = ?,
			email = ?,
			website = ?,
			latitude = ?,
			longitude = ?,
			vision = ?,
			mission = ?,
			history = ?,
			description = ?,
			region = ?,
			hamlet_one = ?,
			hamlet_two = ?,
			total_family = ?,
			total_rt = ?,
			total_rw = ?,
			rt_hamlet_one = ?,
			rt_hamlet_two = ?,
			rw_hamlet_one = ?,
			rw_hamlet_two = ?,
			north_border = ?,
			east_border = ?,
			south_border = ?,
			west_border = ?,
			area = ?,
			population = ?,
			total_male = ?,
			total_female = ?,
			demographic_religions = ?,
			demographic_religion_rt = ?,
			demographic_education = ?,
			demographic_occupation = ?,
			demographic_ages = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	args := append(profileVillageValues(payload.AddProfilePayload), payload.ID)
	result, err := tx.ExecContext(ctx, query, args...)
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
		if _, err := tx.ExecContext(ctx, `DELETE FROM officials WHERE village_id = ? AND position = ?`, payload.ID, headmanPosition); err != nil {
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
		if _, err := tx.ExecContext(ctx, `DELETE FROM officials WHERE village_id = ? AND position <> ?`, payload.ID, headmanPosition); err != nil {
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

	if _, err := tx.ExecContext(ctx, `DELETE FROM officials WHERE village_id = ?`, payload.ID); err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, `DELETE FROM villages WHERE id = ?`, payload.ID)
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
		SELECT v.id, COALESCE(v.name, '') AS name,
			COALESCE(v.province, '') AS province,
			COALESCE(v.regency, '') AS regency,
			COALESCE(v.district, '') AS district,
			COALESCE(v.postal_code, '') AS postal_code,
			COALESCE(v.address, '') AS address,
			COALESCE(v.phone, '') AS phone, COALESCE(v.email, '') AS email, COALESCE(v.website, '') AS website,
			COALESCE(v.latitude, 0) AS latitude,
			COALESCE(v.longitude, 0) AS longitude, COALESCE(v.vision, '') AS vision,
			COALESCE(v.mission, '[]') AS mission,
			COALESCE(v.history, '') AS history, COALESCE(v.description, '') AS description,
			COALESCE(v.region, '') AS region,
			COALESCE(v.hamlet_one, 0) AS hamlet_one,
			COALESCE(v.hamlet_two, 0) AS hamlet_two,
			COALESCE(v.total_family, 0) AS total_family,
			COALESCE(v.rt_hamlet_one, 0) AS rt_hamlet_one,
			COALESCE(v.rt_hamlet_two, 0) AS rt_hamlet_two,
			COALESCE(v.rw_hamlet_one, 0) AS rw_hamlet_one,
			COALESCE(v.rw_hamlet_two, 0) AS rw_hamlet_two,
			CASE
				WHEN COALESCE(v.rt_hamlet_one, 0) + COALESCE(v.rt_hamlet_two, 0) > 0 THEN COALESCE(v.rt_hamlet_one, 0) + COALESCE(v.rt_hamlet_two, 0)
				ELSE COALESCE(v.total_rt, 0)
			END AS total_rt,
			CASE
				WHEN COALESCE(v.rw_hamlet_one, 0) + COALESCE(v.rw_hamlet_two, 0) > 0 THEN COALESCE(v.rw_hamlet_one, 0) + COALESCE(v.rw_hamlet_two, 0)
				ELSE COALESCE(v.total_rw, 0)
			END AS total_rw,
			COALESCE(v.north_border, '') AS north_border,
			COALESCE(v.east_border, '') AS east_border, COALESCE(v.south_border, '') AS south_border,
			COALESCE(v.west_border, '') AS west_border, COALESCE(v.area, '') AS area,
			COALESCE(v.population, '') AS population,
			COALESCE(v.total_male, 0) AS total_male,
			COALESCE(v.total_female, 0) AS total_female,
			COALESCE(v.demographic_religions, '[]') AS demographic_religions,
			COALESCE(v.demographic_religion_rt, '[]') AS demographic_religion_rt,
			COALESCE(v.demographic_education, '[]') AS demographic_education,
			COALESCE(v.demographic_occupation, '[]') AS demographic_occupation,
			COALESCE(v.demographic_ages, '[]') AS demographic_ages,
			COALESCE(v.is_active, FALSE) AS is_active,
			COALESCE(v.created_at, NOW()) AS created_at,
			COALESCE(v.updated_at, NOW()) AS updated_at,
			COALESCE(o.id, 0) AS headman_id,
			COALESCE(o.name, '') AS headman_name,
			COALESCE(o.position, '') AS headman_position,
			COALESCE(o.phone, '') AS headman_phone,
			COALESCE(o.email, '') AS headman_email,
			COALESCE(o.description, '') AS headman_description,
			COALESCE(o.order_number, 0) AS headman_order_number,
			COALESCE(o.is_active, FALSE) AS headman_is_active,
			o.start_date AS headman_start_date, o.finish_date AS headman_finish_date
		FROM villages v
		LEFT JOIN officials o ON o.id = (
			SELECT oi.id
			FROM officials oi
			WHERE oi.village_id = v.id AND (oi.position = 'kepala-desa' OR oi.position = 'kepala desa')
			ORDER BY oi.is_active DESC, oi.finish_date IS NOT NULL ASC, oi.finish_date DESC,
				oi.start_date IS NULL ASC, oi.start_date DESC, oi.id DESC
			LIMIT 1
		)
	`
}

func profileVillageValues(payload AddProfilePayload) []any {
	return []any{
		payload.Name, payload.Province, payload.Regency, payload.District, payload.PostalCode,
		payload.Address, payload.Phone, payload.Email, payload.Website, payload.Latitude,
		payload.Longitude, payload.Vision, types.JSONStringArray(payload.Mission), payload.History, payload.Description,
		payload.Region, payload.HamletOne, payload.HamletTwo, payload.TotalFamily, payload.TotalRT,
		payload.TotalRW, payload.RTHamletOne, payload.RTHamletTwo, payload.RWHamletOne, payload.RWHamletTwo,
		payload.NorthBorder, payload.EastBorder, payload.SouthBorder, payload.WestBorder,
		payload.Area, payload.Population,
		payload.TotalMale, payload.TotalFemale, demographicJSON(payload.DemographicReligions),
		demographicJSON(payload.DemographicReligionRT), demographicJSON(payload.DemographicEducation),
		demographicJSON(payload.DemographicOccupation), demographicJSON(payload.DemographicAges),
	}
}

func insertHeadman(ctx context.Context, tx *sqlx.Tx, villageID uint, headman *ProfileOfficialInput) error {
	query := `
		INSERT INTO officials (
			village_id, name, position, phone, email, description, order_number, is_active,
			start_date, finish_date, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
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
		SET name = ?,
			phone = ?,
			email = ?,
			description = ?,
			order_number = ?,
			is_active = ?,
			start_date = ?,
			finish_date = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE village_id = ? AND position = ?
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
			village_id, name, position, order_number, is_active, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	for _, official := range officials {
		if _, err := tx.ExecContext(ctx, query, villageID, official.Name, official.Position,
			official.OrderNumber, official.IsActive); err != nil {
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
		`INSERT INTO potential_detail (title, detail, description) VALUES (?, ?, ?)`,
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
			village_id, name, position, order_number, is_active, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := r.db.ExecContext(ctx, query,
		villageID, payload.Name, payload.Position, payload.OrderNumber, payload.IsActive,
	)
	return err
}

func (r *repository) ListOfficials(ctx context.Context, payload ListOfficialPayload) ([]OfficialResponse, error) {
	var results []OfficialResponse
	query := `
		SELECT id, COALESCE(village_id, 0) AS village_id, COALESCE(name, '') AS name,
			COALESCE(position, '') AS position, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, FALSE) AS is_active,
			COALESCE(DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s'), '') AS created_at,
			COALESCE(DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s'), '') AS updated_at
		FROM officials
		WHERE (? = 0 OR village_id = ?) AND position <> ?
		ORDER BY order_number ASC, id ASC
		LIMIT ? OFFSET ?
	`
	if err := r.db.SelectContext(ctx, &results, query, payload.VillageID, payload.VillageID, headmanPosition, payload.Limit, payload.Index); err != nil {
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
		SELECT id, COALESCE(village_id, 0) AS village_id, COALESCE(name, '') AS name,
			COALESCE(position, '') AS position, COALESCE(order_number, 0) AS order_number,
			COALESCE(is_active, FALSE) AS is_active,
			COALESCE(DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s'), '') AS created_at,
			COALESCE(DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s'), '') AS updated_at
		FROM officials
		WHERE id = ?
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
		SET name = ?,
			position = ?,
			order_number = ?,
			is_active = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	result, err := r.db.ExecContext(ctx, query,
		payload.Name, payload.Position, payload.OrderNumber, payload.IsActive, payload.ID,
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
	query := `DELETE FROM officials WHERE id = ?`
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
