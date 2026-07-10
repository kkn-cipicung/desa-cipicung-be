package category

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	category := router.Group("/category", utils.AuthMiddleware())
	{
		category.POST("/create", h.Create)
		category.POST("/list", h.List)
		category.POST("/detail", h.FindByID)
		category.POST("/update", h.Update)
		category.POST("/delete", h.Delete)
	}
}
