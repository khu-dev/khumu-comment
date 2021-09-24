package usecase

import (
	rcache "github.com/go-redis/cache/v8"
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/enttest"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/external/khumu"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/repository/cache"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var (
	db                   *ent.Client
	mockSnsClient        *external.MockSnsClient
	mockKhumuApiAdapter  *khumu.MockKhumuAPIAdapter
	mockCommentCacheRepo *cache.MockCommentCacheRepository
	mockLikeCacheRepo    *cache.MockLikeCommentCacheRepository
	commentRepo          repository.CommentRepository
	likeRepo             repository.LikeCommentRepository
	commentUseCase       *CommentUseCase
	likeCommentUseCase   *LikeCommentUseCase
	ctrl                 *gomock.Controller
)

// B는 Before each의 acronym
func BeforeCommentUseCaseTest(tb testing.TB) {
	ctrl = gomock.NewController(tb)
	db = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//db = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&_fk=1")

	mockSnsClient = external.NewMockSnsClient(ctrl)
	mockKhumuApiAdapter = khumu.NewMockKhumuAPIAdapter(ctrl)
	mockCommentCacheRepo = cache.NewMockCommentCacheRepository(ctrl)
	mockLikeCacheRepo = cache.NewMockLikeCommentCacheRepository(ctrl)
	commentRepo = repository.NewCommentRepository(db, mockCommentCacheRepo, true)
	likeRepo = repository.NewLikeCommentRepository(db, mockLikeCacheRepo)

	commentUseCase = &CommentUseCase{
		Repo:            commentRepo,
		entclient:       db,
		SnsClient:       mockSnsClient,
		khumuAPIAdapter: mockKhumuApiAdapter,
		likeRepo:        likeRepo,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repo:        likeRepo,
		CommentRepo: commentRepo,
	}

	mockSnsClient.EXPECT().PublishMessage(gomock.Any()).DoAndReturn(
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

// A는 After each의 acronym
func A(tb testing.TB) {
	if err := db.Close(); err != nil {
		tb.Error(err)
	}
	ctrl.Finish()
	//time.Sleep(time.Second)
}
