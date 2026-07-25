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

func (h *Handler) CreateOfficial(c *gin.Context) {
	var payload AddOfficialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.CreateOfficial(c.Request.Context(), payload); err != nil {
		handleProfileError(c, err, "Failed to create official")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Official created successfully", nil)
}

func (h *Handler) ListOfficials(c *gin.Context) {
	var payload ListOfficialPayload
	_ = c.ShouldBindJSON(&payload)

	officials, err := h.service.ListOfficials(c.Request.Context(), payload)
	if err != nil {
		handleProfileError(c, err, "Failed to list officials")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Officials retrieved successfully", officials)
}

func (h *Handler) DetailOfficial(c *gin.Context) {
	var payload OfficialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid official ID", nil)
		return
	}

	official, err := h.service.FindOfficialByID(c.Request.Context(), payload)
	if err != nil {
		handleProfileError(c, err, "Failed to get official detail")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Official detail retrieved successfully", official)
}

func (h *Handler) UpdateOfficial(c *gin.Context) {
	var payload EditOfficialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid official ID", nil)
		return
	}

	if err := h.service.UpdateOfficial(c.Request.Context(), payload); err != nil {
		handleProfileError(c, err, "Failed to update official")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Official updated successfully", nil)
}

func (h *Handler) DeleteOfficial(c *gin.Context) {
	var payload OfficialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid official ID", nil)
		return
	}

	if err := h.service.DeleteOfficial(c.Request.Context(), payload); err != nil {
		handleProfileError(c, err, "Failed to delete official")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Official deleted successfully", nil)
}

func handleProfileError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrProfileNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Profile not found", err)
		return
	}
	if errors.Is(err, ErrOfficialNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Official not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
