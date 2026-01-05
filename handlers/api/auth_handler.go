package api

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/Alvan191/Simple_IG.git/utils"
	"github.com/gofiber/fiber/v2"
)

func RegistAPI(ctx *fiber.Ctx) error {
	var input models.RegisterInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Input tidak valid",
		})
	}

	hash, _ := utils.HashPassword(input.Password)

	var exist models.Users
	result := config.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&exist)
	if result.Error == nil {
		return ctx.Status(409).JSON(fiber.Map{
			"error": "Username atau Email telah digunakan",
		})
	}

	user := models.Users{
		Username: input.Username,
		Email:    input.Email,
		Password: hash,
	}

	config.DB.Create(&user)

	return ctx.JSON(fiber.Map{
		"succes":  true,
		"message": "Registrasi berhasil",
	})
}

func LoginAPI(ctx *fiber.Ctx) error {
	var input models.LoginInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Input tidak valid",
		})
	}

	var user models.Users
	config.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Email tidak ditemukan",
		})
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Password salah",
		})
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Gagal generate token",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"token":   token,
		"user": models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}
