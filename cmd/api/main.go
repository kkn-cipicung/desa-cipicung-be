package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"cipicung.id/be/config"
	"cipicung.id/be/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	_, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Printf("Server is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
