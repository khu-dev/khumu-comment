package repository

import (
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

func NewGorm() *gorm.DB{
	var db *gorm.DB
	var connectionError error
	if config.Config.DB.Kind == "sqlite3"{
		log.Println("Connecting DB to " + config.Config.DB.SQLite3.FilePath)
		db, connectionError = gorm.Open(sqlite.Open(config.Config.DB.SQLite3.FilePath), &gorm.Config{})
	} else if config.Config.DB.Kind == "mysql"{
		dest := config.Config.DB.MySQL.Host + strconv.Itoa(config.Config.DB.MySQL.Port)
		log.Println("Connecting DB to " + dest)
		db, connectionError = gorm.Open(mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
			config.Config.DB.MySQL.User,
			config.Config.DB.MySQL.Password,
			config.Config.DB.MySQL.Host,
			config.Config.DB.MySQL.Port,
			config.Config.DB.MySQL.DatabaseName,
		)), &gorm.Config{})
	}
	if connectionError != nil {
		log.Print("Failed to connect database. Retry after 3 seconds")
		time.Sleep(time.Duration(3000) * time.Millisecond)
		return NewGorm()
	}else{
		return db
	}

}

func NewTestGorm() *gorm.DB{
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil{log.Panic(err)}

	err = db.AutoMigrate(&model.Board{}, &model.KhumuUser{}, &model.Article{}, &model.Comment{}, &model.LikeComment{})
	if err != nil{log.Panic(err)}

	return db
}