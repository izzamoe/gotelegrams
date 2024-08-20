package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ChatID     string `json:"chatid"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Berapakali Berapakali
}

// tablename use pointer
func (*User) TableName() string {
	return "users"
}
