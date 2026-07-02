package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Message: message,
		Data:    data,
	})
}

func ErrorResponseJSON(c *gin.Context, statusCode int, message string, err error) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	c.JSON(statusCode, ErrorResponse{
		Message: message,
		Error:   errStr,
	})
}
