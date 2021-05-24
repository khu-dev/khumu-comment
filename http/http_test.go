package http

import (
    "github.com/golang/mock/gomock"
    "github.com/khu-dev/khumu-comment/ent"
    "github.com/khu-dev/khumu-comment/ent/enttest"
    "github.com/khu-dev/khumu-comment/external"
    "github.com/khu-dev/khumu-comment/test"
    "github.com/khu-dev/khumu-comment/usecase"
    "github.com/labstack/echo/v4"
    "testing"
)

var (
	repo               *ent.Client
	mockSnsClient      *external.MockSnsClient
	commentUseCase     *usecase.CommentUseCase
	likeCommentUseCase *usecase.LikeCommentUseCase
    commentEcho   *echo.Echo
	commentRouter *CommentRouter
)

// B는 Before each의 acronym
func BeforeCommentUseCaseTest(tb testing.TB) {
	ctrl = gomock.NewController(tb)
	repo = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	mockSnsClient = external.NewMockSnsClient(ctrl)
	commentUseCase = &CommentUseCase{
		Repo:      repo,
		SnsClient: mockSnsClient,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repo: repo,
	}

	mockSnsClient.EXPECT().PublishMessage(gomock.Any()).DoAndReturn(
		func(message interface{}) {
			tb.Log("그냥 테스트라서 푸시 알림 패스")
		}).AnyTimes()

	test.SetUpUsers(repo)
	test.SetUpArticles(repo)
	test.SetUpComments(repo)
}

// A는 After each의 acronym
func A(tb testing.TB) {
	repo.Close()
}
