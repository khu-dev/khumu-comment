package http

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	e              *echo.Echo
	commentRouter  *CommentRouter
	commentUseCase *CommentUseCaseMock
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
	e = echo.New()
	commentUseCase = &CommentUseCaseMock{}
	assert.NotNil(t, commentUseCase)
	commentsMock = append(commentsMock, &model.Comment{ID: 1})
	commentsMock = append(commentsMock, &model.Comment{ID: 2})
}

func TestNewCommentRouter(t *testing.T) {
	commentRouter = NewCommentRouter(e.Group("/"), &CommentUseCaseMock{})
	assert.NotNil(t, commentRouter)
}

func TestCommentUseCaseList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	context := e.NewContext(req, rec)
	err := commentRouter.ListComment(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body, _ := ioutil.ReadAll(rec.Body)
	log.Println("BODY", string(body))
}
