package http

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

func NewCommentRouter(root *RootRouter, uc usecase.CommentUseCaseInterface) *CommentRouter {
	group := root.Group.Group("/comments")
	commentRouter := &CommentRouter{group, uc}
	group.POST("", commentRouter.Create)
	group.GET("", commentRouter.List)
	group.GET("/:id", commentRouter.Get)
	return commentRouter
}

type CommentResponse struct {
	StatusCode int         `json:"statusCode"`
	Comments   interface{} `json:"comments,omitempty"` //this contains any format of comments
	Comment interface{} `json:"comment,omitempty"`
}

type CommentRouter struct {
	*echo.Group
	UC usecase.CommentUseCaseInterface
}

func (r *CommentRouter) Create(c echo.Context) error {
	log.Println("CommentRouter_Create")
	// 먼저 빈 Comment를 생성하고 거기에 값을 대입받는다.러 그렇지 않으면 nil 참조 에
	var comment *model.Comment = &model.Comment{}
	err := c.Bind(comment)
	if err != nil{
		log.Print(err)
		return err
	}
	authorUsername := comment.Author.Username

	if c.Get("user_id").(string) != authorUsername{
		return c.JSON(401, "Unauthorized error. The author has not been set as you.")
	}

	comment, err = r.UC.Create(comment)
	if err != nil{
		log.Print(err)
		return err
	}

	return c.JSON(200, CommentResponse{200, nil,comment})
}

func (r *CommentRouter) List(c echo.Context) error {
	log.Println("CommentRouter_List")
	username := c.Get("user_id").(string)
	if username == "" {log.Println("No user_id in context")}
	comments, err := r.UC.List(username, &repository.CommentQueryOption{})
	if err != nil {return err}

	return c.JSON(200, CommentResponse{200, comments, nil})
}

func (r *CommentRouter) Get(c echo.Context) error {
	log.Println("CommentRouter Get")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {return err}

	comment, err := r.UC.Get(id, &repository.CommentQueryOption{})
	//return c.JSON(200, model.String(comments[0]))
	return c.JSON(200, comment)
}
