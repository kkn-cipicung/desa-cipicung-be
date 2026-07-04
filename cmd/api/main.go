package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"cipicung.id/be/config"
	"cipicung.id/be/internal/auth"
	"cipicung.id/be/internal/category"
	"cipicung.id/be/internal/dashboard"
	"cipicung.id/be/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(api)

	dashboardRepository := dashboard.NewRepository(db)
	dashboardService := dashboard.NewService(dashboardRepository)
	dashboardHandler := dashboard.NewHandler(dashboardService)
	dashboardHandler.RegisterRoutes(api)

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)
	categoryHandler.RegisterRoutes(api)

	log.Printf("Server is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
