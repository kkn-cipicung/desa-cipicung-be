package business

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
	var payload AddBusinessPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleBusinessError(c, err, "Failed to create business")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Business created successfully", nil)
}

func (h *Handler) List(c *gin.Context) {
	var payload ListBusinessPayload
	if err := c.ShouldBindJSON(&payload); err != nil && !errors.Is(err, io.EOF) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	businesses, err := h.service.List(c.Request.Context(), payload)
	if err != nil {
		handleBusinessError(c, err, "Failed to get businesses")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Businesses retrieved successfully", businesses)
}

func (h *Handler) FindByID(c *gin.Context) {
	var payload BusinessPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid business ID", nil)
		return
	}

	business, err := h.service.FindByID(c.Request.Context(), payload)
	if err != nil {
		handleBusinessError(c, err, "Failed to get business")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Business retrieved successfully", business)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditBusinessPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid business ID", nil)
		return
	}

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handleBusinessError(c, err, "Failed to update business")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Business updated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload BusinessPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid business ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleBusinessError(c, err, "Failed to delete business")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Business deleted successfully", nil)
}

func handleBusinessError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrBusinessNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Business not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
