package mapdata

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
	var payload AddMapPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleMapError(c, err, "Failed to create map")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Map created successfully", nil)
}

func (h *Handler) Detail(c *gin.Context) {
	mapDetail, err := h.service.Detail(c.Request.Context())
	if err != nil {
		handleMapError(c, err, "Failed to get map")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Map retrieved successfully", mapDetail)
}

func (h *Handler) FindActive(c *gin.Context) {
	mapDetail, err := h.service.FindActive(c.Request.Context())
	if err != nil {
		handleMapError(c, err, "Failed to get active map")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Active map retrieved successfully", mapDetail)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditMapPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid map ID", nil)
		return
	}

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handleMapError(c, err, "Failed to update map")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Map updated successfully", nil)
}

func (h *Handler) Activate(c *gin.Context) {
	var payload MapPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Activate(c.Request.Context(), payload); err != nil {
		handleMapError(c, err, "Failed to activate map")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Map activated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload MapPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid map ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleMapError(c, err, "Failed to delete map")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Map deleted successfully", nil)
}

func handleMapError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrMapNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Map not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
