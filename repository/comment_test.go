package repository

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"testing"
)

var (
	db              *gorm.DB
	commentRepository      CommentRepositoryInterface
	likeCommentRepository LikeCommentRepositoryInterface
	authorIDForList string = "Arsenal"
	commentID       int
)

// 그냥 Init하는데, 이때에도 테스트 적용
func TestInit(t *testing.T) {
	// container build
	cont := dig.New()
	err := cont.Provide(NewTestGorm)
	assert.Nil(t, err)

	err = cont.Provide(NewCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(NewLikeCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Invoke(func(database *gorm.DB){
		db = database
	})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = cont.Invoke(func(cr CommentRepositoryInterface, lcr LikeCommentRepositoryInterface){
		commentRepository = cr
		likeCommentRepository = lcr
	})

	t.Run("Create sample users to preload in list comment", func(t *testing.T) {
		for _, user := range test.UsersData{
			username := user.Username
			t.Log("Create a user named ", username)
			err = cont.Invoke(func(db *gorm.DB){
				dbErr := db.Create(&user).Error
				assert.Nil(t, dbErr)
				assert.Equal(t, username, user.Username)
			})
			assert.Nil(t, err)
		}
	})

	assert.Nil(t, err)
}

func TestCommentRepositoryGorm_Create(t *testing.T) {
	t.Run("Anonymous comment", func(t *testing.T){
		parentID := uint(1)
		comment := &model.Comment{
			Kind:           "anonymous",
			Author: &model.KhumuUserSimple{Username: "jinsu"},
			ArticleID:      1,
			Content:        "테스트로 작성한 익명 코멘트입니다.",
			ParentID:       &parentID,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 익명 코멘트입니다.", created.Content)
	})
	t.Run("Named comment", func(t *testing.T){
		parentID := uint(1)
		comment := &model.Comment{
			Kind:           "named",
			Author: &model.KhumuUserSimple{Username: "jinsu"},
			ArticleID:      1,
			Content:        "테스트로 작성한 기명 코멘트입니다.",
			ParentID:       &parentID,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "named", created.Kind)
		assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 기명 코멘트입니다.", created.Content)

	})
}

func TestCommentRepositoryGorm_Get(t *testing.T) {
	comment := commentRepository.Get(commentID)
	assert.NotNil(t, comment)
}

func TestCommentRepositoryGorm_List(t *testing.T) {
	comments := commentRepository.List(&CommentQueryOption{})
	assert.NotEmpty(t, comments)
	assert.Equal(t, "jinsu", comments[0].Author.Username)
	assert.Equal(t, "jinsu", comments[1].Author.Username)
	commentID = int(comments[0].ID)
	comments = commentRepository.List(&CommentQueryOption{AuthorID: authorIDForList})
	assert.Len(t, comments, 0)
}

func TestLikeCommentRepositoryGorm_Create(t *testing.T) {
	likeBefore := &model.LikeComment{CommentID: 1, Username: "jinsu"}
	likeAfter, err := likeCommentRepository.Create(likeBefore)
	assert.Nil(t, err)
	assert.NotNil(t, likeAfter)
	assert.Equal(t, likeBefore.CommentID, likeAfter.CommentID)
	assert.Equal(t, likeBefore.Username, likeAfter.Username)
}