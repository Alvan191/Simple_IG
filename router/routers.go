package router

import (
	"github.com/Alvan191/Simple_IG.git/handlers"
	"github.com/Alvan191/Simple_IG.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {
	// app.Static("/", "./views")
	app.Get("/insta", handlers.GetContent)
	app.Post("/insta", middleware.AuthRequired, handlers.PostContent)

	app.Post("/users/regist", handlers.Register)
	app.Post("/users/login", handlers.Login)

	//route comment
	app.Post("/posts/:post_id/comments", middleware.AuthRequired, handlers.CreateComment)
	app.Get("/posts/:post_id/comments", handlers.GetCommentsByPost)
}
