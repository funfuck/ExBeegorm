package models

import "github.com/jinzhu/gorm"

type Prof struct {
	Name string
	gorm.Model
}
