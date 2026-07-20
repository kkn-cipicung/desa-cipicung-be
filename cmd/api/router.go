package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"cipicung.id/be/internal/auth"
	"cipicung.id/be/internal/business"
	"cipicung.id/be/internal/category"
	"cipicung.id/be/internal/contact"
	"cipicung.id/be/internal/dashboard"
	"cipicung.id/be/internal/gallery"
	mapdata "cipicung.id/be/internal/map"
	"cipicung.id/be/internal/news"
	"cipicung.id/be/internal/potential"
	"cipicung.id/be/internal/profile"
	"cipicung.id/be/utils"
)

func setupRouter(r *gin.Engine, db *sqlx.DB) {
	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
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
	r.Static("/uploads", "./uploads")

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		utils.SuccessResponse(c, http.StatusOK, "pong", nil)
	})

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(api)

	dashboardRepository := dashboard.NewRepository(db)
	dashboardService := dashboard.NewService(dashboardRepository)
	dashboardHandler := dashboard.NewHandler(dashboardService)
	dashboardHandler.RegisterRoutes(api)

	mapRepository := mapdata.NewRepository(db)
	mapService := mapdata.NewService(mapRepository)
	mapHandler := mapdata.NewHandler(mapService)
	mapHandler.RegisterRoutes(api)

	galleryRepository := gallery.NewRepository(db)
	galleryService := gallery.NewService(galleryRepository)
	galleryHandler := gallery.NewHandler(galleryService)
	galleryHandler.RegisterRoutes(api)

	contactRepository := contact.NewRepository(db)
	contactService := contact.NewService(contactRepository)
	contactHandler := contact.NewHandler(contactService)
	contactHandler.RegisterRoutes(api)

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

	profileRepository := profile.NewRepository(db)
	profileService := profile.NewService(profileRepository)
	profileHandler := profile.NewHandler(profileService)
	profileHandler.RegisterRoutes(api)

	businessRepository := business.NewRepository(db)
	businessService := business.NewService(businessRepository)
	businessHandler := business.NewHandler(businessService)
	businessHandler.RegisterRoutes(api)
}
