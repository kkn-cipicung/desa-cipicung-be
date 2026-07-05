package potential

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
	var payload AddPotentialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusInternalServerError, "Failed to create potential", err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Potential created successfully", nil)
}

func (h *Handler) List(c *gin.Context) {
	var payload ListPotentialPayload
	if err := c.ShouldBindJSON(&payload); err != nil && !errors.Is(err, io.EOF) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	potentialList, err := h.service.List(c.Request.Context(), payload)
	if err != nil {
		utils.ErrorResponseJSON(c, http.StatusInternalServerError, "Failed to get potentials", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Potentials retrieved successfully", potentialList)
}

func (h *Handler) FindByID(c *gin.Context) {
	var payload PotentialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid potential ID", nil)
		return
	}

	potentialItem, err := h.service.FindByID(c.Request.Context(), payload)
	if err != nil {
		handlePotentialError(c, err, "Failed to get potential")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Potential retrieved successfully", potentialItem)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditPotentialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid potential ID", nil)
		return
	}

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handlePotentialError(c, err, "Failed to update potential")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Potential updated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload PotentialPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid potential ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handlePotentialError(c, err, "Failed to delete potential")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Potential deleted successfully", nil)
}

func handlePotentialError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrPotentialNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Potential not found", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
