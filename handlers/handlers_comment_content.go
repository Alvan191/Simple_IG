package handlers

import (
	"strconv"
	"time"

	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/gofiber/fiber/v2"
)

func CreateComment(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int) // mengmbil user dari jwt
	postID, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Post ID tidak valid",
		})
	}

	var input struct {
		Content string `json:"content"`
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

	var post models.Insta
	if err := config.DB.First(&post, postID).Error; err != nil {
		return ctx.Status(404).SendString("Postingan tidak ditemukan")
	}

	config.DB.Create(&comment).Preload("User").First(&comment, comment.ID)

	// return ctx.JSON(fiber.Map{
	// 	"success": true,
	// 	"data":    comment,
	// })
	return ctx.Redirect("/")
}

func GetCommentsByPost(ctx *fiber.Ctx) error {
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

	loc, _ := time.LoadLocation("Asia/Jakarta")
	for i := range comments {
		comments[i].CreatedAt = comments[i].CreatedAt.In(loc)
	}

	// return ctx.JSON(comments)
	return ctx.Render("home", fiber.Map{
		"Coments": comments,
	})
}
