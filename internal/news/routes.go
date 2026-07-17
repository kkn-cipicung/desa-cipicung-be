package news

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	news := router.Group("/news")
	{
		news.POST("/create", utils.AuthMiddleware(), h.Create)
		news.POST("/list", h.List)
		news.POST("/detail", h.FindByID)
		news.POST("/update", utils.AuthMiddleware(), h.Update)
		news.POST("/delete", utils.AuthMiddleware(), h.Delete)
		news.POST("/find-by-date", h.FindByDate)
	}
}
