package repository

import (
	"context"
	"database/sql"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"testing"
)

var (
	gormCommentRepository     *CommentRepositoryGorm
	gormLikeCommentRepository *LikeCommentRepositoryGorm
)

func BeforeGormCommentRepository(tb testing.TB) {
	db := NewTestGorm()
	migrateAll(tb, db)
	gormCommentRepository = NewCommentRepositoryGorm(db).(*CommentRepositoryGorm)
	gormLikeCommentRepository = NewLikeCommentRepositoryGorm(db).(*LikeCommentRepositoryGorm)
	test.SetUp()
}

func migrateAll(tb testing.TB, db *gorm.DB) {
	err := db.AutoMigrate(&model.Board{}, &model.KhumuUser{}, &model.Article{}, &model.Comment{}, &model.LikeComment{})
	if err != nil {
		tb.Fatal(err)
	} else {
		tb.Log("Migrate all!")
	}
}

func AfterGormCommentRepository(tb testing.TB) {
	err := gormCommentRepository.DB.Migrator().DropTable(&model.LikeComment{}, &model.Comment{}, &model.Article{}, &model.KhumuUserSimple{}, &model.Board{})
	if err != nil {
		logrus.Fatal(err)
	}
}

func TestCommentRepositoryGorm_Create(t *testing.T) {
	//parentID0 := 0
	var parentID1 int64 = 1
	t.Run("jinsu의_익명_댓글", func(t *testing.T) {
		BeforeGormCommentRepository(t)
		defer AfterGormCommentRepository(t)

		comment := test.Comment1JinsuAnnonymous
		comment.State = "exists"
		//		comment := &model.Comment{
		//	Kind:      "anonymous",
		//	AuthorUsername: "jinsu",
		//	ArticleID: 1,
		//	Content:   "테스트로 작성한 익명 코멘트입니다.",
		//	ParentID:  &parentID1,
		//}
		created, err := gormCommentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "jinsu", created.AuthorUsername)
		// test상 foreinkey 제약도 안 걸고, User도 생성 안했으면 nil이라 생략
		//assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, comment.Content, created.Content)
	})
	t.Run("jinsu의_기명_댓글", func(t *testing.T) {
		BeforeGormCommentRepository(t)
		defer AfterGormCommentRepository(t)

		comment := &model.Comment{
			Kind:           "named",
			AuthorUsername: "jinsu",
			ArticleID:      null.Int{sql.NullInt64{1, true}},
			Content:        "테스트로 작성한 기명 코멘트입니다.",
			ParentID:       null.Int{sql.NullInt64{parentID1, true}},
		}
		created, err := gormCommentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "named", created.Kind)
		assert.Equal(t, "jinsu", created.AuthorUsername)
		//assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 기명 코멘트입니다.", created.Content)
	})
	t.Run("somebody의_익명_댓글", func(t *testing.T) {
		BeforeGormCommentRepository(t)
		defer AfterGormCommentRepository(t)

		comment := &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "somebody",
			ArticleID:      null.Int{sql.NullInt64{1, true}},
			Content:        "테스트로 작성한 somebody의 기명 코멘트입니다.",
			ParentID:       null.Int{sql.NullInt64{parentID1, true}},
		}
		created, err := gormCommentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "somebody", created.AuthorUsername)
		//assert.Equal(t, "somebody", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 somebody의 기명 코멘트입니다.", created.Content)
	})
	t.Run("jinsu의_익명_스터디_댓글", func(t *testing.T) {
		BeforeGormCommentRepository(t)
		defer AfterGormCommentRepository(t)

		comment := test.Comment1JinsuAnnonymous
		comment.State = "exists"
		comment.ArticleID = null.Int{sql.NullInt64{}}
		comment.StudyArticleID = null.Int{sql.NullInt64{1, true}}
		//		comment := &model.Comment{
		//	Kind:      "anonymous",
		//	AuthorUsername: "jinsu",
		//	ArticleID: 1,
		//	Content:   "테스트로 작성한 익명 코멘트입니다.",
		//	ParentID:  &parentID1,
		//}
		created, err := gormCommentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "jinsu", created.AuthorUsername)
		// test상 foreinkey 제약도 안 걸고, User도 생성 안했으면 nil이라 생략
		//assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, comment.Content, created.Content)
	})
}

func TestCommentRepositoryGorm_Update(t *testing.T) {
	t.Run("jinsu's anonymous comment", func(t *testing.T) {
		BeforeGormCommentRepository(t)
		defer AfterGormCommentRepository(t)
		opt := map[string]interface{}{
			"content": "수정된 테스트로 작성된 익명 코멘트입니다.",
		}
		gormCommentRepository.Create(test.Comment1JinsuAnnonymous)
		updated, err := gormCommentRepository.Update(test.Comment1JinsuAnnonymous.ID, opt)
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, opt["content"].(string), updated.Content)
	})
}

