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

func NewCommentUseCase(r repository.CommentRepositoryInterface) CommentUseCaseInterface {
	return &CommentUseCase{Repository: r}
}

func (uc *CommentUseCase) List(c echo.Context) []*model.Comment {

	log.Println("CommentUseCase List")
	comments := uc.Repository.List(&repository.CommentQueryOption{})
	parents := uc.listParentWithChildren(comments)
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return nil
	}

	for _, p := range parents {
		for _, c := range p.Children {
			uc.hideAuthor(c, userID)
		}
		uc.hideAuthor(p, userID)
	}

	return parents
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment {
	var parents []*model.Comment

	for _, comment := range allComments {
		if comment.ParentID == nil {
			parents = append(parents, comment)
		}
	}

	return parents
}

func (uc *CommentUseCase) hideAuthor(c *model.Comment, requestUsername string) {
	if c.Kind == "anonymous" && c.AuthorUsername != requestUsername {
		c.Author.Username = "익명"
	}
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
