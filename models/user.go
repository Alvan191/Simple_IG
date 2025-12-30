package models

type Users struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(100);unique" json:"username"`
	Email    string `gorm:"type:varchar(100);unique" json:"email"`
	Password string `json:"-"`
}
