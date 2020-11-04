package http

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	e              *echo.Echo
	commentRouter  *CommentRouter
	commentUseCase usecase.CommentUseCaseInterface
	commentsMock   []*model.Comment
)

type CommentUseCaseMock struct{}

func (uc *CommentUseCaseMock) List(c echo.Context) []*model.Comment {
	return commentsMock
}

func (uc *CommentUseCaseMock) Get(c echo.Context) *model.Comment {
	return commentsMock[0]
}

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
		e = echo.New()
		mockRoot := RootRouter{e.Group("/")}
		commentRouter = NewCommentRouter(&mockRoot, uc)
		commentUseCase = uc
	})
	assert.Nil(t, err)
}

func TestCommentRouter_List(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	context := e.NewContext(req, rec)
	context.Set("user_id", "admin")
	err := commentRouter.List(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body, _ := ioutil.ReadAll(rec.Body)
	log.Println("BODY", string(body))
}
