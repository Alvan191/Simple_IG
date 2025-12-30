package router

import (
	"github.com/Alvan191/Simple_IG.git/handlers"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {
	// app.Static("/", "./views")
	app.Get("/insta", handlers.GetContent)
	app.Post("/insta", handlers.PostContent)
}
