package handlers

import (
	"time"

	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
	"github.com/gofiber/fiber/v2"
)

// var validate = validator.New()

func PostContent(ctx *fiber.Ctx) error {
	var postContent models.Insta
	if err := ctx.BodyParser(&postContent); err != nil {
		return ctx.Status(400).SendString("Isi postingan tidak valid")
	}

	//ambil user_id dari middleware
	userID := ctx.Locals("user_id").(int)
	postContent.UserID = uint(userID)

	config.DB.Create(&postContent)
	return ctx.Redirect("/")
}

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

	userID := ctx.Locals("user_id").(int)

	return ctx.Render("home", fiber.Map{
		"Posts":         getContent,
		"CurrentUserID": uint(userID),
	})
}

func UpdateContent(ctx *fiber.Ctx) error {
	now := time.Now().UTC()
	id := ctx.Params("id")
	var updateContent models.Insta
	if err := ctx.BodyParser(&updateContent); err != nil {
		return ctx.Status(400).SendString("Bad request")
	}

	updates := map[string]interface{}{
		"content":    updateContent.Content,
		"updated_at": &now,
	}
	config.DB.Model(&models.Insta{}).Where("id = ?", id).Updates(updates)

	config.DB.First(&updateContent, id)
	return ctx.Redirect("/")
}

func EditContent(ctx *fiber.Ctx) error { // ini untuk memunculkan data lama ketika klik edit
	id := ctx.Params("id")

	var post models.Insta
	config.DB.First(&post, id)

	return ctx.Render("update", fiber.Map{
		"Post": post,
	})
}

func DeleteContent(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	result := config.DB.Delete(&models.Insta{}, id)
	if result.Error != nil {
		return ctx.Status(500).SendString("Failed to delete content")
	}

	if result.RowsAffected == 0 {
		return ctx.Status(404).SendString("content not found")
	}

	return ctx.Redirect("/")
}
