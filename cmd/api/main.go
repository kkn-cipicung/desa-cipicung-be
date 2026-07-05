package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cipicung.id/be/config"
	"cipicung.id/be/internal/auth"
	"cipicung.id/be/internal/category"
	"cipicung.id/be/internal/dashboard"
	"cipicung.id/be/internal/news"
	"cipicung.id/be/internal/potential"
	"cipicung.id/be/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://admin.cipicung.id",
			"https://cipicung.id",
		},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           43200 * time.Second,
	}
	r.Use(cors.New(corsConfig))

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

	newsRepository := news.NewRepository(db)
	newsService := news.NewService(newsRepository)
	newsHandler := news.NewHandler(newsService)
	newsHandler.RegisterRoutes(api)

	potentialRepository := potential.NewRepository(db)
	potentialService := potential.NewService(potentialRepository)
	potentialHandler := potential.NewHandler(potentialService)
	potentialHandler.RegisterRoutes(api)

	log.Printf("Server is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
