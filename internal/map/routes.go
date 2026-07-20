package mapdata

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	mapGroup := router.Group("/map")
	{
		mapGroup.POST("/create", utils.AuthMiddleware(), h.Create)
		mapGroup.POST("/detail", h.Detail)
		mapGroup.POST("/active", h.FindActive)
		mapGroup.POST("/update", utils.AuthMiddleware(), h.Update)
		mapGroup.POST("/activate", utils.AuthMiddleware(), h.Activate)
		mapGroup.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
