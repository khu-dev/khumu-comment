package http

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
	group.PUT("", router.Toggle)
	return router
}

type CommentResponse struct {
	Data   *model.Comment `json:"data"` //this contains any format of comments
	Message string `json:"message"`
}

type CommentsResponse struct {
	Data   []*model.Comment `json:"data"` //this contains any format of comments
	Message string `json:"message"`
}

type LikeCommentResponse struct {
	Data   bool `json:"data"` //this contains any format of comments
	Message string `json:"message"`
}

// @Tags Comment
// @Summary Comment를 생성합니다.
// @Description 사용 가능한 필드는 주로 Get API의 응답에 있는 필드와 유사합니다.
// @Description author field는 요청자의 Authorization header의 값을 이용합니다.
// @name create-comment
// @Accept  application/json
// @Produce  application/json
// @Param article body int true "어떤 게시물의 댓글인지"
// @Param kind body string false "익명인지, 기명인지"
// @Param content body string true "댓글 내용"
// @Router /api/comments [post]
// @Success 200 {object} CommentResponse
func (r *CommentRouter) Create(c echo.Context) error {
	logrus.Debug("CommentRouter_Create")
	// 먼저 빈 Comment를 생성하고 거기에 값을 대입받는다.러 그렇지 않으면 nil 참조 에
	var comment *model.Comment = &model.Comment{Author: &model.KhumuUserSimple{}}
	err := c.Bind(comment)

	if err != nil{
		log.Print(err)
		return err
	}

	comment.AuthorUsername = c.Get("user_id").(string)
	comment, err = r.UC.Create(comment)
	if err != nil{
		log.Print(err)
		return err
	}

	return c.JSON(200, CommentResponse{Data: comment})
}

// @Tags comment
// @Summary Comment List를 조회합니다.
// @Description
// @name list-comment
// @Produce  application/json
// @Param article query int true "admin group이 아닌 이상은 게시물 id를 꼭 정의해야합니다."
// @Router /api/comments [get]
// @Success 200 {object} CommentsResponse
func (r *CommentRouter) List(c echo.Context) error {
	logrus.Debug("CommentRouter_List")
	username := c.Get("user_id").(string)
	if username == "" {
		return c.JSON(403, CommentResponse{Message: "No user_id in context"})
	}
	if !isAdmin(username){
		log.Println(c.QueryParams())
		if c.QueryParam("article") == ""{
			//return c.JSON(400, CommentResponse{StatusCode: 401, Message: ""})
			return c.JSON(http.StatusBadRequest, CommentResponse{Message: "article in query string is required"})
		}
	}

	opt := &repository.CommentQueryOption{}
	articleIDString := c.QueryParam("article")
	if articleIDString == ""{articleIDString="0"}
	articleID, err := strconv.Atoi(articleIDString)
	//if articleID == {articleID=0}
	if err!=nil{
		logrus.WithField("article", articleIDString).Error(err)
		return c.JSON(400, CommentResponse{Message: "article should be int"})
	}
	opt.ArticleID = articleID

	commentIDString := c.QueryParam("comment")

	if commentIDString == ""{commentIDString="0"}
	commentID, err := strconv.Atoi(commentIDString)

	if err != nil {
		logrus.Println(err, commentIDString, commentID)
		return err
	}
	opt.CommentID = commentID

	comments, err := r.UC.List(username, opt)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(200, CommentsResponse{Data: comments})
}

// @Tags Comment
// @Summary Comment 조회합니다.
// @Description
// @name get-comment
// @Produce  application/json
// @Param id path int true "Comment ID"
// @Router /api/comments/{id} [get]
// @Success 200 {object} CommentResponse
func (r *CommentRouter) Get(c echo.Context) error {
	logrus.Debug("CommentRouter Get")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		return err
	}

	comment, err := r.UC.Get(id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(200, comment)
}

// @Tags Like Comment
// @Summary Comment에 대한 "좋아요"를 생성하거나 삭제합니다.
// @Description 현재 좋아요 상태이면 삭제, 좋아요 상태가 아니면 생성합니다.
// @name create-like-comment
// @Produce  application/json
// @Param comment body int true "좋아요할 comment의 ID"
// @Router /api/like-comments [put]
// @Success 200 {object} CommentResponse
func (r *LikeCommentRouter) Toggle(c echo.Context) error {
	logrus.Debug("LikeCommentRouter_Toggle")
	var likeComment *model.LikeComment = &model.LikeComment{Comment:&model.Comment{}, User:&model.KhumuUserSimple{}}
	username := c.Get("user_id").(string)
	body := &echo.Map{}
	err := c.Bind(body)

	likeComment.Username = username
	if err != nil{
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, LikeCommentResponse{Message: err.Error()})
	}

	commentIDFloat64, ok := (*body)["comment"].(float64)
	commentID := int(commentIDFloat64)
	if !ok{
		logrus.Error("Wrong comment ID format")
		return c.JSON(http.StatusBadRequest, LikeCommentResponse{Message: "comment 필드가 올바른 int 값이 아닙니다."})
	}
	likeComment.CommentID =  commentID

	isCreated, err := r.UC.Toggle(likeComment)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, LikeCommentResponse{Message: err.Error()})
	}

	if isCreated{
		return c.JSON(201, LikeCommentResponse{Data: isCreated})
	} else{
		return c.NoContent(204)
	}

}
