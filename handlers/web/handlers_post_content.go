package handlers

import (
	"fmt"
	"strconv"
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
		Preload("Coments", "deleted_at IS NULL").
		Preload("Coments.User").  //mengambil user pemilik komentar
		Preload("User").          //mengambil user pemilik post
		Order("created_at DESC"). //mengurutkan post berdasarkan waktu dibuat dari terbaru ke terlama
		Find(&getContent).Error; err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	for i := range getContent {
		// post time
		getContent[i].CreatedAt = getContent[i].CreatedAt.In(loc)

		// comment time
		for j := range getContent[i].Coments {
			getContent[i].Coments[j].CreatedAt =
				getContent[i].Coments[j].CreatedAt.In(loc)
		}
	}

	userID := ctx.Locals("user_id").(int)

	var user models.Users
	config.DB.First(&user, userID)

	return ctx.Render("home", fiber.Map{
		"Posts":         getContent,
		"CurrentUserID": uint(userID),
		"CurrentUser":   user.Username,
	})
}

func UpdateContent(ctx *fiber.Ctx) error {
	now := time.Now().UTC()
	id := ctx.Params("id")
	var updateContent models.Insta
	if err := ctx.BodyParser(&updateContent); err != nil {
		return ctx.Status(400).SendString("Bad request")
	}

	userID := ctx.Locals("user_id").(int)

	result := config.DB.
		Model(&models.Insta{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"content":    updateContent.Content,
			"updated_at": now,
		})

	if result.RowsAffected == 0 {
		return ctx.Status(403).SendString("Tidak punya izin mengedit postingan ini")
	}

	return ctx.Redirect("/")
}

func EditContent(ctx *fiber.Ctx) error { // ini untuk memunculkan data lama ketika klik edit
	id := ctx.Params("id")
	userID := ctx.Locals("user_id").(int)

	var post models.Insta
	result := config.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&post)

	if result.RowsAffected == 0 {
		return ctx.Status(403).SendString("Tidak punya izin mengedit postingan ini")
	}

	return ctx.Render("update", fiber.Map{
		"Post": post,
	})
}

func DeleteContent(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	postId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return ctx.Status(400).SendString("Invalid ID")
	}

	userID := ctx.Locals("user_id").(int)

	result := config.DB.
		Where("id = ? AND user_id = ?", postId, userID).
		Delete(&models.Insta{})
	if result.Error != nil {
		return ctx.Status(500).SendString("Failed to delete content")
	}

	if result.RowsAffected == 0 {
		return ctx.Status(403).SendString("Tidak punya izin menghapus postingan ini")
	}

	fmt.Println("Rows affected:", result.RowsAffected)
	return ctx.Redirect("/")
}
