package news

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	news := router.Group("/news", utils.AuthMiddleware())
	{
		news.POST("/create", h.Create)
		news.POST("/list", h.List)
		news.POST("/detail", h.FindByID)
		news.POST("/update", h.Update)
		news.POST("/delete", h.Delete)
		news.POST("/find-by-date", h.FindByDate)
	}
}
