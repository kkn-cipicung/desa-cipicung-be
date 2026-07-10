package dashboard

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	dashboard := router.Group("/dashboard", utils.AuthMiddleware())
	{
		dashboard.POST("/create", h.Create)
		dashboard.POST("/list", h.List)
		dashboard.POST("/detail", h.FindByID)
		dashboard.POST("/update", h.Update)
		dashboard.POST("/delete", h.Delete)
	}
}
