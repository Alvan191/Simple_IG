package models

import "time"

type Comments struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	User Users `gorm:"foreignKey:UserID"`
}
