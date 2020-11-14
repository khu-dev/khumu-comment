// repository 계층과 분리된 채 usecase 계층만의 테스트를 진행
// repository를 간단히 mock 한 뒤 기명, 익명 kind의 comment 들이
// desired한 방향으로 얻어지는 지 확인해본다.
package usecase

import (
	"fmt"
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
	likeCommentUseCase LikeCommentUseCaseInterface
)

func TestSetUp(t *testing.T) {
	// build container
	cont := dig.New()
	err := cont.Provide(repository.NewTestGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewLikeCommentRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(repository.NewUserRepositoryGorm)
	assert.Nil(t, err)

	err = cont.Provide(NewCommentUseCase)
	assert.Nil(t, err)
	err = cont.Invoke(func(uc CommentUseCaseInterface) {
		commentUseCase = uc
	})
	assert.Nil(t, err)

	err = cont.Provide(NewLikeCommentUseCase)
	assert.Nil(t, err)
	err = cont.Invoke(func(uc LikeCommentUseCaseInterface) {
		likeCommentUseCase = uc
	})
	assert.Nil(t, err)

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

	// 내가 사용할 원본 데이터가 잘 만들어져있는가
	assert.GreaterOrEqual(t, len(test.CommentsData), 3) // e1 >= 3
}

func TestCommentUseCase_Create(t *testing.T){
	t.Run("My anonymous comment", func(t *testing.T){
		c := test.CommentsData["JinsuAnonymousComment"] // 0 번 인덱스는 익명 댓글
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("My named comment", func(t *testing.T){
		c := test.CommentsData["JinsuNamedComment"] // 1번 인덱스는 기명 댓글
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("Others anonymous comment", func(t *testing.T){
		c := test.CommentsData["SomebodyAnonymousComment"] // 1번 인덱스는 기명 댓글
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})
}

func TestCommentUseCase_Get(t *testing.T) {

}

func TestCommentUseCase_Update(t *testing.T) {
	toUpdate, _ := commentUseCase.Get(1)
	updated, err := commentUseCase.Update(toUpdate.ID, map[string]interface{}{
		"content": "수정된 1번 코멘트입니다.",
	})
	assert.NoError(t, err)
	assert.NotEqual(t, toUpdate, updated)
	assert.Equal(t, "수정된 1번 코멘트입니다.", updated.Content)
}

// setup에서 parent comment 1개, 그 comment를 참조하는 child comment 2개 생성
// 이후 child comment 삭제 => 실제로 DB에서 삭제
// parent comment 삭제 시 parent comment의 state는 deleted
// 남아있던 child comment는 여전히 parent를 참조.
func TestCommentUseCase_Delete(t *testing.T) {
	var err error
	var parent, child1, child2 *model.Comment
	t.Run("Setup", func(t *testing.T) {
		parent, err = commentUseCase.Create(&model.Comment{
			Author: &model.KhumuUserSimple{Username: "jinsu"},
			Content: "A parent comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)

		child1, err = commentUseCase.Create(&model.Comment{
			Author: &model.KhumuUserSimple{Username: "somebody"},
			Content: "The first child comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)

		child2, err = commentUseCase.Create(&model.Comment{
			Author: &model.KhumuUserSimple{Username: "somebody"},
			Content: "The second child comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)
	})
	fmt.Println(parent,child1,child2)

	t.Run("The second child comment", func(t *testing.T) {
		deleted, err := commentUseCase.Delete(child2.ID)
		assert.Nil(t, err)
		assert.NotNil(t, deleted)

		_, err = commentUseCase.Get(deleted.ID)
		assert.NoError(t, err)
	})

	// parent comment는 실제로 삭제되는 것이 아니라, kind가 deleted 로 변경될 뿐.
	t.Run("The parent comment", func(t *testing.T) {
		assert.Equal(t, "exists", parent.State)
		updatedParent, err := commentUseCase.Update(parent.ID, map[string]interface{}{
			"state": "deleted",
		})
		assert.NoError(t, err)
		assert.Equal(t, "deleted", updatedParent.State)
		// 삭제된 댓글의 작성자는 무언가를 통해 익명처리가 되어야함.
		assert.NotEqual(t, parent.AuthorUsername, updatedParent.AuthorUsername)
	})
}

func TestLikeCommentUseCase_List(t *testing.T) {
	// Nothing.
}

func TestLikeCommentUseCase_Toggle(t *testing.T) {
	t.Run("Somebody likes jinsu's comment", func(t *testing.T) {
		commentID := 1
		created, err := likeCommentUseCase.Toggle(
		&model.LikeComment{
			CommentID: commentID,
			Username: test.UsersData["Somebody"].Username,
		})
		assert.Nil(t, err)
		assert.True(t, created)
	})

	t.Run("Somebody likes jinsu's comment again so delete", func(t *testing.T) {
		commentID := 1
		created, err := likeCommentUseCase.Toggle(
		&model.LikeComment{
			CommentID: commentID,
			Username: test.UsersData["Somebody"].Username,
		})
		assert.Nil(t, err)
		assert.False(t, created)
	})

	t.Run("jinsu likes jinsu's comment", func(t *testing.T){
		created, err := likeCommentUseCase.Toggle(
		&model.LikeComment{
			CommentID: 1,
			Username: test.UsersData["Jinsu"].Username,
		})
		assert.NotNil(t, err)
		assert.False(t, created)
	})

	t.Run("Bad request to create a like comment", func(t *testing.T){
	})

}

func TestCommentUseCase_List(t *testing.T) {
	var resultComments []*model.Comment
	t.Run("Set up", func(t *testing.T) {
		created, err := likeCommentUseCase.Toggle(
		&model.LikeComment{
			CommentID: 1,
			Username: test.UsersData["Puppy"].Username,
		})
		assert.Nil(t, err)
		assert.True(t, created)

		resultComments, err = commentUseCase.List("jinsu", &repository.CommentQueryOption{})
		assert.Nil(t, err)
	})

	for _, c := range resultComments{
		fmt.Println(c.ID, c.Author.Username)
	}

	t.Run("Jinsu's anonymous comment", func(t *testing.T) {
		c := resultComments[0]
		assert.Equal(t, c.Kind, "anonymous")
		assert.Equal(t, "jinsu", c.AuthorUsername)
		assert.Equal(t, "jinsu", c.Author.Username)
		assert.Equal(t, 1, c.LikeCommentCount)
	})
	t.Run("Jinsu's named comment", func(t *testing.T) {
		c := resultComments[1]
		assert.Equal(t, "named", c.Kind)
		assert.Equal(t, "jinsu", c.AuthorUsername)
		assert.Equal(t, "jinsu", c.Author.Username)
		assert.Equal(t, 0, c.LikeCommentCount)
	})
	t.Run("Somebody's anonymous comment", func(t *testing.T) {
		c := resultComments[2]
		assert.Equal(t, "anonymous", c.Kind)
		assert.Equal(t, "익명", c.AuthorUsername)
		assert.Equal(t, "익명", c.Author.Username)
		assert.Equal(t, 0, c.LikeCommentCount)
	})
}
