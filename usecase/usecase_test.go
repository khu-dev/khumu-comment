package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/enttest"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/external/khumu"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var (
	repo                *ent.Client
	mockSnsClient       *external.MockSnsClient
	mockRedisAdapter    *external.MockRedisAdapter
	mockKhumuApiAdapter *khumu.MockKhumuAPIAdapter
	commentUseCase      *CommentUseCase
	likeCommentUseCase  *LikeCommentUseCase
	ctrl                *gomock.Controller
)

// B는 Before each의 acronym
func BeforeCommentUseCaseTest(tb testing.TB) {
	ctrl = gomock.NewController(tb)
	repo = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	mockSnsClient = external.NewMockSnsClient(ctrl)
	mockRedisAdapter = external.NewMockRedisAdapter(ctrl)
	mockKhumuApiAdapter = khumu.NewMockKhumuAPIAdapter(ctrl)
	commentUseCase = &CommentUseCase{
		Repo:            repository.NewCommentRepository(repo),
		entclient:       repo,
		SnsClient:       mockSnsClient,
		khumuAPIAdapter: mockKhumuApiAdapter,
		redisAdapter:    mockRedisAdapter,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repo:        repository.NewLikeCommentRepository(repo),
		CommentRepo: repository.NewCommentRepository(repo),
	}

	mockSnsClient.EXPECT().PublishMessage(gomock.Any()).DoAndReturn(
		func(message interface{}) error {
			tb.Log("그냥 테스트라서 푸시 알림 패스")
			return nil
		}).AnyTimes()
	mockRedisAdapter.EXPECT().Refresh(gomock.AssignableToTypeOf(1)).AnyTimes()
	mockRedisAdapter.EXPECT().GetAllByArticle(gomock.AssignableToTypeOf(1)).AnyTimes()
	mockKhumuApiAdapter.EXPECT().IsAuthor(gomock.Any(), gomock.Any()).DoAndReturn(func(articleID int, authorID string) <-chan bool {
		ch := make(chan bool, 1)
		ch <- false
		return ch
	}).AnyTimes()
	test.SetUpUsers(repo)
	test.SetUpArticles(repo)
	test.SetUpStudyArticles(repo)
	test.SetUpComments(repo)
}

// A는 After each의 acronym
func A(tb testing.TB) {
	repo.Close()
}
