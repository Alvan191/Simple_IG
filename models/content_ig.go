package models

import "time"

//nama type harus sesuai nama tabel database nanti eror gorm tidak mengenali class struct jika tidak sama dengan nama table, misal type Content struct maka => "Error 1146 (42S02): Table 'simple_ig.content' doesn't exist" harusnya type Insta struct => "simple_ig.insta"
type Insta struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content" validate:"required"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
