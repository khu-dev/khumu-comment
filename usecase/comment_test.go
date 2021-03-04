// 현재는 거의 usecase level에서 repository 계층까지 테스트하는 셈인데
// 좀 더 순수한 service 계층의 logic을 테스트할 수 있도록 바뀌었으면 좋겠다.
package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockCommentRepository *repository.MockCommentRepositoryInterface
	mockLikeCommentRepository *repository.MockLikeCommentRepositoryInterface
	commentUseCase     *CommentUseCase
	likeCommentUseCase *LikeCommentUseCase
	ctrl *gomock.Controller
)

func TestMain(m *testing.M) {
	commentUseCase = &CommentUseCase{
		Repository: mockCommentRepository,
		//LikeCommentRepository: likeCommentUseCase
	}

	m.Run()
}

// B는 Before each의 acronym
func BeforeCommentUseCaseTest(t *testing.T) {
	ctrl = gomock.NewController(t)

	mockCommentRepository = repository.NewMockCommentRepositoryInterface(ctrl)
	mockLikeCommentRepository = repository.NewMockLikeCommentRepositoryInterface(ctrl)
	commentUseCase = &CommentUseCase{
		Repository: mockCommentRepository,
		LikeCommentRepository: mockLikeCommentRepository,
	}
	//test.setUp(db
	// = make([]*model.Comment, 0)
	//for _, comment := range test.CommentsData {
	//	data = append(data, comment)
	//}
}

// A는 After each의 acronym
func A(tb testing.TB) {
	//test.CleanUp(db)
}

//func (m *MyMockedObject) DoSomething(number int) (bool, error) {
//
//  args := m.Called(number)
//  return args.Bool(0), args.Error(1)
//
//}

func TestCommentUseCase_Get(t *testing.T) {}

func TestLikeCommentUseCase_List(t *testing.T) {
	BeforeCommentUseCaseTest(t)
	test.SetUp()
	mockCommentRepository.EXPECT().List(gomock.Any()).Return(test.Comments).AnyTimes()
	mockLikeCommentRepository.EXPECT().List(gomock.Any()).Return([]*model.LikeComment{}).AnyTimes()
	comments, err := commentUseCase.List("jinsu", &repository.CommentQueryOption{})
	assert.NoError(t, err)
	for _, comment := range comments {

		assert.Equal(t, comment.AuthorUsername, comment.Author.Username)
		if comment.Kind == "named" {
			assert.NotEqual(t, AnonymousCommentUsername, comment.Author.Username)
			assert.NotEqual(t, AnonymousCommentNickname, comment.Author.Nickname)
			assert.NotEqual(t, DeletedCommentNickname, comment.Author.Username)
			assert.NotEqual(t, DeletedCommentUsername, comment.Author.Nickname)
		} else if comment.Kind == "anonymous" {
			assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
		} else if comment.Content == "deleted" {
			assert.Equal(t, DeletedCommentNickname, comment.Author.Username)
			assert.Equal(t, DeletedCommentContent, comment.Content)
		}
		if comment.ID == test.Comment1JinsuAnnonymous.ID {
			assert.True(t, comment.IsAuthor)
		}
	}
}


