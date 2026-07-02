package config

import (
	"log"
	"os"
	"strconv"

	"cipicung.id/be/pkg/database"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DB   database.Config
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return Config{
		Port: getEnv("PORT", "8080"),
		DB: database.Config{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "cipicung"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
