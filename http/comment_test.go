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
	"testing"
)

var (
	commentEcho              *echo.Echo
	commentRouter  *CommentRouter
	commentUseCase usecase.CommentUseCaseInterface
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

func TestInit(t *testing.T) {
	cont := dig.New()
	err := cont.Provide(repository.NewTestGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewUserRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(usecase.NewCommentUseCase)
	assert.Nil(t, err)

	err = cont.Invoke(func(uc usecase.CommentUseCaseInterface) {
		commentEcho = echo.New()
		mockRoot := RootRouter{commentEcho.Group("/")}
		commentRouter = NewCommentRouter(&mockRoot, uc)
		commentUseCase = uc
	})

	t.Run("Create a user jinsu to preload in list comment", func(t *testing.T) {
		user := &model.KhumuUserSimple{Username: "jinsu", Type: "active"}
		err = cont.Invoke(func(db *gorm.DB){
			dbErr := db.Create(&user).Error
			assert.Nil(t, dbErr)
			assert.Equal(t, "jinsu", user.Username)
		})
		assert.Nil(t, err)
	})

	t.Run("Create a user somebody who is not me to preload in list comment", func(t *testing.T) {

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

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		err = commentRouter.Create(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	//Unauthenticated user 혹은 자신의 username이 아닌 author username에 대한 test
}

func TestCommentRouter_List(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	context := commentEcho.NewContext(req, rec)
	context.Set("user_id", "admin")
	err := commentRouter.List(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body, _ := ioutil.ReadAll(rec.Body)
	log.Println("BODY", string(body))
}
