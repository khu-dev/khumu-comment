package repository

import (
	"fmt"
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
)

// 그냥 Init하는데, 이때에도 테스트 적용
func TestSetUp(t *testing.T) {
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
	t.Run("jinsu's anonymous comment", func(t *testing.T){
		comment := &model.Comment{
			Kind:           "anonymous",
			Author: &model.KhumuUserSimple{Username: "jinsu"},
			ArticleID:      1,
			Content:        "테스트로 작성한 익명 코멘트입니다.",
			ParentID:       1,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 익명 코멘트입니다.", created.Content)
	})
	t.Run("jinsu's named comment", func(t *testing.T){
		comment := &model.Comment{
			Kind:           "named",
			Author: &model.KhumuUserSimple{Username: "jinsu"},
			ArticleID:      1,
			Content:        "테스트로 작성한 기명 코멘트입니다.",
			ParentID:       1,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "named", created.Kind)
		assert.Equal(t, "jinsu", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 기명 코멘트입니다.", created.Content)
	})
	t.Run("somebody's anonymous comment", func(t *testing.T){
		comment := &model.Comment{
			Kind:           "anonymous",
			Author: &model.KhumuUserSimple{Username: "somebody"},
			ArticleID:      1,
			Content:        "테스트로 작성한 somebody의 기명 코멘트입니다.",
			ParentID:       0,
		}
		created, err := commentRepository.Create(comment)
		assert.Nil(t, err)
		assert.Equal(t, "anonymous", created.Kind)
		assert.Equal(t, "somebody", created.Author.Username)
		assert.Equal(t, "테스트로 작성한 somebody의 기명 코멘트입니다.", created.Content)
	})
}

func TestCommentRepositoryGorm_Update(t *testing.T) {
	t.Run("jinsu's anonymous comment", func(t *testing.T){
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

	// 삭제할 코멘트와, 그 코멘트에 대한 좋아요 생성
	t.Run("Setup", func(t *testing.T) {

		commentToDelete, err = commentRepository.Create(&model.Comment{
			Author: &model.KhumuUserSimple{Username: "jinsu"},
			Content: "A comment to be deleted.",
		})
		assert.Nil(t, err)

		likeToDelete, err = likeCommentRepository.Create(&model.LikeComment{
			CommentID: commentToDelete.ID,
			Username: "somebody",
		})
		assert.NotNil(t, likeToDelete)
	})

	t.Run("Delete a comment", func(t *testing.T) {
		deleted, err := commentRepository.Delete(commentToDelete.ID)
		assert.Nil(t, err)
		assert.NotNil(t, deleted)

		c, _ := commentRepository.Get(commentToDelete.ID)
		assert.Nil(t, c)
		// 좋아요의 cascade 확인
		likes := likeCommentRepository.List(&LikeCommentQueryOption{CommentID: commentToDelete.ID})
		assert.Equal(t, 0, len(likes))

	})
}

// somebody가 1번 코멘트를 좋아도록합니다.
func TestLikeCommentRepositoryGorm_Create(t *testing.T) {
	likeBefore := &model.LikeComment{CommentID: 1, Username: "somebody"}
	likeAfter, err := likeCommentRepository.Create(likeBefore)
	assert.Nil(t, err)
	assert.NotNil(t, likeAfter)
	assert.Equal(t, likeBefore.CommentID, likeAfter.CommentID)
	assert.Equal(t, likeBefore.Username, likeAfter.Username)
}

func TestLikeCommentRepositoryGorm_Delete(t *testing.T) {
	setupLike := &model.LikeComment{CommentID: 2, Username: "somebody"}
	_, err := likeCommentRepository.Create(setupLike)
	assert.Nil(t, err)

	err = likeCommentRepository.Delete(setupLike.ID)
	assert.Nil(t, err)
}

func TestLikeCommentRepositoryGorm_List(t *testing.T) {
	// 위의 테스트에서 1번 코멘트에 대한 like comment를 생성했음.
	t.Run("Somebody likes comment 1.", func(t *testing.T) {
		likes := likeCommentRepository.List(&LikeCommentQueryOption{})
		assert.GreaterOrEqual(t, 1, len(likes))

		likes = likeCommentRepository.List(&LikeCommentQueryOption{CommentID: 1})
		assert.Equal(t,1, len(likes))

		likes = likeCommentRepository.List(&LikeCommentQueryOption{CommentID: 1, Username: "somebody"})
		assert.Equal(t,1, len(likes))

		// 존재하지 않는 유저
		likes = likeCommentRepository.List(&LikeCommentQueryOption{CommentID: 1, Username: "wizardofoz"})
		assert.Equal(t,0, len(likes))
	})

	t.Run("Somebody doesn't like comment 3.", func(t *testing.T) {
		likes := likeCommentRepository.List(&LikeCommentQueryOption{CommentID: 3})
		assert.Equal(t, 0, len(likes))
	})
}

func TestCommentRepositoryGorm_Get(t *testing.T) {
	comment, err := commentRepository.Get(1)

	assert.Nil(t, err)
	assert.NotNil(t, comment)
}

func TestCommentRepositoryGorm_List(t *testing.T) {
	comments := commentRepository.List(&CommentQueryOption{})
	for _, l := range comments{
		fmt.Println(l)
	}

	assert.NotEmpty(t, comments)
	assert.Equal(t, "jinsu", comments[0].Author.Username)
	assert.Equal(t, "jinsu", comments[1].Author.Username)
	t.Run("List comments written by somebody", func(t *testing.T) {
		comments = commentRepository.List(&CommentQueryOption{AuthorUsername: "somebody"})
		assert.GreaterOrEqual(t, len(comments), 1)
		assert.Equal(t, comments[0].AuthorUsername, "somebody")
		assert.Equal(t, comments[0].Author.Username, "somebody")
	})
}
