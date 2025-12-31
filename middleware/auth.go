package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthRequired(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token tidak ditemukan",
		})
	}

	tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

	//menyiapkan claims
	claims := jwt.MapClaims{}

	//parse dan validasi token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRETKEY")), nil
	})
	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	//ambil claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	//ambil user_id dari claims
	userID := int(claims["user_id"].(float64))

	ctx.Locals("user_id", userID)

	return ctx.Next()
}
