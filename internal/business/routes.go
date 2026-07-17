package business

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	business := router.Group("/business")
	{
		business.POST("/create", utils.AuthMiddleware(), h.Create)
		business.POST("/list", h.List)
		business.POST("/detail", h.FindByID)
		business.POST("/update", utils.AuthMiddleware(), h.Update)
		business.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
