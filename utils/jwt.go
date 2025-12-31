package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtKey := []byte(os.Getenv("JWT_SECRETKEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
