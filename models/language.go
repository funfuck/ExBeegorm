package models

import "github.com/jinzhu/gorm"

type Language struct {
	Name string
	gorm.Model
}