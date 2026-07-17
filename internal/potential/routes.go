package potential

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	potential := router.Group("/potential")
	{
		potential.POST("/create", utils.AuthMiddleware(), h.Create)
		potential.POST("/list", h.List)
		potential.POST("/detail", h.FindByID)
		potential.POST("/update", utils.AuthMiddleware(), h.Update)
		potential.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
