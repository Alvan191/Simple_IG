package api

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func PostContentAPI(ctx *fiber.Ctx) error {
	var postContent models.Insta
	if err := ctx.BodyParser(&postContent); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := validate.Struct(&postContent)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := ctx.Locals("user_id").(int)
	postContent.UserID = uint(userID)

	config.DB.Create(&postContent)

	resp := models.InstaResponse{
		ID:        postContent.ID,
		UserID:    postContent.UserID,
		Content:   postContent.Content,
		CreatedAt: postContent.CreatedAt,
	}

	return ctx.Status(201).JSON(resp)
}

func GetContentAPI(ctx *fiber.Ctx) error {
	var getContent []models.Insta
	result := config.DB.Preload("Coments").Find(&getContent)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "tidak menemukan content",
		})
	}

	var response []models.InstaResponse

	for _, content := range getContent {
		response = append(response, models.InstaResponse{
			ID:           content.ID,
			UserID:       content.UserID,
			Content:      content.Content,
			CreatedAt:    content.CreatedAt,
			UpdatedAt:    content.UpdatedAt,
			CommentCount: int64(len(content.Coments)),
		})
	}

	return ctx.JSON(response)
}
