package http

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var (
	ctrl                   *gomock.Controller
	mockCommentUseCase     *usecase.MockCommentUseCaseInterface
	mockLikeCommentUseCase *usecase.MockLikeCommentUseCaseInterface

	commentEcho   *echo.Echo
	commentRouter *CommentRouter
)

func TestMain(m *testing.M) {
	m.Run()
}

// B는 Before each의 acronym
func BeforeCommentRouterTest(t testing.TB) {
	test.SetUp()
	ctrl = gomock.NewController(t)
	mockCommentUseCase = usecase.NewMockCommentUseCaseInterface(ctrl)
	mockLikeCommentUseCase = usecase.NewMockLikeCommentUseCaseInterface(ctrl)
	mockCommentUseCase.EXPECT().Create(gomock.Any()).DoAndReturn(func(c *_model.Comment) (*_model.Comment, error) {
		c = test.Comment1JinsuAnnonymous
		return c, nil
	}).AnyTimes()
	mockCommentUseCase.EXPECT().List(gomock.Any(), gomock.Any()).DoAndReturn(func(username string, opt *_repository.CommentQueryOption) ([]*_model.Comment, error) {
		return test.Comments, nil
	}).AnyTimes()
	mockCommentUseCase.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(username string, id int) (*_model.Comment, error) {
		return test.Comment1JinsuAnnonymous, nil
	}).AnyTimes()
	mockCommentUseCase.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(username string, id int, opt map[string]interface{}) (*_model.Comment, error) {
		return test.Comment1JinsuAnnonymous, nil
	}).AnyTimes()
	mockCommentUseCase.EXPECT().Delete(gomock.Any()).DoAndReturn(func(id int) (*_model.Comment, error) {
		return test.Comment1JinsuAnnonymous, nil
	}).AnyTimes()

	commentEcho = echo.New()
	fakeRoot := RootRouter{commentEcho.Group("/comments")}
	commentRouter = NewCommentRouter(&fakeRoot, mockCommentUseCase, mockLikeCommentUseCase)
}

// A는 After each의 acronym
func AfterCommentRouterTest(tb testing.TB) {
}

func TestCommentRouter_Create(t *testing.T) {
	BeforeCommentRouterTest(t)
	defer AfterCommentRouterTest(t)
	defaultComment := map[string]interface{}{
		"kind": "anonymous",
		"author": map[string]interface{}{
			"username": "jinsu",
		},
		"article": 1,
		"content": "jinsu의 익명 댓글을 테스트하고 있습니다.\nhello, world",
		"parent":  nil,
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
	BeforeCommentRouterTest(t)
	defer AfterCommentRouterTest(t)
	t.Run("Admin user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "admin")
		err := commentRouter.List(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Jinsu user with required arguments", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments?article=1", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		err := commentRouter.List(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Jinsu user without required arguments", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		err := commentRouter.List(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

func TestCommentRouter_Get(t *testing.T) {
	BeforeCommentRouterTest(t)
	defer AfterCommentRouterTest(t)
	t.Run("Get a comment", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.Set("user_id", "jinsu")
		context.SetParamNames("id")
		context.SetParamValues("1")
		//con
		err := commentRouter.Get(context)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCommentRouter_Delete(t *testing.T) {
	t.Run("부모 댓글 삭제", func(t *testing.T) {
		BeforeCommentRouterTest(t)
		defer AfterCommentRouterTest(t)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.SetParamNames("id")
		context.SetParamValues(strconv.Itoa(1))

		err := commentRouter.Delete(context)
		assert.NoError(t, err)
		assert.Equal(t, 204, rec.Code)
		//assert.Nil(t, err)
		//assert.Equal(t, http.StatusNoContent, rec.Code)
		//
		//// 테스트 시나리오 상 ID:1 댓글은 부모 댓글이다.
		//deletedStateParentComment, err := mockCommentUseCase.Get("jinsu", 1)
		//assert.NoError(t, err)
		//assert.Equal(t, "deleted", deletedStateParentComment.State)
		//assert.Equal(t, usecase.DeletedCommentContent, deletedStateParentComment.Content)
	})

	t.Run("자식 댓글 삭제", func(t *testing.T) {
		BeforeCommentRouterTest(t)
		defer AfterCommentRouterTest(t)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		context := commentEcho.NewContext(req, rec)
		context.SetParamNames("id")
		context.SetParamValues(strconv.Itoa(test.Comment6JinsuNamedFromComment1.ID))
		//con
		err := commentRouter.Delete(context)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)

		//deletedStateChildComment, err := mockCommentUseCase.Get("jinsu", test.Comment6JinsuNamedFromComment1.ID)
		//assert.NoError(t, err)
		//assert.Equal(t, "deleted", deletedStateChildComment.State)
	})
}

func TestLikeCommentRouter_Toggle(t *testing.T) {
	t.Run("Somebody likes and dislikes jinsu's comment 1.", func(t *testing.T) {
		BeforeCommentRouterTest(t)
		defer AfterCommentRouterTest(t)

		// like
		func() {
			mockLikeCommentUseCase.EXPECT().Toggle(gomock.Any()).Return(true, nil)
			req := httptest.NewRequest(http.MethodPatch, "/", nil)
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
		}()

		requestBody, err := json.Marshal(
			map[string]interface{}{
				"comment": 1,
			})

		// dislike by liking again
		func() {
			mockLikeCommentUseCase.EXPECT().Toggle(gomock.Any()).Return(false, nil)
			req := httptest.NewRequest(http.MethodPut, "/comments/1/likes", bytes.NewReader(requestBody))
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
		}()
	})
}
