package models

import "github.com/jinzhu/gorm"

type Email struct {
	Email   string
	UserID  uint
	Member Member
	gorm.Model
}
