package utils

import (
	"net/http"

	"cipicung.id/be/utils/token"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey   = "user_id"
	ContextUsernameKey = "username"
	ContextUserRoleKey = "user_role"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(token.AccessTokenCookieName)
		if err != nil {
			authHeader := c.GetHeader("Authorization")
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				accessToken = authHeader[7:]
			} else {
				ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
				c.Abort()
				return
			}
		}

		claims, err := token.ParseAccessToken(accessToken)
		if err != nil {
			ErrorResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextUsernameKey, claims.Username)
		c.Set(ContextUserRoleKey, claims.Role)
		c.Next()
	}
}

func UserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get(ContextUserIDKey)
	if !exists {
		return 0, false
	}

	switch value := userID.(type) {
	case uint:
		return value, value != 0
	case int:
		return uint(value), value > 0
	case int64:
		return uint(value), value > 0
	case float64:
		return uint(value), value > 0
	default:
		return 0, false
	}
}
