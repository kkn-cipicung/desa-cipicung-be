package category

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
	var payload AddCategoryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Create(c.Request.Context(), payload); err != nil {
		handleCategoryError(c, err, "Failed to create category")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Category created successfully", nil)
}

func (h *Handler) List(c *gin.Context) {
	var payload ListCategoryPayload
	if err := c.ShouldBindJSON(&payload); err != nil && !errors.Is(err, io.EOF) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	categories, err := h.service.List(c.Request.Context(), payload)
	if err != nil {
		handleCategoryError(c, err, "Failed to get categories")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

func (h *Handler) FindByID(c *gin.Context) {
	var payload CategoryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	category, err := h.service.FindByID(c.Request.Context(), payload)
	if err != nil {
		handleCategoryError(c, err, "Failed to get category")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", category)
}

func (h *Handler) Update(c *gin.Context) {
	var payload EditCategoryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	if err := h.service.Update(c.Request.Context(), payload); err != nil {
		handleCategoryError(c, err, "Failed to update category")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category updated successfully", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	var payload CategoryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if payload.ID == 0 {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), payload); err != nil {
		handleCategoryError(c, err, "Failed to delete category")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category deleted successfully", nil)
}

func handleCategoryError(c *gin.Context, err error, message string) {
	if errors.Is(err, ErrCategoryNotFound) {
		utils.ErrorResponseJSON(c, http.StatusNotFound, "Category not found", err)
		return
	}
	if errors.Is(err, utils.ErrInvalidPayload) {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	utils.ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
}
