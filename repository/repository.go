package repository

import (
	"github.com/khu-dev/khumu-comment/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func NewGorm() *gorm.DB{
	log.Println("Connecting DB to " + config.Config.DB.SQLite3.FilePath)
	db, err := gorm.Open(sqlite.Open(config.Config.DB.SQLite3.FilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
