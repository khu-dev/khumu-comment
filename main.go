package main

import (
  "github.com/khu-dev/khumu-comment/config"
  "github.com/khu-dev/khumu-comment/repository"
  "github.com/khu-dev/khumu-comment/usecase"
  "github.com/khu-dev/khumu-comment/http"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "log"
)



func main(){
  Run()
}

func Run(){
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
  e.Logger.Fatal(e.Start(config.Config.Host+":"+config.Config.Port))
}

func tmp(){
  // 연습 참고


  // Migrate the schema
  //db.AutoMigrate(&Product{})

  // Create
  //db.Create(&Product{Code: "D42", Price: 100})

  // Read
  //var user models.KhumuUser
  //var article models.Article
  //var board models.Board
  //var users []*models.KhumuUser
  //var comments []*models.Comment
  //db.Find(&users, "username='Park'") // find product with integer primary key
  ////fmt.Printf("%#v\n", user)
  //db.First(&article)
  ////fmt.Println(article)
  //db.Find(&board, fmt.Sprintf("id=%d", article.BoardID))
  ////fmt.Println(board)
  //models.PrintModel(&user)
  ////fmt.Println(users[0])
  //db.Find(&comments)
  //models.PrintModel(comments[0])


  //db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  //db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  //db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  //db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  //db.Delete(&product, 1)
}