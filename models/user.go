package models

import "gorm.io/gorm"

type Users struct { //database
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(100);uniqueIndex" json:"username"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password  string         `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type RegisterInput struct { //req API
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct { //req API
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct { //data tanpa pw
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
