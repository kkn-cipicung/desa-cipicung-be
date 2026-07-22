package profile

import (
	"errors"
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

func (h *Handler) Create(c *gin.Context) {
	var payload AddProfilePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleProfileError(c, err, "Failed to save profile")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile saved successfully", nil)
}

func (h *Handler) Detail(c *gin.Context) {
	profile, err := h.service.Detail(c.Request.Context())
	if err != nil {
		handleProfileError(c, err, "Failed to get profile")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", profile)
}

func (h *Handler) FindRegionBoundary(c *gin.Context) {
	profile, err := h.service.FindRegionBoundary(c.Request.Context())
	if err != nil {
		handleProfileError(c, err, "Failed to get profile region boundary")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile region boundary retrieved successfully", profile)
}

func (h *Handler) FindVisionMission(c *gin.Context) {
	profile, err := h.service.FindVisionMission(c.Request.Context())
	if err != nil {
		handleProfileError(c, err, "Failed to get profile vision mission")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile vision mission retrieved successfully", profile)
}

func (h *Handler) FindGovernmentStructure(c *gin.Context) {
	structure, err := h.service.FindGovernmentStructure(c.Request.Context())
	if err != nil {
		handleProfileError(c, err, "Failed to get government structure")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Government structure retrieved successfully", structure)
}

func (h *Handler) FindResourcePotential(c *gin.Context) {
	potentials, err := h.service.FindResourcePotential(c.Request.Context())
	if err != nil {
		handleProfileError(c, err, "Failed to get resource potential")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Resource potential retrieved successfully", potentials)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload ProfilePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid profile ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleProfileError(c, err, "Failed to delete profile")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile deleted successfully", nil)
}

func handleProfileError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrProfileNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Profile not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
