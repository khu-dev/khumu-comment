// repository 계층과 분리된 채 usecase 계층만의 테스트를 진행
// repository를 간단히 mock 한 뒤 기명, 익명 kind의 comment 들이
// desired한 방향으로 얻어지는 지 확인해본다.
package usecase

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"testing"
)

// 언젠가 모킹을 사용할 것이라면 이 타입을 구현. 메소드는 현재 사용하지 않으므로 주석처리했다.
type CommentRepositoryMock struct{}
//func (r *CommentRepositoryMock) Create(comment *model.Comment) error {
//	commentsData = append(commentsData, comment)
//	return nil
//}
//
//// QyeryOption기능은 제외하고 mock
//func (r *CommentRepositoryMock) List(opt *repository.CommentQueryOption) []*model.Comment {
//	return commentsData
//}
//
//func (r *CommentRepositoryMock) Get(id int) *model.Comment {
//	for _, comment := range commentsData {
//		if int(comment.ArticleID) == id {
//			return comment
//		}
//	}
//	return nil
//}

var (
	commentUseCase CommentUseCaseInterface
)

func TestInit(t *testing.T) {
	// build container
	cont := dig.New()
	err := cont.Provide(repository.NewTestGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewUserRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(NewCommentUseCase)
	assert.Nil(t, err)

	err = cont.Invoke(func(uc CommentUseCaseInterface) {
		commentUseCase = uc
	})

	assert.Nil(t, err)

	t.Run("Create a user jinsu to preload in list comment", func(t *testing.T) {
		user := &model.KhumuUserSimple{Username: "jinsu", Type: "active"}
		err = cont.Invoke(func(db *gorm.DB){
			dbErr := db.Create(&user).Error
			assert.Nil(t, dbErr)
			assert.Equal(t, "jinsu", user.Username)
		})
		assert.Nil(t, err)
	})

	t.Run("Create a user somebody who is not me to preload in list comment", func(t *testing.T) {
		user := &model.KhumuUserSimple{Username: "somebody", Type: "active"}
		err = cont.Invoke(func(db *gorm.DB){
			dbErr := db.Create(&user).Error
			assert.Nil(t, dbErr)
			assert.Equal(t, "somebody", user.Username)
		})
		assert.Nil(t, err)
	})

	// 내가 사용할 원본 데이터가 잘 만들어져있는가
	assert.GreaterOrEqual(t, len(test.CommentsData), 3) // e1 >= 3
}

func createCommentData(t *testing.T){
	assert.Len(t, test.CommentsData, 3)
}

func TestCommentUseCase_Create(t *testing.T){
	t.Run("My anonymous comment", func(t *testing.T){
		c := test.CommentsData[0] // 0 번 인덱스는 익명 댓글
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("My named comment", func(t *testing.T){
		c := test.CommentsData[1] // 1번 인덱스는 기명 댓글
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("Others anonymous comment", func(t *testing.T){
		c := test.CommentsData[2] // 1번 인덱스는 기명 댓글
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})
}

func TestCommentUseCase_List(t *testing.T) {

	resultComments, err := commentUseCase.List("jinsu", &repository.CommentQueryOption{})
	assert.Nil(t, err)

	t.Run("My anonymous comment", func(t *testing.T) {
		c := resultComments[0]
		assert.Equal(t, c.Kind, "anonymous")
		assert.Equal(t, "jinsu", c.AuthorUsername)
		assert.Equal(t, "jinsu", c.Author.Username)
	})
	t.Run("My named comment", func(t *testing.T) {
		c := resultComments[1]
		assert.Equal(t, "named", c.Kind)
		assert.Equal(t, "jinsu", c.AuthorUsername)
		assert.Equal(t, "jinsu", c.Author.Username)
	})
	t.Run("Others anonymous comment", func(t *testing.T) {
		c := resultComments[2]
		assert.Equal(t, "anonymous", c.Kind)
		assert.Equal(t, "익명", c.AuthorUsername)
		assert.Equal(t, "익명", c.Author.Username)
	})
}


