package business

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	business := router.Group("/business", utils.AuthMiddleware())
	{
		business.POST("/create", h.Create)
		business.POST("/list", h.List)
		business.POST("/detail", h.FindByID)
		business.POST("/update", h.Update)
		business.POST("/delete", h.Delete)
	}
}
