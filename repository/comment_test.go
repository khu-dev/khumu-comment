package repository

import (
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"io/ioutil"
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
	// 원본 test db를 카피해서 작업
	config.Load()
	originalDataPath := config.Config.DB.SQLite3.FilePath
	config.LoadTestConfig()
	testDataPath := config.Config.DB.SQLite3.FilePath
	copyBackupData(originalDataPath, testDataPath, t)

	// container build
	cont := dig.New()
	err := cont.Provide(NewGorm)
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

	assert.Nil(t, err)
}

func TestCommentRepositoryGorm_List(t *testing.T) {
	comments := commentRepository.List(&CommentQueryOption{})
	assert.NotEmpty(t, comments)
	commentID = int(comments[0].ID)
	comments = commentRepository.List(&CommentQueryOption{AuthorID: authorIDForList})
	assert.Len(t, comments, 0)
}

func TestCommentRepositoryGorm_Get(t *testing.T) {
	comment := commentRepository.Get(commentID)
	assert.NotNil(t, comment)
}

func TestCommentRepositoryGorm_Create(t *testing.T) {
	t.Run("Anonymous comment", func(t *testing.T){
		parentID := uint(1)
		comment := &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "jinsu",
			ArticleID:      1,
			Content:        "테스트로 작성한 익명 코멘트입니다.",
			ParentID:       &parentID,
		}
		err := commentRepository.Create(comment)
		assert.Nil(t, err)
	})
	t.Run("Named comment", func(t *testing.T){
		parentID := uint(1)
		comment := &model.Comment{
			Kind:           "named",
			AuthorUsername: "jinsu",
			ArticleID:      1,
			Content:        "테스트로 작성한 기명 코멘트입니다.",
			ParentID:       &parentID,
		}
		err := commentRepository.Create(comment)
		assert.Nil(t, err)
	})
}

func TestLikeCommentRepositoryGorm_Create(t *testing.T) {
	likeBefore := &model.LikeComment{CommentID: 1, Username: "jinsu"}
	likeAfter, err := likeCommentRepository.Create(likeBefore)
	assert.Nil(t, err)
	assert.NotNil(t, likeAfter)
	assert.Equal(t, likeBefore.CommentID, likeAfter.CommentID)
	assert.Equal(t, likeBefore.Username, likeAfter.Username)
}


func copyBackupData(src, dest string, t *testing.T) {
	input, err := ioutil.ReadFile(src)
	assert.Nil(t, err)
	err = ioutil.WriteFile(dest, input, 0644)
	assert.Nil(t, err)
}
