package potential

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	potential := router.Group("/potential", utils.AuthMiddleware())
	{
		potential.POST("/create", h.Create)
		potential.POST("/list", h.List)
		potential.POST("/detail", h.FindByID)
		potential.POST("/update", h.Update)
		potential.POST("/delete", h.Delete)
	}
}
