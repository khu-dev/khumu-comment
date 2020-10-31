package repository

import (
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"testing"
)

var (
	db              *gorm.DB
	repository      *CommentRepositoryGorm
	authorIDForList string = "Arsenal"
	commentID       int
)

// 그냥 Init하는데, 이때에도 테스트 적용
func TestInit(t *testing.T) {
	config.Load()
	originalDataPath := config.Config.DB.SQLite3.FilePath
	config.LoadTestConfig()
	testDataPath := config.Config.DB.SQLite3.FilePath
	copyBackupData(originalDataPath, testDataPath, t)
	database, err := gorm.Open(sqlite.Open(testDataPath), &gorm.Config{})
	db = database
	assert.Nil(t, err)
	assert.NotNil(t, db)
	repository = &CommentRepositoryGorm{DB: db}

}

func TestList(t *testing.T) {
	comments := repository.List(&CommentQueryOption{})
	assert.NotEmpty(t, comments)
	commentID = int(comments[0].ID)
	comments = repository.List(&CommentQueryOption{AuthorID: authorIDForList})
	assert.Len(t, comments, 0)
}

func TestGet(t *testing.T) {
	comment := repository.Get(commentID)
	assert.NotNil(t, comment)
}

func TestCreateAnonymousComment(t *testing.T){
	parentID := uint(3)
	comment := &model.Comment{
		Kind: "anonymous",
		AuthorUsername: "jinsu",
		ArticleID: 1,
		Content: "테스트로 작성한 코멘트입니다.",
		ParentID: &parentID,
	}
	err := repository.Create(comment)
	assert.Nil(t, err)
}

func copyBackupData(src, dest string, t *testing.T){
	input, err := ioutil.ReadFile(src)
	assert.Nil(t, err)
	err = ioutil.WriteFile(dest, input, 0644)
	assert.Nil(t, err)
}