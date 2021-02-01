package repository

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"testing"
)

var (
	db                    *gorm.DB
	commentRepository     CommentRepositoryInterface
	likeCommentRepository LikeCommentRepositoryInterface
)

func TestMain(m *testing.M) {
	cont := BuildContainer()
	err := cont.Invoke(func(database *gorm.DB, cr CommentRepositoryInterface, lcr LikeCommentRepositoryInterface) {
		db = database
		commentRepository = cr
		likeCommentRepository = lcr
	})
	if err != nil {
		logrus.Fatal(err)
	}

	m.Run()
}

// B는 Before each의 acronym
func B(tb testing.TB) {
	test.SetUp(db)
}

// A는 After each의 acronym
func A(tb testing.TB) {
	test.CleanUp(db)
}

// 처음 Test 환경을 만들 때에 필요한 SetUp 작업들
// 주로 IoC 컨테이너
func BuildContainer() *dig.Container {
	cont := dig.New()
	err := cont.Provide(NewTestGorm)
	if err != nil {
		logrus.Fatal(err)
	}
	err = cont.Provide(NewCommentRepositoryGorm)
	if err != nil {
		logrus.Fatal(err)
	}
	err = cont.Provide(NewLikeCommentRepositoryGorm)
	if err != nil {
		logrus.Fatal(err)
	}

	return cont
}

func TestCommentRepositoryGorm_Create(t *testing.T) {
	//parentID0 := 0
	parentID1 := 1
	t.Run("jinsu의_익명_댓글", func(t *testing.T) {
		B(t)
		defer A(t)

		comment := &model.Comment{
			Kind:      "anonymous",
			Author:    &model.KhumuUserSimple{Username: "jinsu"},
			ArticleID: 1,
			Content:   "테스트로 작성한 익명 코멘트입니다.",
			ParentID:  &parentID1,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 익명 코멘트입니다.", created.Content)
	})
	t.Run("jinsu의_기명_댓글", func(t *testing.T) {
		B(t)
		defer A(t)

		comment := &model.Comment{
			Kind:      "named",
			Author:    &model.KhumuUserSimple{Username: "jinsu"},
			ArticleID: 1,
			Content:   "테스트로 작성한 기명 코멘트입니다.",
			ParentID:  &parentID1,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "named", created.Kind)
		assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 기명 코멘트입니다.", created.Content)
	})
	t.Run("somebody의_익명_댓글", func(t *testing.T) {
		B(t)
		defer A(t)

		comment := &model.Comment{
			Kind:      "anonymous",
			Author:    &model.KhumuUserSimple{Username: "somebody"},
			ArticleID: 1,
			Content:   "테스트로 작성한 somebody의 기명 코멘트입니다.",
			ParentID:  &parentID1,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "somebody", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 somebody의 기명 코멘트입니다.", created.Content)
	})
}

func TestCommentRepositoryGorm_Update(t *testing.T) {
	t.Run("jinsu's anonymous comment", func(t *testing.T) {
		B(t)
		defer A(t)
		opt := map[string]interface{}{
			"content": "수정된 테스트로 작성된 익명 코멘트입니다.",
		}
		updated, err := commentRepository.Update(1, opt)
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
		B(t)
		defer A(t)

		// set up
		// 삭제할 코멘트와, 그 코멘트에 대한 좋아요 생성
		func() {
			commentToDelete, err = commentRepository.Create(&model.Comment{
				Author:  &model.KhumuUserSimple{Username: "jinsu"},
				Content: "A comment to be deleted.",
			})
			assert.Nil(t, err)

			likeToDelete, err = likeCommentRepository.Create(&model.LikeComment{
				CommentID: commentToDelete.ID,
				Username:  "somebody",
			})
			assert.NotNil(t, likeToDelete)

			deleted, err := commentRepository.Delete(commentToDelete.ID)
			assert.Nil(t, err)
			assert.NotNil(t, deleted)
		}()

		c, _ := commentRepository.Get(commentToDelete.ID)
		assert.Nil(t, c)
		// 좋아요의 cascade 확인
		likes := likeCommentRepository.List(&LikeCommentQueryOption{CommentID: commentToDelete.ID})
		assert.Equal(t, 0, len(likes))

	})
}

// somebody가 1번 코멘트를 좋아도록합니다.
func TestLikeCommentRepositoryGorm_Create(t *testing.T) {
	B(t)
	defer A(t)
	likeBefore := &model.LikeComment{CommentID: 1, Username: "somebody"}
	likeAfter, err := likeCommentRepository.Create(likeBefore)
	assert.Nil(t, err)
	assert.NotNil(t, likeAfter)
	assert.Equal(t, likeBefore.CommentID, likeAfter.CommentID)
	assert.Equal(t, likeBefore.Username, likeAfter.Username)
}

func TestLikeCommentRepositoryGorm_Delete(t *testing.T) {
	B(t)
	defer A(t)
	setupLike := &model.LikeComment{CommentID: 2, Username: "somebody"}
	_, err := likeCommentRepository.Create(setupLike)
	assert.Nil(t, err)

	err = likeCommentRepository.Delete(setupLike.ID)
	assert.Nil(t, err)
}

func TestLikeCommentRepositoryGorm_List(t *testing.T) {
	// 별로 List할 일 없는 듯.
}

func TestCommentRepositoryGorm_Get(t *testing.T) {
	B(t)
	defer A(t)
	comment, err := commentRepository.Get(1)

	assert.Nil(t, err)
	assert.NotNil(t, comment)
}

func TestCommentRepositoryGorm_List(t *testing.T) {
	B(t)
	defer A(t)
	comments := commentRepository.List(&CommentQueryOption{})

	assert.NotEmpty(t, comments)
	assert.NotZero(t, comments[0].Author.Username)
	assert.NotZero(t, comments[1].Author.Username)
	t.Run("List comments written by somebody", func(t *testing.T) {
		comments = commentRepository.List(&CommentQueryOption{AuthorUsername: "somebody"})
		logrus.Warn(comments)
		assert.GreaterOrEqual(t, len(comments), 1)
		assert.Equal(t, comments[0].AuthorUsername, "somebody")
		assert.Equal(t, comments[0].Author.Username, "somebody")
	})
}
