package repository

import (
	"github.com/khu-dev/khumu-comment/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	config.LoadTestConfig()
	database, err := gorm.Open(sqlite.Open(config.Config.DB.SQLite3.FilePath), &gorm.Config{})
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
