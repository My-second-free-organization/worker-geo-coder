package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   string   `json:"sub"`
	Email    string   `json:"email"`
	TenantID string   `json:"tenant_id"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, email, tenantID string, roles []string, secret string) (string, error) {
	claims := Claims{
		UserID: userID, Email: email, TenantID: tenantID, Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour))},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func ValidateToken(tokenStr, secret string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	return claims, err
}
