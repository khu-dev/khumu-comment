// repository 계층과 분리된 채 usecase 계층만의 테스트를 진행
// repository를 간단히 mock 한 뒤 기명, 익명 kind의 comment 들이
// desired한 방향으로 얻어지는 지 확인해본다.
package usecase

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"testing"
)

// 언젠가 모킹을 사용할 것이라면 이 타입을 구현. 메소드는 현재 사용하지 않으므로 주석처리했다.
//type CommentRepositoryMock struct{}
//func (r *CommentRepositoryMock) Create(comment *model.Comment) (*model.Comment, error){
//	return nil, nil
//}
//func (r *CommentRepositoryMock) List(opt *repository.CommentQueryOption) []*model.Comment{
//	return nil
//}
//func (r *CommentRepositoryMock) Get(id int) (*model.Comment, error){
//	return nil, nil
//}
//func (r *CommentRepositoryMock) Update(id int, opt map[string]interface{}) (*model.Comment, error){
//	return nil, nil
//}
//func (r *CommentRepositoryMock) Delete(id int) (*model.Comment, error){
//	return nil, nil
//}

var (
	db *gorm.DB
	commentUseCase CommentUseCaseInterface
	likeCommentUseCase LikeCommentUseCaseInterface
)

func TestMain(m *testing.M){
	cont := BuildContainer()
	err := cont.Invoke(func(conn *gorm.DB, cuc CommentUseCaseInterface, lcuc LikeCommentUseCaseInterface) {
		db = conn
		commentUseCase = cuc
		likeCommentUseCase = lcuc
	})
	if err != nil{
		logrus.Fatal(err)
	}

	m.Run()
}


// B는 Before each의 acronym
func B(tb testing.TB){
	test.SetUp(db)
}

// A는 After each의 acronym
func A(tb testing.TB){
	test.CleanUp(db)
}

func BuildContainer() (*dig.Container){
	cont := dig.New()
	err := cont.Provide(repository.NewTestGorm)
	if err != nil{
		logrus.Fatal(err)
	}
	err = cont.Provide(repository.NewCommentRepositoryGorm)
	if err != nil{
		logrus.Fatal(err)
	}
	err = cont.Provide(repository.NewLikeCommentRepositoryGorm)
	if err != nil{
		logrus.Fatal(err)
	}
	err = cont.Provide(NewCommentUseCase)
	if err != nil{
		logrus.Fatal(err)
	}
	err = cont.Provide(NewLikeCommentUseCase)
	if err != nil{
		logrus.Fatal(err)
	}
	return cont
}

