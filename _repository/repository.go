package _repository

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/sirupsen/logrus"
	"time"
)

func NewEnt() *ent.Client {
	// parseTime=true가 없을 시
	// Error: unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	// ref: https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	drv, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			config.Config.DB.MySQL.User,
			config.Config.DB.MySQL.Password,
			config.Config.DB.MySQL.Host,
			config.Config.DB.MySQL.DatabaseName,
		))
	if err != nil {
		logrus.Panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	ent.Debug()
	ent.Log(func(i ...interface{}) {
		logrus.Warn(i...)
	})
	client := ent.NewClient(ent.Driver(drv))
	return client
}

// Deprecated. 이제 Gorm이 아니라 facebook/ent를 사용한다.

//func NewGorm() *gorm.DB {
//	var db *gorm.DB
//	var connectionError error
//	if config.Config.DB.Kind == "sqlite3" {
//		logrus.Println("Connecting DB to " + config.Config.DB.SQLite3.FilePath)
//		db, connectionError = gorm.Open(sqlite.Open(config.Config.DB.SQLite3.FilePath), &gorm.Config{})
//	} else if config.Config.DB.Kind == "mysql" {
//		dest := config.Config.DB.MySQL.Host + strconv.Itoa(config.Config.DB.MySQL.Port)
//		log.Println("Connecting DB to " + dest)
//		// loc을 통해 timezone을 설정해줘야만함.
//		db, connectionError = gorm.Open(mysql.Open(fmt.Sprintf(
//			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
//			config.Config.DB.MySQL.User,
//			config.Config.DB.MySQL.Password,
//			config.Config.DB.MySQL.Host,
//			config.Config.DB.MySQL.Port,
//			config.Config.DB.MySQL.DatabaseName,
//		)), &gorm.Config{})
//	}
//	if connectionError != nil {
//		log.Print("Failed to connect database. Retry after 3 seconds")
//		time.Sleep(time.Duration(3000) * time.Millisecond)
//		return NewGorm()
//	} else {
//		return db
//	}
//}
//
//
//
//func NewTestGorm() *gorm.DB {
//	var db *gorm.DB
//	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
//	if err != nil {
//		log.Panic(err)
//	}
//	// this allows foreign key contraints
//	// 켤까 싶었는데, 켜면 테스트 할 때 유저랑 게시판이랑 다 만들어야함 ㅜㅜ
//	// db.Exec("PRAGMA foreign_keys=ON")
//	return db
//}
