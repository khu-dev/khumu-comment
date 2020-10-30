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
	comments := uc.Repository.List(&repository.CommentQueryOption{})
	parents := uc.listParentWithChildren(comments)

	// Do Something
	//for _, parent := range parents{
	//	for _, child := range parent.Children{
	//
	//	}
	//}
	return parents
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment{
	var parents []*model.Comment

	for _, comment := range allComments{
		if comment.ParentID == 0{
			parents = append(parents, comment)
		}
	}

	return parents
}

func (uc *CommentUseCase) Get(c echo.Context) *model.Comment {
	log.Println("CommentUseCase Get")
	comment := uc.Repository.Get(c.Param("id"))
	return comment
}
