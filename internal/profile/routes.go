package profile

import (
	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	profile := router.Group("/profile")
	{
		profile.POST("", h.Detail)
		profile.POST("/detail", h.Detail)
		profile.POST("/active", h.Detail)
		profile.POST("/create", utils.AuthMiddleware(), h.Create)
		profile.POST("/region-boundary", h.FindRegionBoundary)
		profile.POST("/vision-mission", h.FindVisionMission)
		profile.POST("/government-structure", h.FindGovernmentStructure)
		profile.POST("/resource-potential", h.FindResourcePotential)
		profile.POST("/delete", utils.AuthMiddleware(), h.Delete)

		// Official CRUD endpoints (supporting both /official/... and /officials/...)
		profile.POST("/official/create", utils.AuthMiddleware(), h.CreateOfficial)
		profile.POST("/official/list", h.ListOfficials)
		profile.POST("/official/detail", h.DetailOfficial)
		profile.POST("/official/update", utils.AuthMiddleware(), h.UpdateOfficial)
		profile.POST("/official/delete", utils.AuthMiddleware(), h.DeleteOfficial)

		profile.POST("/officials/create", utils.AuthMiddleware(), h.CreateOfficial)
		profile.POST("/officials/list", h.ListOfficials)
		profile.POST("/officials/detail", h.DetailOfficial)
		profile.POST("/officials/update", utils.AuthMiddleware(), h.UpdateOfficial)
		profile.POST("/officials/delete", utils.AuthMiddleware(), h.DeleteOfficial)
	}
}
