package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

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
	c.JSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
		Error:   PublicErrorDetail(statusCode, err),
	})
}

func ErrorResponseJSONWithDetail(c *gin.Context, statusCode int, message string, err error) {
	var detail string
	if err != nil {
		detail = err.Error()
	}
	c.JSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
		Error:   detail,
	})
}

func PublicErrorDetail(statusCode int, err error) string {
	if err == nil {
		return ""
	}

	var unmarshalTypeErr *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeErr) {
		field := unmarshalTypeErr.Field
		if field == "" {
			field = "request body"
		}
		return fmt.Sprintf("%s must be %s", field, jsonTypeName(unmarshalTypeErr.Type))
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return "request body contains invalid JSON"
	}

	errMessage := err.Error()
	switch {
	case strings.Contains(errMessage, "json: cannot unmarshal"):
		return "request body contains invalid field type"
	case strings.Contains(errMessage, "invalid character"):
		return "request body contains invalid JSON"
	case errors.Is(err, ErrInvalidPayload):
		return strings.TrimPrefix(errMessage, ErrInvalidPayload.Error()+": ")
	default:
		return errMessage
	}
}

func jsonTypeName(t reflect.Type) string {
	if t == nil {
		return "the expected type"
	}
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "a number"
	case reflect.Bool:
		return "a boolean"
	case reflect.String:
		return "a string"
	case reflect.Slice, reflect.Array:
		return "an array"
	case reflect.Map, reflect.Struct:
		return "an object"
	default:
		return "the expected type"
	}
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
