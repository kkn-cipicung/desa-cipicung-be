package profile

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	profile := router.Group("/profile")
	{
		profile.POST("/create", utils.AuthMiddleware(), h.Create)
		profile.POST("/detail", h.Detail)
		profile.POST("/region-boundary", h.FindRegionBoundary)
		profile.POST("/vision-mission", h.FindVisionMission)
		profile.POST("/government-structure", h.FindGovernmentStructure)
		profile.POST("/resource-potential", h.FindResourcePotential)
		profile.POST("/delete", utils.AuthMiddleware(), h.Delete)
	}
}
