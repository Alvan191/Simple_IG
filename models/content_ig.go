package models

import (
	"time"

	"gorm.io/gorm"
)

// nama type harus sesuai nama tabel database nanti eror gorm tidak mengenali class struct jika tidak sama dengan nama table, misal type Content struct maka => "Error 1146 (42S02): Table 'simple_ig.content' doesn't exist" harusnya type Insta struct => "simple_ig.insta"
type Insta struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `json:"user_id"`

	Content   string         `json:"content" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt *time.Time     `gorm:"<-:update" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	User    Users      `gorm:"foreignKey:UserID"`
	Coments []Comments `gorm:"foreignKey:PostID"`
}

type InstaResponse struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`

	CommentCount int64
}
