package contact

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
	var payload AddContactPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleContactError(c, err, "Failed to create contact")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Contact created successfully", nil)
}

func (h *Handler) Detail(c *gin.Context) {
	contact, err := h.service.Detail(c.Request.Context())
	if err != nil {
		handleContactError(c, err, "Failed to get contact")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Contact retrieved successfully", contact)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditContactPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid contact ID", nil)
		return
	}

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handleContactError(c, err, "Failed to update contact")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Contact updated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload ContactPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid contact ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleContactError(c, err, "Failed to delete contact")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Contact deleted successfully", nil)
}

func handleContactError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrContactNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Contact not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
