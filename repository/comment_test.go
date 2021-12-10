package repository

import (
	rcache "github.com/go-redis/cache/v8"
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/enttest"
	"github.com/khu-dev/khumu-comment/repository/cache"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ctrl             *gomock.Controller
	commentCacheRepo *cache.MockCommentCacheRepository
)

func BeforeEach(t *testing.T) {
	ctrl = gomock.NewController(t)
	commentCacheRepo = cache.NewMockCommentCacheRepository(ctrl)
	commentCacheRepo.EXPECT().FindAllParentCommentsByArticleID(gomock.Any()).Return(nil, rcache.ErrCacheMiss).AnyTimes()
	commentCacheRepo.EXPECT().SetCommentsByArticleID(gomock.Any(), gomock.Any()).AnyTimes()
}

func Test_commentRepository_Create(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		BeforeEach(t)
		db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer db.Close()
		repo := NewCommentRepository(db, commentCacheRepo, true)
		test.SetUpUsers(db)
		test.SetUpArticles(db)

		created, err := repo.Create(&data.CommentInput{
			Author:  test.UserPuppy.ID,
			Article: &test.Articles[0].ID,
			Content: "이것은 댓글",
		})
		assert.NoError(t, err)
		assert.NotNil(t, created)
		assert.Equal(t, "이것은 댓글", created.Content)
		assert.NotNil(t, created.Edges.Author)
		assert.Equal(t, test.UserPuppy.ID, created.Edges.Author.ID)
		assert.Equal(t, test.UserPuppy.Nickname, created.Edges.Author.Nickname)
	})
	//t.Run("에러 래핑", func(t *testing.T) {
	//	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//	defer db.Close()
	//	repo := NewCommentRepository(db)
	//	test.SetUpUsers(db)
	//	test.SetUpArticles(db)
	//	_, err := repo.Create(&data.CommentInput{
	//		Author:  "there is no author like this",
	//		Article: &test.Articles[0].ID,
	//		Content: "이것은 댓글",
	//	})
	//	assert.ErrorIs(t, err, errorz.ErrResourceNotFound)
	//})
}

func TestCommentRepository_FindAllParentsByAuthorID(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		BeforeEach(t)
		db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer db.Close()
		repo := NewCommentRepository(db, commentCacheRepo, true)
		test.SetUpUsers(db)
		test.SetUpArticles(db)

		created1, err := repo.Create(&data.CommentInput{
			Author:  test.UserPuppy.ID,
			Article: &test.Articles[0].ID,
			Content: "이것은 댓글",
		})
		assert.NoError(t, err)
		created2, err := repo.Create(&data.CommentInput{
			Author:  test.UserJinsu.ID,
			Article: &test.Articles[1].ID,
			Content: "이것은 댓글",
		})
		assert.NoError(t, err)
		result, err := repo.FindAllParentCommentsByAuthorID(test.UserPuppy.ID)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, created1.ID, result[0].ID)

		result, err = repo.FindAllParentCommentsByArticleID(test.Articles[1].ID)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, created2.ID, result[0].ID)
		assert.Equal(t, test.Articles[1].ID, result[0].ID)
		//assert.Equal(t, test.UserPuppy.ID, created.Edges.Author.ID)
		//assert.Equal(t, test.UserPuppy.Nickname, created.Edges.Author.Nickname)
	})
}

func TestCommentRepository_Get(t *testing.T) {
	t.Run("ErrResourceNotFound 에러 랩핑", func(t *testing.T) {
		BeforeEach(t)
		db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer db.Close()
		repo := NewCommentRepository(db, commentCacheRepo, true)
		test.SetUpUsers(db)
		test.SetUpArticles(db)
		// error type warpping 테스트
		_, err := repo.Get(10000)
		assert.True(t, ent.IsNotFound(err))
	})
}
