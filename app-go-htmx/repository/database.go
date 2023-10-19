package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	database, err := gorm.Open(sqlite.Open("../repository/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Student{}, &Session{})

	db = database
}