func TestCommentUseCase_Create(t *testing.T) {
	t.Run("My anonymous comment", func(t *testing.T) {
		BeforeCommentUseCaseTest(t)
		defer A(t)
		c := &model.Comment{
			Kind:           "anonymous",
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
}

func TestCommentUseCase_Update(t *testing.T) {
	BeforeCommentUseCaseTest(t)
	defer A(t)
	toUpdate, _ := commentUseCase.Get("", 1)
	updated, err := commentUseCase.Update("", toUpdate.ID, map[string]interface{}{
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
	BeforeCommentUseCaseTest(t)
	defer A(t)
	var err error
	var parent, child1, child2 *model.Comment
	t.Run("Setup", func(t *testing.T) {
		parent, err = commentUseCase.Create(&model.Comment{
			Author:  &model.KhumuUserSimple{Username: "jinsu"},
			Content: "A parent comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)

		child1, err = commentUseCase.Create(&model.Comment{
			Author:  &model.KhumuUserSimple{Username: "somebody"},
			Content: "The first child comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)
		assert.NotNil(t, child1)

		child2, err = commentUseCase.Create(&model.Comment{
			Author:  &model.KhumuUserSimple{Username: "somebody"},
			Content: "The second child comment to setup CommentRepositoryGorm_Delete.",
		})
		assert.Nil(t, err)
	})

	t.Run("두번째 대댓글 삭제", func(t *testing.T) {
		deleted, err := commentUseCase.Delete(child2.ID)
		assert.NoError(t, err)
		assert.NotNil(t, deleted)
		assert.Equal(t, child2.ID, deleted.ID)

		afterDeleted, err := commentUseCase.Get("", child2.ID)
		assert.NotNil(t, afterDeleted)
		assert.NoError(t, err)
		assert.Equal(t, "deleted", afterDeleted.State)
		assert.Equal(t, DeletedCommentContent, afterDeleted.Content)
	})

	// parent comment는 실제로 삭제되는 것이 아니라, kind가 deleted 로 변경될 뿐.
	t.Run("부모 댓글 삭제", func(t *testing.T) {
		assert.Equal(t, "exists", parent.State)
		deletedParent, err := commentUseCase.Delete(parent.ID)
		assert.NoError(t, err)
		assert.NotNil(t, deletedParent)
		assert.Equal(t, parent.ID, deletedParent.ID)

		// 삭제된 댓글의 작성자는 무언가를 통해 익명처리가 되어야함.
		afterDeleted, err := commentUseCase.Get("", parent.ID)
		assert.NotNil(t, afterDeleted)
		assert.NoError(t, err)
		assert.Equal(t, "deleted", afterDeleted.State)
		assert.Equal(t, DeletedCommentContent, afterDeleted.Content)
	})
}

// 시나리오
// somebody가 jinsu의 댓글인 1번 댓글을 좋아요.
//func TestLikeCommentUseCase_Toggle(t *testing.T) {
//	BeforeCommentUseCaseTest(t)
//	defer A(t)
//	t.Run("Somebody toggle(create&delete) jinsu's comment", func(t *testing.T) {
//		commentID := test.CommentsData["JinsuNamedComment"].ID
//		// toggle to create
//		func() {
//			created, err := likeCommentUseCase.Toggle(
//				&model.LikeComment{
//					CommentID: commentID,
//					Username:  test.UsersData["Somebody"].Username,
//				})
//			assert.Nil(t, err)
//			assert.True(t, created)
//		}()
//		// toggle to delete
//		func() {
//			deleted, err := likeCommentUseCase.Toggle(
//				&model.LikeComment{
//					CommentID: commentID,
//					Username:  test.UsersData["Somebody"].Username,
//				})
//			assert.Nil(t, err)
//			assert.False(t, deleted)
//		}()
//
//		// toggle to create again
//		func() {
//			created, err := likeCommentUseCase.Toggle(
//				&model.LikeComment{
//					CommentID: commentID,
//					Username:  test.UsersData["Somebody"].Username,
//				})
//			assert.Nil(t, err)
//			assert.True(t, created)
//		}()
//	})
//
//	t.Run("jinsu likes jinsu's comment", func(t *testing.T) {
//		created, err := likeCommentUseCase.Toggle(
//			&model.LikeComment{
//				CommentID: test.CommentsData["JinsuNamedComment"].ID,
//				Username:  test.UsersData["Jinsu"].Username,
//			})
//		assert.NotNil(t, err)
//		assert.False(t, created)
//	})
//}
//
//// 시나리오
//// 기존에 somebody가 jinsu의 1번 댓글을 좋아요.
//// => 좋아요 1개
//func TestCommentUseCase_List(t *testing.T) {
//	t.Run("Article에 대한 LikeCount 필드", func(t *testing.T) {
//		BeforeCommentUseCaseTest(t)
//		defer A(t)
//		// create like count for test
//		func() {
//			// 1번 somebody가 jinsu의 코멘트에 좋아요
//			_, err := likeCommentUseCase.Toggle(&model.LikeComment{Username: "somebody", CommentID: test.CommentsData["JinsuNamedComment"].ID})
//			assert.NoError(t, err)
//		}()
//		comments, err := commentUseCase.List("jinsu", nil)
//		assert.NoError(t, err)
//		foundDesiredComment := false
//		for _, c := range comments {
//			if c.ID == test.CommentsData["JinsuNamedComment"].ID {
//				foundDesiredComment = true
//				assert.False(t, c.Liked)
//				assert.Equal(t, 1, c.LikeCommentCount)
//			}
//		}
//		assert.True(t, foundDesiredComment)
//	})
//	t.Run("jinsu 자기 댓글", func(t *testing.T) {
//		BeforeCommentUseCaseTest(t)
//		defer A(t)
//		commentID := test.CommentsData["JinsuNamedComment"].ID
//		comments, err := commentUseCase.List("jinsu", &repository.CommentQueryOption{CommentID: commentID})
//		assert.NoError(t, err)
//		logrus.Error(comments)
//		c := comments[0]
//		assert.Equal(t, "named", c.Kind)
//		assert.Equal(t, "jinsu", c.AuthorUsername)
//		assert.Equal(t, "jinsu", c.Author.Username)
//		assert.False(t, c.Liked) // 자신의 코멘트에 대한 liked
//		assert.True(t, c.IsAuthor)
//	})
//	t.Run("somebody가 jinsu 댓글", func(t *testing.T) {
//		BeforeCommentUseCaseTest(t)
//		defer A(t)
//		commentID := test.CommentsData["JinsuAnonymousComment"].ID
//		comments, err := commentUseCase.List("somebody", &repository.CommentQueryOption{CommentID: commentID})
//		assert.NoError(t, err)
//		c := comments[0]
//		assert.Equal(t, "anonymous", c.Kind)
//		assert.Equal(t, AnonymousCommentUsername, c.AuthorUsername)
//		assert.Equal(t, AnonymousCommentUsername, c.Author.Username)
//		assert.Equal(t, AnonymousCommentNickname, c.Author.Nickname)
//		assert.False(t, c.IsAuthor)
//	})
//}
