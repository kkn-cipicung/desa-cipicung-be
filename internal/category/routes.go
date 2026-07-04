package category

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	// category := router.Group("/category", utils.AuthMiddleware())
	category := router.Group("/category")
	{
		category.POST("/create", h.Create)
		category.POST("/list", h.List)
		category.POST("/find", h.FindByID)
		category.POST("/update", h.Update)
		category.POST("/delete", h.Delete)
	}
}
