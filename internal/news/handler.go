package news

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"cipicung.id/be/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var payload AddNewsPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	userID, ok := utils.UserIDFromContext(c)
	if !ok {
		utils.ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	payload.UploadedBy = userID

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleNewsError(c, err, "Failed to create news")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "News created successfully", nil)
}

func (h *Handler) List(c *gin.Context) {
	var payload ListNewsPayload
	if err := c.ShouldBindJSON(&payload); err != nil && !errors.Is(err, io.EOF) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	newsList, err := h.service.List(c.Request.Context(), payload)
	if err != nil {
		handleNewsError(c, err, "Failed to get news")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "News retrieved successfully", newsList)
}

func (h *Handler) FindByID(c *gin.Context) {
	var payload NewsByIdPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid news ID", nil)
		return
	}

	newsItem, err := h.service.FindByID(c.Request.Context(), payload)
	if err != nil {
		handleNewsError(c, err, "Failed to get news")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "News retrieved successfully", newsItem)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditNewsPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid news ID", nil)
		return
	}

	userID, ok := utils.UserIDFromContext(c)
	if !ok {
		utils.ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	payload.UploadedBy = userID

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handleNewsError(c, err, "Failed to update news")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "News updated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload NewsPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid news ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleNewsError(c, err, "Failed to delete news")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "News deleted successfully", nil)
}

func (h *Handler) FindByDate(c *gin.Context) {
	var payload NewsByDatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.Date == "" {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid date", nil)
		return
	}

	parsedDate, err := utils.ParseDate(payload.Date)
	if err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid date format, expected YYYY-MM-DD", err)
		return
	}

	payload.Date = parsedDate

	newsList, err := h.service.FindByDate(c.Request.Context(), payload)
	if err != nil {
		handleNewsError(c, err, "Failed to get news by date")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "News retrieved successfully", newsList)
}

func (h *Handler) FindHeader(c *gin.Context) {
	var payload NewsByIdPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid news ID", nil)
		return
	}

	header, err := h.service.FindHeader(c.Request.Context(), payload)
	if err != nil {
		handleNewsError(c, err, "Failed to get news header")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "News header retrieved successfully", header)
}

func handleNewsError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrNewsNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "News not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23503":
			utils.ErrorResponseJSON(c, http.StatusBadRequest, "Referenced data does not exist", databaseErrorDetail(err, pqErr))
			return
		case "23502", "22001", "22P02":
			utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid news data", databaseErrorDetail(err, pqErr))
			return
		case "23505":
			utils.ErrorResponseJSON(c, http.StatusConflict, "News data already exists", databaseErrorDetail(err, pqErr))
			return
		}
	}

	utils.ErrorResponseJSONWithDetail(c, http.StatusInternalServerError, message, err)
}

func databaseErrorDetail(err error, pqErr *pq.Error) error {
	detail := pqErr.Detail
	if detail == "" {
		detail = pqErr.Message
	}
	return fmt.Errorf("%v: %s", err, detail)
}
