package http

import (
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"log"
)

func NewCommentRouter(parent *echo.Group, uc usecase.CommentUseCaseInterface) *CommentRouter {
	group := parent.Group("/comments")
	commentRouter := &CommentRouter{group, uc}
	group.GET("", commentRouter.ListComment)
	group.GET("/:id", commentRouter.GetComment)
	return commentRouter
}

type CommentResponse struct {
	StatusCode int         `json:"statusCode"`
	Comments   interface{} `json:"comments"` //this contains any format of comments
}

type CommentRouter struct {
	*echo.Group
	UC usecase.CommentUseCaseInterface
}

func (r *CommentRouter) ListComment(c echo.Context) error {
	log.Println("CommentRouter List")

	comments := r.UC.List(c)
	//return c.JSON(200, model.String(comments[0]))
	return c.JSON(200, CommentResponse{200, comments})
}

func (r *CommentRouter) GetComment(c echo.Context) error {
	log.Println("CommentRouter Get")

	comment := r.UC.Get(c)
	//return c.JSON(200, model.String(comments[0]))
	return c.JSON(200, comment)
}
