package main

import (
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/http"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	Run()
}

func Run() {
	log.Println("Connecting DB to " + config.Config.DB.SQLite3.FilePath)
	db, err := gorm.Open(sqlite.Open(config.Config.DB.SQLite3.FilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userRepository := &repository.UserRepositoryGorm{DB: db}
	commentRepository := &repository.CommentRepositoryGorm{DB: db}
	commentUC := &usecase.CommentUseCase{
		Repository: commentRepository,
	}
	e := http.NewEcho(userRepository, commentUC)

	e.Logger.Print("Started Server")
	e.Logger.Fatal(e.Start(config.Config.Host + ":" + config.Config.Port))
}
