package http

import (
	"errors"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CommentRouter struct {
	*echo.Group
	commentUC usecase.CommentUseCaseInterface
	likeUC    usecase.LikeCommentUseCaseInterface
}

func NewCommentRouter(root *RootRouter, commentUC usecase.CommentUseCaseInterface, likeUC usecase.LikeCommentUseCaseInterface) *CommentRouter {
	group := root.Group.Group("/comments")
	commentRouter := &CommentRouter{group, commentUC, likeUC}

	// comment는 Update API를 제공하지 않는다.
	group.POST("", commentRouter.Create)
	group.GET("", commentRouter.List)
	group.GET("/:id", commentRouter.Get)
	group.PATCH("/:id", commentRouter.Update)
	group.DELETE("/:id", commentRouter.Delete)
	group.PATCH("/:id/likes", commentRouter.LikeToggle)

	return commentRouter
}


type CommentResponse struct {
	Data    *data.CommentOutput `json:"data"` //this contains any format of comments
	Message string       `json:"message"`
}

type CommentsResponse struct {
	Data    []*data.CommentOutput `json:"data"` //this contains any format of comments
	Message string           `json:"message"`
}

type LikeCommentResponse struct {
	Data    bool   `json:"data"` //this contains any format of comments
	Message string `json:"message"`
}

func (r *CommentRouter) Create(c echo.Context) error {
	logrus.Debug("CommentRouter_Create")
	// 먼저 빈 Comment를 생성하고 거기에 값을 대입받는다. 그렇지 않으면 nil 참조 에러
	var commentInput *data.CommentInput = &data.CommentInput{}
	err := c.Bind(commentInput)

	if err != nil {
		logrus.Error(err)
		return c.JSON(400, CommentResponse{Data: nil, Message: err.Error()})
	}
	commentInput.Author = c.Get("user_id").(string)
	comment, err := r.commentUC.Create(commentInput)
	if err != nil {
		logrus.Error(err)
		return c.JSON(400, CommentResponse{Data: nil, Message: err.Error()})
	}

	return c.JSON(200, CommentResponse{Data: comment})
}

func (r *CommentRouter) List(c echo.Context) error {
	logrus.Debug("CommentRouter_List")
	username := c.Get("user_id").(string)
	if username == "" {
		return c.JSON(403, CommentResponse{Message: "No user_id in context"})
	}
	if !isAdmin(username) {
		logrus.Println(c.QueryParams())
		if c.QueryParam("article") == "" {
			//return c.JSON(400, CommentResponse{StatusCode: 401, Message: ""})
			return c.JSON(http.StatusBadRequest, CommentResponse{Message: "article in query string is required"})
		}
	}

	opt := &repository.CommentQueryOption{}
	articleIDString := c.QueryParam("article")
	if articleIDString == "" {
		articleIDString = "0"
	}
	articleID, err := strconv.Atoi(articleIDString)
	//if articleID == {articleID=0}
	if err != nil {
		logrus.WithField("article", articleIDString).Error(err)
		return c.JSON(400, CommentResponse{Message: "article should be int"})
	}
	opt.ArticleID = articleID

	commentIDString := c.QueryParam("comment")

	if commentIDString == "" {
		commentIDString = "0"
	}
	commentID, err := strconv.Atoi(commentIDString)

	if err != nil {
		logrus.Println(err, commentIDString, commentID)
		return err
	}
	opt.CommentID = commentID

	comments, err := r.commentUC.List(username, opt)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(200, CommentsResponse{Data: comments})
}

func (r *CommentRouter) Get(c echo.Context) error {
	logrus.Debug("CommentRouter Get")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		return err
	}

	comment, err := r.commentUC.Get(c.Get("user_id").(string), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, CommentResponse{Data: nil, Message: "No comment with id=" + strconv.Itoa(id)})
		}

		logrus.Error(err)
		return err
	}

	return c.JSON(200, comment)
}

func (r *CommentRouter) Update(c echo.Context) error {
	logrus.Debug("CommentRouter_Update")
	// 먼저 빈 Comment를 생성하고 거기에 값을 대입받는다. 그렇지 않으면 nil 참조 에러
	body := make(map[string]interface{})
	err := c.Bind(&body)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		return err
	}

	username := c.Get("user_id").(string)
	updated, err := r.commentUC.Update(username, id, body)
	if err != nil {
		logrus.Error(err)
		return c.JSON(400, CommentResponse{Data: nil, Message: err.Error()})
	}

	return c.JSON(200, CommentResponse{Data: updated})
}

func (r *CommentRouter) Delete(c echo.Context) error {
	logrus.Debug("CommentRouter Delete")
	username := c.Get("user_id").(string)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = r.commentUC.Delete(username, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, CommentResponse{Data: nil, Message: "No comment with id=" + strconv.Itoa(id)})
		}

		logrus.Error(err)
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// @Tags Like Comment
// @Summary Comment에 대한 "좋아요"를 생성하거나 삭제합니다.
// @Description 현재 좋아요 상태이면 삭제, 좋아요 상태가 아니면 생성합니다.
// @name create-like-comment
// @Produce  application/json
// @Param comment body int true "좋아요할 comment의 ID"
// @Router /api/like-comments [put]
// @Success 200 {object} CommentResponse
func (r *CommentRouter) LikeToggle(c echo.Context) error {
	logrus.Debug("LikeCommentRouter_Toggle")
	var likeComment *model.LikeComment = &model.LikeComment{Comment: &model.Comment{}, User: &model.KhumuUserSimple{}}
	username := c.Get("user_id").(string)
	likeComment.Username = username

	commentID, err := strconv.Atoi(c.Param("id"))
	logrus.Warn(commentID)
	if err != nil {
		logrus.Error("Wrong comment ID format")
		return c.JSON(http.StatusBadRequest, LikeCommentResponse{Message: "comment 필드가 올바른 int 값이 아닙니다."})
	}
	likeComment.CommentID = commentID

	isCreated, err := r.likeUC.Toggle(likeComment)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, LikeCommentResponse{Message: err.Error()})
	}

	if isCreated {
		return c.JSON(201, LikeCommentResponse{Data: isCreated})
	} else {
		return c.NoContent(204)
	}

}
