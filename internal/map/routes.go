package mapdata

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	mapGroup := router.Group("/map")
	{
		mapGroup.POST("/create", utils.AuthMiddleware(), h.Create)
		mapGroup.POST("/main", h.FindActive)
		mapGroup.POST("/", h.FindActive)
		mapGroup.POST("", h.FindActive)
		mapGroup.POST("/list", h.List)
		mapGroup.POST("/update", utils.AuthMiddleware(), h.Update)
		mapGroup.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
