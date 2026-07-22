package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Connect(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	DB = db
	log.Println("Successfully connected to PostgreSQL database:", cfg.DBName)
	runAutoMigrations(db)

	return db, nil
}

func runAutoMigrations(db *sqlx.DB) {
	query := `
		ALTER TABLE villages
			ADD COLUMN IF NOT EXISTS website VARCHAR(255),
			ADD COLUMN IF NOT EXISTS region TEXT,
			ADD COLUMN IF NOT EXISTS hamlet_one BIGINT,
			ADD COLUMN IF NOT EXISTS hamlet_two BIGINT,
			ADD COLUMN IF NOT EXISTS north_border TEXT,
			ADD COLUMN IF NOT EXISTS east_border TEXT,
			ADD COLUMN IF NOT EXISTS south_border TEXT,
			ADD COLUMN IF NOT EXISTS west_border TEXT,
			ADD COLUMN IF NOT EXISTS area TEXT,
			ADD COLUMN IF NOT EXISTS is_active BOOLEAN NOT NULL DEFAULT FALSE;

		CREATE TABLE IF NOT EXISTS potential_detail (
			id BIGSERIAL PRIMARY KEY,
			title TEXT NOT NULL DEFAULT '',
			detail TEXT NOT NULL DEFAULT '',
			description TEXT NOT NULL DEFAULT ''
		);
	`
	if _, err := db.Exec(query); err != nil {
		log.Println("Auto migration notice:", err)
	}
}
