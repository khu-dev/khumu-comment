package usecase

import (
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)
import "github.com/khu-dev/khumu-comment/model"

type CommentUseCaseInterface interface {
	List(c echo.Context) []*model.Comment
	Get(c echo.Context) *model.Comment
}

type CommentUseCase struct {
	Repository repository.CommentRepositoryInterface
}

func (uc *CommentUseCase) List(c echo.Context) []*model.Comment {
	log.Println("CommentUseCase List")
	comments := uc.Repository.List(&repository.CommentQueryOption{})
	parents := uc.listParentWithChildren(comments)

	return parents
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment {
	var parents []*model.Comment

	for _, comment := range allComments {
		if comment.ParentID == 0 {
			parents = append(parents, comment)
		}
	}

	return parents
}

func (uc *CommentUseCase) Get(c echo.Context) *model.Comment {
	log.Println("CommentUseCase Get")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	comment := uc.Repository.Get(id)
	return comment
}
