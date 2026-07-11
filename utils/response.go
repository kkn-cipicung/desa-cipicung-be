package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrInactiveUser       = errors.New("user is inactive")
	ErrUsernameTaken      = errors.New("username already taken")
	ErrInvalidPayload     = errors.New("invalid request payload")
)

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

func ErrorResponseJSON(c *gin.Context, statusCode int, message string, err error) {
	var errStr string
	if err != nil && statusCode < http.StatusInternalServerError {
		errStr = err.Error()
	}
	c.JSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
		Error:   errStr,
	})
}

func AuthErrorResponse(c *gin.Context, err error, message string) {
	switch {
	case errors.Is(err, ErrInvalidAuthPayload):
		ErrorResponseJSON(c, http.StatusBadRequest, "Invalid request payload", err)
	case errors.Is(err, ErrUsernameTaken):
		ErrorResponseJSON(c, http.StatusConflict, "Username already taken", err)
	case errors.Is(err, ErrInvalidCredentials):
		ErrorResponseJSON(c, http.StatusUnauthorized, "Invalid username or password", err)
	case errors.Is(err, ErrInactiveUser):
		ErrorResponseJSON(c, http.StatusForbidden, "User is inactive", err)
	default:
		ErrorResponseJSON(c, http.StatusInternalServerError, message, err)
	}
}
