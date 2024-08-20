package entity

import "gorm.io/gorm"

type Berapakali struct {
	gorm.Model
	BerapaX int  `json:"berapax"`
	UserID  uint `json:"userid"`
}

// tablename use pointer
func (*Berapakali) TableName() string {
	return "berapakali"
}
