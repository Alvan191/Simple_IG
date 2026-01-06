package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthRequiredWeb(ctx *fiber.Ctx) error {
	tokenStr := ctx.Cookies("jwt")
	if tokenStr == "" {
		return ctx.Redirect("/login")
	}

	//menyiapkan claims
	claims := jwt.MapClaims{}

	//parse dan validasi token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRETKEY")), nil
	})
	if err != nil || !token.Valid {
		return ctx.Redirect("/login")
	}
	userID := int(claims["user_id"].(float64))

	ctx.Locals("user_id", userID)

	return ctx.Next()
}

func AuthRequiredAPI(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token tidak ditemukan",
		})
	}

	tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRETKEY")), nil
	})
	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token tidak valid",
		})
	}

	userID := int(claims["user_id"].(float64))

	ctx.Locals("user_id", userID)

	return ctx.Next()
}
