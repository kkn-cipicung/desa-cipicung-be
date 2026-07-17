package category

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	category := router.Group("/category")
	{
		category.POST("/create", utils.AuthMiddleware(), h.Create)
		category.POST("/list", h.List)
		category.POST("/detail", h.FindByID)
		category.POST("/update", utils.AuthMiddleware(), h.Update)
		category.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