func TestCommentUseCase_Create(t *testing.T){
	t.Run("My anonymous comment", func(t *testing.T){
		B(t)
		defer A(t)
		c := &model.Comment{
			Kind: "anonymous",
			State: "exists",
			AuthorUsername: "jinsu",
			ArticleID: 1,
			Content: "새로운 테스트 댓글",
		}
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("My named comment", func(t *testing.T){
		B(t)
		defer A(t)
		c := &model.Comment{
			Kind:           "named",
			State:          "exists",
			AuthorUsername: "jinsu",
			ArticleID:      1,
			Content:        "새로운 테스트 댓글",
		}
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("Others anonymous comment", func(t *testing.T){
		B(t)
		defer A(t)
		c := &model.Comment{
			Kind: "named",
			State: "exists",
			AuthorUsername: "jinsu",
			ArticleID: 1,
			Content: "새로운 테스트 댓글",
		}
		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})
}

func TestCommentUseCase_Get(t *testing.T) {
	B(t)
	defer A(t)
}

func TestCommentUseCase_Update(t *testing.T) {
	B(t)
	defer A(t)
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
	B(t)
	defer A(t)
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
		assert.NotNil(t, child1)

		child2, err = commentUseCase.Create(&model.Comment{
			Author: &model.KhumuUserSimple{Username: "somebody"},
			Content: "The second child comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)
	})

	t.Run("두번째 대댓글 삭제", func(t *testing.T) {
		deleted, err := commentUseCase.Delete(child2.ID)
		assert.NoError(t, err)
		assert.NotNil(t, deleted)
		assert.Equal(t, child2.ID, deleted.ID)

		afterDeleted, err := commentUseCase.Get(child2.ID)
		assert.NotNil(t, afterDeleted)
		assert.NoError(t, err)
		assert.Equal(t, "deleted", afterDeleted.State)
		assert.Equal(t, model.DeletedCommentContent, afterDeleted.Content)
	})

	// parent comment는 실제로 삭제되는 것이 아니라, kind가 deleted 로 변경될 뿐.
	t.Run("부모 댓글 삭제", func(t *testing.T) {
		assert.Equal(t, "exists", parent.State)
		deletedParent, err := commentUseCase.Delete(parent.ID)
		assert.NoError(t, err)
		assert.NotNil(t, deletedParent)
		assert.Equal(t, parent.ID, deletedParent.ID)

		// 삭제된 댓글의 작성자는 무언가를 통해 익명처리가 되어야함.
		afterDeleted, err := commentUseCase.Get(parent.ID)
		assert.NotNil(t, afterDeleted)
		assert.NoError(t, err)
		assert.Equal(t, "deleted", afterDeleted.State)
		assert.Equal(t, model.DeletedCommentContent, afterDeleted.Content)
	})
}

func TestLikeCommentUseCase_List(t *testing.T) {
	// Nothing.
	B(t)
	defer A(t)
}


// 시나리오
// somebody가 jinsu의 댓글인 1번 댓글을 좋아요.
func TestLikeCommentUseCase_Toggle(t *testing.T) {
	B(t)
	defer A(t)
	t.Run("Somebody toggle(create&delete) jinsu's comment", func(t *testing.T) {
		commentID := test.CommentsData["JinsuNamedComment"].ID
		// toggle to create
		func(){
			created, err := likeCommentUseCase.Toggle(
			&model.LikeComment{
				CommentID: commentID,
				Username: test.UsersData["Somebody"].Username,
			})
			assert.Nil(t, err)
			assert.True(t, created)
		}()
		// toggle to delete
		func(){
			deleted, err := likeCommentUseCase.Toggle(
			&model.LikeComment{
				CommentID: commentID,
				Username: test.UsersData["Somebody"].Username,
			})
			assert.Nil(t, err)
			assert.False(t, deleted)
		}()

		// toggle to create again
		func(){
			created, err := likeCommentUseCase.Toggle(
			&model.LikeComment{
				CommentID: commentID,
				Username: test.UsersData["Somebody"].Username,
			})
			assert.Nil(t, err)
			assert.True(t, created)
		}()
	})

	t.Run("jinsu likes jinsu's comment", func(t *testing.T){
		created, err := likeCommentUseCase.Toggle(
		&model.LikeComment{
			CommentID: test.CommentsData["JinsuNamedComment"].ID,
			Username: test.UsersData["Jinsu"].Username,
		})
		assert.NotNil(t, err)
		assert.False(t, created)
	})
}

// 시나리오
// 기존에 somebody가 jinsu의 1번 댓글을 좋아요.
// => 좋아요 1개
func TestCommentUseCase_List(t *testing.T) {
	t.Run("Article에 대한 LikeCount 필드", func(t *testing.T) {
		B(t)
		defer A(t)
		// create like count for test
		func(){
			// 1번 somebody가 jinsu의 코멘트에 좋아요
			_, err := likeCommentUseCase.Toggle(&model.LikeComment{Username: "somebody", CommentID: test.CommentsData["JinsuNamedComment"].ID})
			assert.NoError(t, err)
		}()
		comments, err := commentUseCase.List("jinsu", nil)
		assert.NoError(t, err)
		foundDesiredComment := false
		for _, c := range comments{
			if c.ID == test.CommentsData["JinsuNamedComment"].ID{
				foundDesiredComment = true
				assert.False(t, c.Liked)
				assert.Equal(t, 1, c.LikeCommentCount)
			}
		}
		assert.True(t, foundDesiredComment)
	})
	t.Run("jinsu 자기 댓글", func(t *testing.T) {
		B(t)
		defer A(t)
		commentID := test.CommentsData["JinsuNamedComment"].ID
		comments, err := commentUseCase.List("jinsu", &repository.CommentQueryOption{CommentID: commentID})
		assert.NoError(t, err)
		logrus.Error(comments)
		c := comments[0]
		assert.Equal(t, "named", c.Kind)
		assert.Equal(t, "jinsu", c.AuthorUsername)
		assert.Equal(t, "jinsu", c.Author.Username)
		assert.False(t, c.Liked) // 자신의 코멘트에 대한 liked
	})
	t.Run("somebody 자기 댓글", func(t *testing.T) {
		B(t)
		defer A(t)
		commentID := test.CommentsData["SomebodyAnonymousComment"].ID
		comments, err := commentUseCase.List("somebody", &repository.CommentQueryOption{CommentID: commentID})
		assert.NoError(t, err)
		c := comments[0]
		assert.Equal(t, "anonymous", c.Kind)
		assert.Equal(t, "somebody", c.AuthorUsername)
		assert.Equal(t, "somebody", c.Author.Username)
		//isAuthor 필드 도입
	})
}