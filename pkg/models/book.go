package models

import (
	"book-crud-project/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Title       string `gorm:""json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
