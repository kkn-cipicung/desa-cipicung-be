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

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "3306"))

	return Config{
		Port: getEnv("PORT", "8080"),
		DB: database.Config{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
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
