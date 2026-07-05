package potential

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	potential := router.Group("/potential")
	{
		potential.POST("/create", h.Create)
		potential.POST("/list", h.List)
		potential.POST("/find", h.FindByID)
		potential.POST("/update", h.Update)
		potential.POST("/delete", h.Delete)
	}
}
