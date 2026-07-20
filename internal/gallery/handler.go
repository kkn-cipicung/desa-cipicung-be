package gallery

import (
	"errors"
	"io"
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
	var payload AddGalleryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	userID, ok := utils.UserIDFromContext(c)
	if !ok {
		utils.ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	payload.CreatedBy = userID

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleGalleryError(c, err, "Failed to create gallery")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Gallery created successfully", nil)
}

func (h *Handler) List(c *gin.Context) {
	var payload ListGalleryPayload
	if err := c.ShouldBindJSON(&payload); err != nil && !errors.Is(err, io.EOF) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	galleries, err := h.service.List(c.Request.Context(), payload)
	if err != nil {
		handleGalleryError(c, err, "Failed to get galleries")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Galleries retrieved successfully", galleries)
}

func (h *Handler) FindByID(c *gin.Context) {
	var payload GalleryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid gallery ID", nil)
		return
	}

	gallery, err := h.service.FindByID(c.Request.Context(), payload)
	if err != nil {
		handleGalleryError(c, err, "Failed to get gallery")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Gallery retrieved successfully", gallery)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditGalleryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid gallery ID", nil)
		return
	}

	userID, ok := utils.UserIDFromContext(c)
	if !ok {
		utils.ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	payload.UpdatedBy = userID

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handleGalleryError(c, err, "Failed to update gallery")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Gallery updated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload GalleryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid gallery ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleGalleryError(c, err, "Failed to delete gallery")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Gallery deleted successfully", nil)
}

func handleGalleryError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrInvalidGalleryCategory) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid gallery category", err)
		return
	}
	if errors.Is(err, ErrGalleryNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Gallery not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
