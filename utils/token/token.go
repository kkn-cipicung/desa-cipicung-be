package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	AccessTokenCookieName  = "access_token"
	RefreshTokenCookieName = "refresh_token"

	AccessTokenType  = "access"
	RefreshTokenType = "refresh"

	defaultIssuer            = "cipicung.id"
	defaultSecret            = "cipicung-dev-secret"
	defaultAccessExpiration  = 15 * time.Minute
	defaultRefreshExpiration = 7 * 24 * time.Hour
)

var (
	ErrInvalidToken   = errors.New("invalid token")
	ErrExpiredToken   = errors.New("expired token")
	ErrInvalidSigning = errors.New("invalid token signing method")
	ErrMissingSecret  = errors.New("jwt secret is required")
)

type Claims struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	Issuer    string `json:"iss"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

type Pair struct {
	AccessToken  string
	RefreshToken string
}

type jwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

func GenerateAccessToken(userID uint, username, role string) (string, error) {
	return generateToken(userID, username, role, AccessTokenType, defaultAccessExpiration)
}

func GenerateRefreshToken(userID uint, username, role string) (string, error) {
	return generateToken(userID, username, role, RefreshTokenType, defaultRefreshExpiration)
}

func GeneratePair(userID uint, username, role string) (*Pair, error) {
	accessToken, err := GenerateAccessToken(userID, username, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken(userID, username, role)
	if err != nil {
		return nil, err
	}

	return &Pair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func AccessTokenMaxAge() int {
	return int(defaultAccessExpiration.Seconds())
}

func RefreshTokenMaxAge() int {
	return int(defaultRefreshExpiration.Seconds())
}

func generateToken(userID uint, username, role, tokenType string, expiration time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:    userID,
		Username:  username,
		Role:      role,
		TokenType: tokenType,
		Issuer:    defaultIssuer,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(expiration).Unix(),
	}

	jwtSecret, err := secret()
	if err != nil {
		return "", err
	}

	return generate(claims, jwtSecret)
}

func ParseAccessToken(accessToken string) (*Claims, error) {
	return parse(accessToken, AccessTokenType)
}

func ParseRefreshToken(refreshToken string) (*Claims, error) {
	return parse(refreshToken, RefreshTokenType)
}

func parse(value, expectedTokenType string) (*Claims, error) {
	parts := strings.Split(value, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidToken
	}

	jwtSecret, err := secret()
	if err != nil {
		return nil, err
	}

	signingInput := parts[0] + "." + parts[1]
	expectedSignature := sign(signingInput, jwtSecret)
	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return nil, ErrInvalidToken
	}

	var header jwtHeader
	if err := decode(parts[0], &header); err != nil {
		return nil, err
	}

	if header.Algorithm != "HS256" || header.Type != "JWT" {
		return nil, ErrInvalidSigning
	}

	var claims Claims
	if err := decode(parts[1], &claims); err != nil {
		return nil, err
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return nil, ErrExpiredToken
	}

	if claims.TokenType != expectedTokenType {
		return nil, ErrInvalidToken
	}

	return &claims, nil
}

func generate(claims Claims, secret string) (string, error) {
	header := jwtHeader{
		Algorithm: "HS256",
		Type:      "JWT",
	}

	encodedHeader, err := encode(header)
	if err != nil {
		return "", err
	}

	encodedClaims, err := encode(claims)
	if err != nil {
		return "", err
	}

	signingInput := encodedHeader + "." + encodedClaims
	signature := sign(signingInput, secret)

	return fmt.Sprintf("%s.%s", signingInput, signature), nil
}

func encode(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(data), nil
}

func decode(value string, target any) error {
	data, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

func sign(value, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(value))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func secret() (string, error) {
	value := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if value == "" {
		if isProduction() {
			return "", ErrMissingSecret
		}
		return defaultSecret, nil
	}

	return value, nil
}

func isProduction() bool {
	for _, key := range []string{"APP_ENV", "GO_ENV", "GIN_MODE"} {
		value := strings.TrimSpace(os.Getenv(key))
		if strings.EqualFold(value, "production") || strings.EqualFold(value, "release") {
			return true
		}
	}

	return false
}
