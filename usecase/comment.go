package usecase

import (
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"log"
)
import "github.com/khu-dev/khumu-comment/model"

type CommentUseCase struct{
	Repository repository.CommentRepository
}

func (uc *CommentUseCase) List(c echo.Context) []*model.Comment {
	log.Println("CommentUseCase List")
	comments := uc.Repository.List(&repository.CommentQueryOption{ArticleID: 7})
	//for _, c := range comments{
	//	model.PrintModel(c)
	//}
	return comments
}

func (uc *CommentUseCase) Get(c echo.Context) *model.Comment {
	log.Println("CommentUseCase Get")
	comment := uc.Repository.Get(c.Param("id"))
	return comment
}
