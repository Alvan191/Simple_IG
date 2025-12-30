package migration

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
)

func MigrateUsersInsta() {
	db := config.DB

	if !db.Migrator().HasTable(&models.Users{}) {
		db.Migrator().CreateTable(&models.Users{})
	}

	if !db.Migrator().HasColumn(&models.Users{}, "Username") {
		db.Migrator().AddColumn(&models.Users{}, "Username")
	}

	if !db.Migrator().HasColumn(&models.Users{}, "Email") {
		db.Migrator().AddColumn(&models.Users{}, "Email")
	}

	if !db.Migrator().HasColumn(&models.Users{}, "Password") {
		db.Migrator().AddColumn(&models.Users{}, "Password")
	}
}
