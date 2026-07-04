package news

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	news := router.Group("/news")
	{
		news.POST("/create", h.Create)
		news.POST("/list", h.List)
		news.POST("/find", h.FindByID)
		news.POST("/update", h.Update)
		news.POST("/delete", h.Delete)
		news.POST("/find-by-date", h.FindByDate)
	}
}
