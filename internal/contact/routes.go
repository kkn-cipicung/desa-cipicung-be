package contact

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	contact := router.Group("/contact")
	{
		contact.POST("/create", utils.AuthMiddleware(), h.Create)
		contact.POST("/detail", h.Detail)
		contact.POST("/update", utils.AuthMiddleware(), h.Update)
		contact.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
