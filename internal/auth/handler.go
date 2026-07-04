package auth

import (
	"net/http"

	"cipicung.id/be/utils"
	tokenutils "cipicung.id/be/utils/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var payload RegisterPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.service.Register(c.Request.Context(), payload); err != nil {
		utils.AuthErrorResponse(c, err, "Failed to register user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", nil)
}

func (h *Handler) Login(c *gin.Context) {
	var payload LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	tokens, err := h.service.Login(c.Request.Context(), payload)
	if err != nil {
		utils.AuthErrorResponse(c, err, "Failed to login user")
		return
	}

	tokenutils.SetAuthCookies(c, tokens)
	utils.SuccessResponse(c, http.StatusOK, "User logged in successfully", gin.H{"access_token": tokens.AccessToken})
}
