package models

type Users struct { //database
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(100);uniqueIndex" json:"username"`
	Email    string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password string `json:"-"`
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
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
