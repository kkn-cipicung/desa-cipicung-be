package infografis

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	infografis := router.Group("/infografis")
	{
		infografis.POST("/detail", h.Detail)
	}
}
