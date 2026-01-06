package api

import (
	"strconv"

	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCommentAPI(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	postID, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Post ID tidak valid",
		})
	}

	var insta models.Insta
	if err := config.DB.First(&insta, postID).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Postingan tidak tersedia",
		})
	}

	var input struct {
		Content string `json:"Content"`
	}
	if err := ctx.BodyParser(&input); err != nil || input.Content == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Content wajib di isi",
		})
	}

	comment := models.Comments{
		Content: input.Content,
		UserID:  uint(userID),
		PostID:  uint(postID),
	}

	config.DB.Create(&comment).Preload("User").First(&comment, comment.ID)

	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    comment,
	})
}

func GetCommentsByPostAPI(ctx *fiber.Ctx) error {
	postID := ctx.Params("post_id")

	var comments []models.Comments

	err := config.DB.
		Preload("User").
		Where("post_id = ?", postID).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Gagal mengambil komentar",
		})
	}

	return ctx.JSON(comments)
}
