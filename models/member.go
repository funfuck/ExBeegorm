package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
)

type Member struct {
	Profile   Prof `gorm:"ForeignKey:ProfileID;"`
	ProfileID uint
	Email []Email
	Language []Language `gorm:"many2many:user_languages;"`
	gorm.Model
}

func GetAllMember() []Member {
	db, _ := gorm.Open("mysql", beego.AppConfig.String("mysql::url"))
	defer db.Close()

	user := []Member{}
	db.Debug().Preload("Profile").Find(&user)
	return user
}

func AddMember() {
	db, _ := gorm.Open("mysql", beego.AppConfig.String("mysql::url"))
	defer db.Close()

	user := Member{
		Profile:Prof{Name:"yyy"},
		Email:[]Email{{Email:"gorm1@mail.com"}, {Email:"gorm2@mail.com"}},
	}
	db.Create(&user)
}