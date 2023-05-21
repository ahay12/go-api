package model

import (
	"github.com/ahay12/go-api/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	database, errs := gorm.Open(mysql.Open("root:ahay@tcp(localhost:3306)/go_api"))
	helper.PanicIfError(errs)

	err := database.AutoMigrate(&Page{})
	if err != nil {
		return
	}

	DB = database
}
