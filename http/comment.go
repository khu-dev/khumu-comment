package http

import (
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"log"
)

func NewCommentRouter(root *RootRouter, uc usecase.CommentUseCaseInterface) *CommentRouter {
	group := root.Group.Group("/comments")
	commentRouter := &CommentRouter{group, uc}
	group.GET("/", commentRouter.List)
	group.GET("/:id", commentRouter.Get)
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

func (r *CommentRouter) List(c echo.Context) error {
	log.Println("CommentRouter List")

	comments := r.UC.List(c)
	//return c.JSON(200, model.String(comments[0]))
	return c.JSON(200, CommentResponse{200, comments})
}

func (r *CommentRouter) Get(c echo.Context) error {
	log.Println("CommentRouter Get")

	comment := r.UC.Get(c)
	//return c.JSON(200, model.String(comments[0]))
	return c.JSON(200, comment)
}
