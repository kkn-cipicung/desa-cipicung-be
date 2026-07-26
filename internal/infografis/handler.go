package infografis

import (
	"net/http"

	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Detail(c *gin.Context) {
	infografis, err := h.service.Detail(c.Request.Context())
	if err != nil {
		utils.ErrorResponseJSON(c, http.StatusInternalServerError, "Failed to get infografis", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Infografis retrieved successfully", infografis)
}
