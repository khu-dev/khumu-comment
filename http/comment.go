package http

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type CommentRouter struct {
	*echo.Group
	UC usecase.CommentUseCaseInterface
}

type LikeCommentRouter struct{
	*echo.Group
	UC usecase.LikeCommentUseCaseInterface
}

func NewCommentRouter(root *RootRouter, uc usecase.CommentUseCaseInterface) *CommentRouter {
	group := root.Group.Group("/comments")
	router := &CommentRouter{group, uc}
	group.POST("", router.Create)
	group.GET("", router.List)
	group.GET("/:id", router.Get)
	return router
}

func NewLikeCommentRouter(root *RootRouter, uc usecase.LikeCommentUseCaseInterface) *LikeCommentRouter {
	group := root.Group.Group("/like-comments")
	router := &LikeCommentRouter{group, uc}
	group.POST("", router.Create)
	return router
}

type CommentResponse struct {
	StatusCode int         `json:"statusCode"`
	Comments   interface{} `json:"comments,omitempty"` //this contains any format of comments
	Comment interface{} `json:"comment,omitempty"`
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
	if err != nil {return err}

	return c.JSON(200, comment)
}

func (r *LikeCommentRouter) Create(c echo.Context) error {
	log.Println("LikeCommentRouter_Create")
	var likeComment *model.LikeComment = &model.LikeComment{}
	err := c.Bind(likeComment)
	if err != nil {return c.JSON(http.StatusBadRequest, err.Error())}

	newLike, err := r.UC.Create(likeComment)
	if err != nil {return c.JSON(http.StatusBadRequest, err.Error())}

	return c.JSON(200, newLike)
}