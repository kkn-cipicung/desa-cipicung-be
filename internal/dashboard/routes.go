package dashboard

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	dashboard := router.Group("/dashboard")
	{
		dashboard.POST("/create", utils.AuthMiddleware(), h.Create)
		dashboard.POST("/list", h.List)
		dashboard.POST("/detail", h.Detail)
		dashboard.POST("/active", h.FindActive)
		dashboard.POST("/update", utils.AuthMiddleware(), h.Update)
		dashboard.POST("/activate", utils.AuthMiddleware(), h.Activate)
		dashboard.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
