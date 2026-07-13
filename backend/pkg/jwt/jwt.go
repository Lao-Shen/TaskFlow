package authjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ============================================================
//  JWT
// ============================================================

func GenerateToken(id uint64, username string) (string, error) {
	claims := &Claims{
		UserID:   id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "TaskFlow",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}
