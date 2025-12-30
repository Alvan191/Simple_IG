package handlers

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetContent(ctx *fiber.Ctx) error {
	var getContent []models.Insta
	config.DB.Find(&getContent)

	return ctx.JSON(getContent)
}

func PostContent(ctx *fiber.Ctx) error {
	var postContent models.Insta
	ctx.BodyParser(&postContent)

	err := validate.Struct(&postContent)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	config.DB.Create(&postContent)

	return ctx.JSON(postContent)
}
