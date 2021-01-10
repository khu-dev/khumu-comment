package http

import (
	"bytes"
	"encoding/json"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var (
	commentEcho              *echo.Echo
	likeCommentEcho *echo.Echo
	commentRouter  *CommentRouter
	commentUseCase usecase.CommentUseCaseInterface
	likeCommentUseCase usecase.LikeCommentUseCaseInterface
)

// 후에 mocking을 사용하게 된다면 이 타입을 이용
type CommentUseCaseMock struct{}
//func (uc *CommentUseCaseMock) List(c echo.Context) []*model.Comment {
//	return commentsMock
//}
//
//func (uc *CommentUseCaseMock) Get(c echo.Context) *model.Comment {
//	return commentsMock[0]
//}

func TestSetUp(t *testing.T) {
	cont := dig.New()
	err := cont.Provide(repository.NewTestGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewLikeCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewUserRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(usecase.NewCommentUseCase)
	assert.Nil(t, err)

	err = cont.Provide(usecase.NewLikeCommentUseCase)
	assert.Nil(t, err)

	err = cont.Invoke(func(commentUC usecase.CommentUseCaseInterface, likeCommentUC usecase.LikeCommentUseCaseInterface) {
		commentEcho = echo.New()
		mockRoot := RootRouter{commentEcho.Group("/comments")}
		commentRouter = NewCommentRouter(&mockRoot, commentUC, likeCommentUC)
		commentUseCase = commentUC
		likeCommentUseCase = likeCommentUC
	})

	t.Run("Create sample users to preload in list comment", func(t *testing.T) {
		for _, user := range test.UsersData{
			username := user.Username
			t.Log("Create a user named ", username)
			err = cont.Invoke(func(db *gorm.DB){
				dbErr := db.Create(&user).Error
				assert.Nil(t, dbErr)
				assert.Equal(t, username, user.Username)
			})
			assert.Nil(t, err)
		}
	})

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(test.CommentsData), 3)
}

func TestCommentRouter_Create(t *testing.T){
	defaultComment := map[string]interface{}{
		"kind":"anonymous",
		"author":map[string]interface{}{
			"username": "jinsu",
		},
		"article": 1,
		"content": "jinsu의 익명 댓글을 테스트하고 있습니다.\nhello, world",
		"parent": nil,
	}

	t.Run("Authenticated user", func(t *testing.T) {
		data, err := json.Marshal(defaultComment)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodGet, "/comments", bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		assert.NotNil(t, commentRouter.commentUC)
		err = commentRouter.Create(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	//Unauthenticated user 혹은 자신의 username이 아닌 author username에 대한 test
}

func TestCommentRouter_List(t *testing.T) {
	t.Run("Admin user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "admin")
		err := commentRouter.List(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		body, _ := ioutil.ReadAll(rec.Body)
		log.Println("BODY", string(body))
	})

	t.Run("Jinsu user with required arguments", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments?article=1", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		err := commentRouter.List(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		body, _ := ioutil.ReadAll(rec.Body)
		log.Println("BODY", string(body))
	})

	t.Run("Jinsu user without required arguments", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		err := commentRouter.List(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, _ := ioutil.ReadAll(rec.Body)
		log.Println("BODY", string(body))
	})

}

func TestCommentRouter_Get(t *testing.T) {
	t.Run("Get a non-existing comment", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		context.SetParamNames("id")
		context.SetParamValues("99999")
		//con
		err := commentRouter.Get(context)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestCommentRouter_Delete(t *testing.T) {
	var (
		tmpParentComment, tmpChildComment *model.Comment
		err error
	)
	t.Run("[Setup] 임시 코멘트 생성", func(t *testing.T) {
		tmpParentComment, err = commentUseCase.Create(&model.Comment{ArticleID: 1, AuthorUsername: "jinsu", Author: &model.KhumuUserSimple{Username: "jinsu"}, ParentID: nil, Content: "삭제될 임시 Parent Comment"})
		assert.NoError(t, err)
		assert.NotNil(t, tmpParentComment)
		tmpChildComment, err = commentUseCase.Create(&model.Comment{ArticleID: 1, AuthorUsername: "somebody", Author: &model.KhumuUserSimple{Username: "somebody"}, ParentID: &tmpParentComment.ID, Content: "삭제될 임시 Child Comment"})
		assert.NoError(t, err)
		assert.NotNil(t, tmpChildComment)
	})

	t.Run("부모 댓글 삭제", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments/" + strconv.Itoa(tmpParentComment.ID), nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.SetParamNames("id")
		context.SetParamValues(strconv.Itoa(tmpParentComment.ID))
		//con
		err := commentRouter.Delete(context)
		assert.Nil(t, err)

		deletedStateParentComment, err := commentUseCase.Get(tmpParentComment.ID)
		assert.Equal(t, "deleted", deletedStateParentComment.State)

		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("자식 댓글 삭제", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments/" + strconv.Itoa(tmpChildComment.ID), nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.SetParamNames("id")
		context.SetParamValues(strconv.Itoa(tmpChildComment.ID))
		//con
		err := commentRouter.Delete(context)
		assert.Nil(t, err)

		deletedStateChildComment, err := commentUseCase.Get(tmpChildComment.ID)
		assert.Equal(t, "deleted", deletedStateChildComment.State)

		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}
//func TestLikeCommentRouter_Get(t *testing.T) {
//	req := httptest.NewRequest(http.MethodGet, "/", nil)
//	rec := httptest.NewRecorder()
//
//	context := commentEcho.NewContext(req, rec)
//	context.Set("user_id", "admin")
//	err := commentRouter.List(context)
//	assert.Nil(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//	body, _ := ioutil.ReadAll(rec.Body)
//	log.Println("BODY", string(body))
//}

func TestLikeCommentRouter_Toggle(t *testing.T) {
	t.Run("Somebody likes jinsu's comment 1.", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/comments/1/likes", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "somebody")
		context.SetParamNames("id")
		context.SetParamValues("1")
		assert.NotNil(t, commentRouter.likeUC)
		err := commentRouter.LikeToggle(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("Somebody doesn't like jinsu's comment 1.", func(t *testing.T) {
		data, err := json.Marshal(
		map[string]interface{}{
			"comment": 1,
		},
	)

	req := httptest.NewRequest(http.MethodPut, "/comments/1/likes", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	context := commentEcho.NewContext(req, rec)
	context.Set("user_id", "somebody")
	context.SetParamNames("id")
	context.SetParamValues("1")

	assert.NotNil(t, commentRouter.likeUC)
	err = commentRouter.LikeToggle(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	body, _ := ioutil.ReadAll(rec.Body)
	log.Println("BODY", string(body))
	})
}