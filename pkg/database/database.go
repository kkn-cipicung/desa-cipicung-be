package database

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	db, err := sqlx.Open("mysql", dsn)
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
	log.Println("Successfully connected to MariaDB database:", cfg.DBName)
	runAutoMigrations(db)

	return db, nil
}

func runAutoMigrations(db *sqlx.DB) {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS potential_detail (
			id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL DEFAULT '',
			detail VARCHAR(255) NOT NULL DEFAULT '',
			description VARCHAR(255) NOT NULL DEFAULT ''
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,
		`ALTER TABLE categories MODIFY COLUMN type VARCHAR(50) NOT NULL DEFAULT ''`,
		`UPDATE galleries g JOIN categories c ON g.category_id = c.id
			SET g.type = c.type
			WHERE c.type IN ('dashboard', 'gallery')`,
		`CREATE INDEX idx_galleries_type ON galleries(type)`,
	}

	for _, query := range statements {
		if _, err := db.Exec(query); err != nil {
			log.Println("Auto migration notice:", err)
		}
	}
}