func TestCommentRepositoryGorm_Delete(t *testing.T) {
	var err error
	var commentToDelete *model.Comment
	var likeToDelete *model.LikeComment

	t.Run("Delete a comment", func(t *testing.T) {
		BeforeGormCommentRepository(t)
		defer AfterGormCommentRepository(t)

		// set up
		// 삭제할 코멘트와, 그 코멘트에 대한 좋아요 생성
		func() {
			// Create에 대한 에러체크는 위에서 했다고 믿음.
			commentToDelete, _ = gormCommentRepository.Create(test.Comment1JinsuAnnonymous)
			likeToDelete, err = gormLikeCommentRepository.Create(&model.LikeComment{
				CommentID: commentToDelete.ID,
				Username:  "somebody",
			})
			assert.NotNil(t, likeToDelete)
			assert.Nil(t, err)

			deleted, err := gormCommentRepository.Delete(commentToDelete.ID)
			assert.Nil(t, err)
			assert.NotNil(t, deleted)
		}()

		c, _ := gormCommentRepository.Get(commentToDelete.ID)
		assert.Nil(t, c)
		// 좋아요의 cascade 확인
		// 테스트 시에는 ForeignKey constraint를 사용하지 않아서 불가능..
		//likes := gormLikeCommentRepository.List(&LikeCommentQueryOption{CommentID: commentToDelete.ID})
		//assert.Equal(t, 0, len(likes))
	})
}

// somebody가 1번 코멘트를 좋아도록합니다.
func TestLikeCommentRepositoryGorm_Create(t *testing.T) {
	BeforeGormCommentRepository(t)
	defer AfterGormCommentRepository(t)
	likeBefore := &model.LikeComment{CommentID: 1, Username: "somebody"}
	likeAfter, err := gormLikeCommentRepository.Create(likeBefore)
	assert.Nil(t, err)
	assert.NotNil(t, likeAfter)
	assert.Equal(t, likeBefore.CommentID, likeAfter.CommentID)
	assert.Equal(t, likeBefore.Username, likeAfter.Username)
}

func TestLikeCommentRepositoryGorm_Delete(t *testing.T) {
	BeforeGormCommentRepository(t)
	defer AfterGormCommentRepository(t)
	setupLike := &model.LikeComment{CommentID: 2, Username: "somebody"}
	_, err := gormLikeCommentRepository.Create(setupLike)
	assert.Nil(t, err)

	err = gormLikeCommentRepository.Delete(setupLike.ID)
	assert.Nil(t, err)
}

func TestLikeCommentRepositoryGorm_List(t *testing.T) {
	// 별로 List할 일 없는 듯.
}

func TestCommentRepositoryGorm_Get(t *testing.T) {
	BeforeGormCommentRepository(t)
	defer AfterGormCommentRepository(t)
	_, _ = gormCommentRepository.Create(test.Comment1JinsuAnnonymous)
	comment, err := gormCommentRepository.Get(1)

	assert.Nil(t, err)
	assert.NotNil(t, comment)
}

func TestCommentRepositoryGorm_List(t *testing.T) {
	BeforeGormCommentRepository(t)
	defer AfterGormCommentRepository(t)
	for _, c := range test.Comments {
		gormCommentRepository.Create(c)
	}

	comments := gormCommentRepository.List(&CommentQueryOption{})
	assert.NotEmpty(t, comments)
	assert.NotZero(t, comments[0].AuthorUsername)
	assert.NotZero(t, comments[1].AuthorUsername)
	t.Run("List comments written by somebody", func(t *testing.T) {
		comments = gormCommentRepository.List(&CommentQueryOption{AuthorUsername: "somebody"})
		assert.GreaterOrEqual(t, len(comments), 1)
		assert.Equal(t, comments[0].AuthorUsername, "somebody")
		//assert.Equal(t, comments[0].Author.Username, "somebody")
	})
}

func TestEntgo(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		logrus.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}

	user, err := client.User.Create().
		SetUsername("jinsu").
		SetPassword("123123").
		SetStudentNumber("2016101168").
		Save(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	logrus.Warn(user)

	user2, err := client.User.Create().
		SetUsername("jinsu2").
		SetPassword("123123").
		SetStudentNumber("2016101168").
		Save(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	logrus.Warn(user2)

	comment, err := client.Comment.Create().
		SetAuthor(user).
		SetContent("hello, world").
		SetState("exists").
		Save(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	logrus.Warn(comment)
	author, err := comment.QueryAuthor().All(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	logrus.Warn("Author:", author)
}
