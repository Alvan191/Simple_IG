package handlers

import (
	"time"

	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetContent(ctx *fiber.Ctx) error {
	var getContent []models.Insta
	if err := config.DB.
		Preload("Coments.User"). //agar komentar dan username muncul
		Preload("User").         // agar username post muncul
		Order("created_at DESC").
		Find(&getContent).Error; err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	for i := range getContent {
		getContent[i].CreatedAt = getContent[i].CreatedAt.In(loc)
	}

	return ctx.Render("home", fiber.Map{
		"Posts": getContent,
	})
}

func PostContent(ctx *fiber.Ctx) error {
	var postContent models.Insta
	if err := ctx.BodyParser(&postContent); err != nil {
		return ctx.Status(400).SendString("Isi postingan tidak valid")
	}

	// err := validate.Struct(&postContent)
	// if err != nil {
	// 	return ctx.Status(400).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	//ambil user_id dari middleware
	userID := ctx.Locals("user_id").(int)
	postContent.UserID = uint(userID)

	config.DB.Create(&postContent)
	return ctx.Redirect("/")
}
