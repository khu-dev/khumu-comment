package usecase

import (
	rcache "github.com/go-redis/cache/v8"
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/enttest"
	"github.com/khu-dev/khumu-comment/infra/khumu"
	"github.com/khu-dev/khumu-comment/infra/message"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/repository/cache"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var (
	db                   *ent.Client
	mockMessagePublisher *message.MockMessagePublisher
	mockKhumuApiAdapter  *khumu.MockKhumuAPIAdapter
	mockCommentCacheRepo *cache.MockCommentCacheRepository
	mockLikeCacheRepo    *cache.MockLikeCommentCacheRepository
	mockCommentRepo      *repository.MockCommentRepository
	mockLikeRepo         *repository.MockLikeCommentRepository

	commentRepo        repository.CommentRepository
	likeRepo           repository.LikeCommentRepository
	commentUseCase     *CommentUseCase
	likeCommentUseCase *LikeCommentUseCase
	ctrl               *gomock.Controller
)

// B는 Before each의 acronym
func BeforeEach(tb testing.TB) {
	ctrl = gomock.NewController(tb)
	db = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//db = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&_fk=1")

	mockMessagePublisher = message.NewMockMessagePublisher(ctrl)
	mockKhumuApiAdapter = khumu.NewMockKhumuAPIAdapter(ctrl)
	mockCommentCacheRepo = cache.NewMockCommentCacheRepository(ctrl)
	mockLikeCacheRepo = cache.NewMockLikeCommentCacheRepository(ctrl)
	commentRepo = repository.NewCommentRepository(db, mockCommentCacheRepo, true)
	likeRepo = repository.NewLikeCommentRepository(db, mockLikeCacheRepo, true)

	commentUseCase = &CommentUseCase{
		Repo:            commentRepo,
		entClient:       db,
		snsClient:       mockMessagePublisher,
		khumuAPIAdapter: mockKhumuApiAdapter,
		likeRepo:        likeRepo,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repo:        likeRepo,
		CommentRepo: commentRepo,
	}

	mockMessagePublisher.EXPECT().Publish(gomock.Any()).DoAndReturn(
		func(message interface{}) error {
			tb.Log("그냥 테스트라서 푸시 알림 패스")
			return nil
		}).AnyTimes()

	mockKhumuApiAdapter.EXPECT().IsAuthor(gomock.Any(), gomock.Any()).DoAndReturn(func(articleID int, authorID string) <-chan bool {
		ch := make(chan bool, 1)
		ch <- false
		return ch
	}).AnyTimes()

	mockCommentCacheRepo.EXPECT().FindAllParentCommentsByArticleID(gomock.Any()).Return(nil, rcache.ErrCacheMiss).AnyTimes()
	mockCommentCacheRepo.EXPECT().SetCommentsByArticleID(gomock.Any(), gomock.Any()).Return().AnyTimes()

	mockLikeCacheRepo.EXPECT().FindAllByCommentID(gomock.Any()).Return([]*ent.LikeComment{}, nil).AnyTimes()
	mockLikeCacheRepo.EXPECT().SetLikesByCommentID(gomock.Any(), gomock.Any()).Return().AnyTimes()

	test.SetUpUsers(db)
	test.SetUpArticles(db)
	test.SetUpStudyArticles(db)
	test.SetUpComments(db)
}

func AfterEach(tb testing.TB) {
	if err := db.Close(); err != nil {
		tb.Error(err)
	}
	ctrl.Finish()
	//time.Sleep(time.Second)
}

func BeforeUnitTest(t *testing.T) {
	ctrl = gomock.NewController(t)

	mockMessagePublisher = message.NewMockMessagePublisher(ctrl)
	mockKhumuApiAdapter = khumu.NewMockKhumuAPIAdapter(ctrl)
	mockCommentRepo = repository.NewMockCommentRepository(ctrl)
	mockLikeRepo = repository.NewMockLikeCommentRepository(ctrl)

	commentUseCase = &CommentUseCase{
		Repo:            mockCommentRepo,
		entClient:       nil,
		snsClient:       mockMessagePublisher,
		khumuAPIAdapter: mockKhumuApiAdapter,
		likeRepo:        mockLikeRepo,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repo:        mockLikeRepo,
		CommentRepo: mockCommentRepo,
	}

	mockMessagePublisher.EXPECT().Publish(gomock.Any()).DoAndReturn(
		func(message interface{}) error {
			t.Log("그냥 테스트라서 푸시 알림 패스")
			return nil
		}).AnyTimes()

	mockKhumuApiAdapter.EXPECT().IsAuthor(gomock.Any(), gomock.Any()).DoAndReturn(func(articleID int, authorID string) <-chan bool {
		ch := make(chan bool, 1)
		ch <- false
		return ch
	}).AnyTimes()
}

func AfterUnitTest(tb testing.TB) {
	ctrl.Finish()
}
