package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"cipicung.id/be/internal/auth"
	"cipicung.id/be/internal/business"
	"cipicung.id/be/internal/category"
	"cipicung.id/be/internal/dashboard"
	"cipicung.id/be/internal/news"
	"cipicung.id/be/internal/potential"
)

func setupRouter(r *gin.Engine, db *sqlx.DB) {
	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:3020",
			"https://admin.cipicung.id",
			"https://cipicung.id",
		},
		AllowMethods: []string{"POST", "GET"},
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

	businessRepository := business.NewRepository(db)
	businessService := business.NewService(businessRepository)
	businessHandler := business.NewHandler(businessService)
	businessHandler.RegisterRoutes(api)
}
