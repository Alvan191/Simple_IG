package main

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/migration"
	"github.com/Alvan191/Simple_IG.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.ConnectDB()
	// migration.MigrateUsersInsta()
	// migration.MigrateUserIDContent()
	// migration.MigrateCommentContentInsta()
	migration.MigrateSimpleIG()

	engine := html.New("./views", ".html") // tambahan untuk kehubung ke html

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.MapRoutes(app)

	app.Listen(":8080")
}
