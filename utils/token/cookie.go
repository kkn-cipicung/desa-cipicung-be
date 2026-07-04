package token

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetAuthCookies(c *gin.Context, pair *Pair) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(AccessTokenCookieName, pair.AccessToken, AccessTokenMaxAge(), "/", "", secureCookie(), true)
	c.SetCookie(RefreshTokenCookieName, pair.RefreshToken, RefreshTokenMaxAge(), "/", "", secureCookie(), true)
}

func ClearAuthCookies(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(AccessTokenCookieName, "", -1, "/", "", secureCookie(), true)
	c.SetCookie(RefreshTokenCookieName, "", -1, "/", "", secureCookie(), true)
}

func secureCookie() bool {
	return strings.EqualFold(strings.TrimSpace(os.Getenv("COOKIE_SECURE")), "true")
}
