package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // import database's driver
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
