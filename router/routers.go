package router

import (
	"github.com/Alvan191/Simple_IG.git/handlers"
	"github.com/Alvan191/Simple_IG.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {
	// app.Static("/", "./views")
	app.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.Render("login", nil)
	})
	app.Get("/register", func(ctx *fiber.Ctx) error {
		return ctx.Render("register", nil)
	})
	// app.Get("/home", middleware.AuthRequired, func(ctx *fiber.Ctx) error {
	// 	return ctx.Render("home", nil)
	// })

	app.Get("/", middleware.AuthRequired, handlers.GetContent)
	app.Post("/insta", middleware.AuthRequired, handlers.PostContent)
	app.Get("/insta/:id/edit", middleware.AuthRequired, handlers.EditContent)
	app.Post("/insta/:id/update", middleware.AuthRequired, handlers.UpdateContent)
	app.Post("/insta/:id/delete", handlers.DeleteContent)

	app.Post("/users/regist", handlers.Register)
	app.Post("/users/login", handlers.Login)

	//route comment
	app.Post("/posts/:post_id/comments", middleware.AuthRequired, handlers.CreateComment)
	app.Get("/posts/:post_id/comments", handlers.GetCommentsByPost)
}
