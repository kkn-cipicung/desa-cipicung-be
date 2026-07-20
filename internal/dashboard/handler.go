package dashboard

import (
	"errors"
	"io"
	"log"
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
	var payload AddDashboardPayload
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
		handleDashboardError(c, err, "Failed to create dashboard")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Dashboard created successfully", nil)
}

func (h *Handler) List(c *gin.Context) {
	var payload ListDashboardPayload
	if err := c.ShouldBindJSON(&payload); err != nil && !errors.Is(err, io.EOF) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	dashboards, err := h.service.List(c.Request.Context(), payload)
	if err != nil {
		handleDashboardError(c, err, "Failed to get dashboards")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboards retrieved successfully", dashboards)
}

func (h *Handler) Detail(c *gin.Context) {
	dashboard, err := h.service.Detail(c.Request.Context())
	if err != nil {
		handleDashboardError(c, err, "Failed to get dashboard")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard retrieved successfully", dashboard)
}

func (h *Handler) FindActive(c *gin.Context) {
	dashboard, err := h.service.FindActive(c.Request.Context())
	if err != nil {
		handleDashboardError(c, err, "Failed to get active dashboard")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Active dashboard retrieved successfully", dashboard)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditDashboardPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid dashboard ID", nil)
		return
	}

	userID, ok := utils.UserIDFromContext(c)
	if !ok {
		utils.ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	payload.UpdatedBy = userID

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		log.Printf("dashboard update failed: id=%d user_id=%d error=%v", payload.ID, userID, err)
		handleDashboardError(c, err, "Failed to update dashboard")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard updated successfully", nil)
}

func (h *Handler) Activate(c *gin.Context) {
	var payload DashboardPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Activate(c.Request.Context(), payload); err != nil {
		handleDashboardError(c, err, "Failed to activate dashboard")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard activated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload DashboardPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid dashboard ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleDashboardError(c, err, "Failed to delete dashboard")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard deleted successfully", nil)
}

func handleDashboardError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrDashboardNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Dashboard not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
