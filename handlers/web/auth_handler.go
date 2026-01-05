package handlers

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/Alvan191/Simple_IG.git/utils"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	var input models.RegisterInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).SendString("Input tidak valid")
	}

	hash, _ := utils.HashPassword(input.Password)

	var exist models.Users
	result := config.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&exist)
	if result.Error == nil {
		return ctx.Status(409).SendString("Username atau Email sudah digunakan")
	}

	user := models.Users{
		Username: input.Username,
		Email:    input.Email,
		Password: hash,
	}

	config.DB.Create(&user)

	//Code Untuk API
	// return ctx.JSON(fiber.Map{
	// 	"success": true,
	// 	"message": "Registrasi Berhasil",
	// })

	return ctx.Redirect("/login", fiber.StatusSeeOther) //code untuk redirect ke halaman .html lain
}

func Login(ctx *fiber.Ctx) error {
	var input models.LoginInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).SendString("Input tidak valid")
	}

	var user models.Users
	config.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 {
		return ctx.Render("login", fiber.Map{
			"Error": "Email tidak ditemukan",
		})
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		return ctx.Render("login", fiber.Map{
			"Error": "Password salah",
		})
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		return ctx.Render("login", fiber.Map{
			"Error": "Gagal generate token",
		})
	}

	// return ctx.JSON(fiber.Map{
	// 	"success": true,
	// 	"token":   token,
	// 	"user": models.UserResponse{
	// 		ID:       int(user.ID),
	// 		Username: user.Username,
	// 		Email:    user.Email,
	// 	},
	// })
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		Path:     "/",
	})

	return ctx.Redirect("/", fiber.StatusSeeOther)
}
