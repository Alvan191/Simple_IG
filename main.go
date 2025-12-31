package main

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/migration"
	"github.com/Alvan191/Simple_IG.git/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	migration.MigrateUsersInsta()
	migration.MigrateUserIDContent()

	app := fiber.New()

	router.MapRoutes(app)

	app.Listen(":8080")
}
