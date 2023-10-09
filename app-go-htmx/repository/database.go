package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	Id        uint32 `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}

var db *gorm.DB

func init() {
	database, err := gorm.Open(sqlite.Open("../repository/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Student{})

	db = database
}
