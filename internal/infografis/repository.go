package infografis

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Detail(ctx context.Context) (*VillageStats, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Detail(ctx context.Context) (*VillageStats, error) {
	var result VillageStats
	query := `
		SELECT
			COALESCE(NULLIF(population, '')::bigint, total_population, 0) AS population,
			COALESCE(total_family, 0) AS family,
			COALESCE(total_male, 0) AS male,
			COALESCE(total_female, 0) AS female,
			COALESCE(hamlet_one, 0) AS hamlet_one,
			COALESCE(hamlet_two, 0) AS hamlet_two,
			COALESCE(demographic_religions, '[]'::jsonb)::text AS religions,
			COALESCE(demographic_religion_rt, '[]'::jsonb)::text AS religion_rt,
			COALESCE(demographic_education, '[]'::jsonb)::text AS education,
			COALESCE(demographic_occupation, '[]'::jsonb)::text AS occupation,
			COALESCE(demographic_ages, '[]'::jsonb)::text AS ages
		FROM villages
		ORDER BY is_active DESC, id DESC
		LIMIT 1
	`

	if err := r.db.GetContext(ctx, &result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &VillageStats{}, nil
		}
		return nil, err
	}

	return &result, nil
}
