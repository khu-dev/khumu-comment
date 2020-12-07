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

var (
	Location *time.Location
)

func init(){
	l, err := time.LoadLocation("Asia/Seoul")
	if err != nil{
		log.Fatal(err)
	} else{
		Location = l
	}

}

func NewGorm() *gorm.DB{
	var db *gorm.DB
	var connectionError error
	if config.Config.DB.Kind == "sqlite3"{
		log.Println("Connecting DB to " + config.Config.DB.SQLite3.FilePath)
		db, connectionError = gorm.Open(sqlite.Open(config.Config.DB.SQLite3.FilePath), &gorm.Config{})
	} else if config.Config.DB.Kind == "mysql"{
		dest := config.Config.DB.MySQL.Host + strconv.Itoa(config.Config.DB.MySQL.Port)
		log.Println("Connecting DB to " + dest)
		// loc을 통해 timezone을 설정해줘야만함.
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
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil{log.Panic(err)}

	err = MigrateAll(db)
	if err != nil{log.Panic(err)}
	// this allows foreign key contraints
	db.Exec("PRAGMA foreign_keys=ON")
	return db
}

func MigrateAll(db *gorm.DB) error{
	err := db.AutoMigrate(&model.Board{}, &model.KhumuUser{}, &model.Article{}, &model.Comment{}, &model.LikeComment{})
	return err
}

func MigrateMinimum(db *gorm.DB) error{
	err := db.AutoMigrate(&model.Comment{}, &model.LikeComment{})
	return err
}