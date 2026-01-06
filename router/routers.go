package router

import (
	"github.com/Alvan191/Simple_IG.git/handlers/api"
	handlers "github.com/Alvan191/Simple_IG.git/handlers/web"
	"github.com/Alvan191/Simple_IG.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(app *fiber.App) {
	//ROUTER WEBSITE REGIST and LOGIN
	app.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.Render("login", nil)
	})
	app.Get("/register", func(ctx *fiber.Ctx) error {
		return ctx.Render("register", nil)
	})
	app.Post("/users/regist", handlers.Register)
	app.Post("/users/login", handlers.Login)

	//ROUTER WEBSITE post content
	app.Get("/", middleware.AuthRequiredWeb, handlers.GetContent)
	app.Post("/insta", middleware.AuthRequiredWeb, handlers.PostContent)
	app.Get("/insta/:id/edit", middleware.AuthRequiredWeb, handlers.EditContent)
	app.Post("/insta/:id/update", middleware.AuthRequiredWeb, handlers.UpdateContent)
	app.Post("/insta/:id/delete", middleware.AuthRequiredWeb, handlers.DeleteContent)

	//ROUTER WEBSITE comment
	app.Post("/posts/:post_id/comments", middleware.AuthRequiredWeb, handlers.CreateComment)
	app.Get("/posts/:post_id/comments", handlers.GetCommentsByPost)

	//ROUTER API REGIST and LOGIN
	app.Post("/users/regist_api", api.RegistAPI)
	app.Post("/users/login_api", api.LoginAPI)

	//ROUTER API post content
	app.Post("/insta/postcontent_api", middleware.AuthRequiredAPI, api.PostContentAPI)
	app.Get("insta/getcontent_api", api.GetContentAPI)

	//ROUTER API comment
	app.Post("/posts/:post_id/postcomments_api", middleware.AuthRequiredAPI, api.CreateCommentAPI)
	app.Get("/posts/:post_id/getcomments_api", middleware.AuthRequiredAPI, api.GetCommentsByPostAPI)
}
