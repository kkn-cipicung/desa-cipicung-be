package gallery

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	gallery := router.Group("/gallery")
	{
		gallery.POST("/create", utils.AuthMiddleware(), h.Create)
		gallery.POST("/list", h.List)
		gallery.POST("/detail", h.FindByID)
		gallery.POST("/update", utils.AuthMiddleware(), h.Update)
		gallery.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
