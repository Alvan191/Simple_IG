package migration

import (
	"github.com/Alvan191/Simple_IG.git/config"
	"github.com/Alvan191/Simple_IG.git/models"
)

// func MigrateUsersInsta() {
// 	db := config.DB

// 	if !db.Migrator().HasTable(&models.Users{}) {
// 		db.Migrator().CreateTable(&models.Users{})
// 	}

// 	if !db.Migrator().HasColumn(&models.Users{}, "Username") {
// 		db.Migrator().AddColumn(&models.Users{}, "Username")
// 	}

// 	if !db.Migrator().HasColumn(&models.Users{}, "Email") {
// 		db.Migrator().AddColumn(&models.Users{}, "Email")
// 	}

// 	if !db.Migrator().HasColumn(&models.Users{}, "Password") {
// 		db.Migrator().AddColumn(&models.Users{}, "Password")
// 	}
// }

// func MigrateUserIDContent() {
// 	db := config.DB

// 	if !db.Migrator().HasColumn(&models.Insta{}, "user_id") {
// 		db.Migrator().AddColumn(&models.Insta{}, "user_id")
// 	}
// }

// func MigrateCommentContentInsta() {
// 	db := config.DB
// 	if db == nil {
// 		log.Fatal("DB is nil")
// 	}

// 	err := db.AutoMigrate(&models.Comments{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Comment migration done")
// }

func MigrateSimpleIG() {
	db := config.DB

	db.AutoMigrate(
		&models.Users{},
		&models.Insta{},
		&models.Comments{},
	)
}
