// 현재는 거의 usecase level에서 repository 계층까지 테스트하는 셈인데
// 좀 더 순수한 service 계층의 logic을 테스트할 수 있도록 바뀌었으면 좋겠다.
package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/enttest"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	repo       *ent.Client
	mockSnsClient               *external.MockSnsClient
	commentUseCase              *CommentUseCase
	likeCommentUseCase          *LikeCommentUseCase
	ctrl                        *gomock.Controller
)

// B는 Before each의 acronym
func BeforeCommentUseCaseTest(t *testing.T) {
	ctrl = gomock.NewController(t)
	repo = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	mockSnsClient = external.NewMockSnsClient(ctrl)
	commentUseCase = &CommentUseCase{
		Repo: repo,
		SnsClient:             mockSnsClient,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repo: repo,
	}

	mockSnsClient.EXPECT().PublishMessage(gomock.Any()).DoAndReturn(
		func(message interface{}) {
			t.Log("그냥 테스트라서 푸시 알림 패스")
		}).AnyTimes()

	test.SetUp(repo)
}

// A는 After each의 acronym
func A(tb testing.TB) {
	repo.Close()
}

//func (m *MyMockedObject) DoSomething(number int) (bool, error) {
//
//  args := m.Called(number)
//  return args.Bool(0), args.Error(1)
//
//}

func TestCommentUseCase_Create(t *testing.T) {
	t.Run("My anonymous article comment", func(t *testing.T) {
		BeforeCommentUseCaseTest(t)
		defer A(t)

	tmp, _ := commentUseCase.Create(&data.CommentInput{
		Author: test.UserJinsu.ID,
		Article: &test.Articles[0].ID,
		Content: "테스트 댓글",
	})

	comment, err := commentUseCase.Get(test.UserJinsu.ID, tmp.Id)
	assert.NoError(t, err)
	// 기본적으로는 익명 댓글임.
	assert.Equal(t, AnonymousCommentUsername, comment.Author.Username)
	assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
	assert.Equal(t, "테스트 댓글", comment.Content)
	})
	//
	//t.Run("My anonymous study article comment", func(t *testing.T) {
	//	BeforeCommentUseCaseTest(t)
	//	defer A(t)
	//	c := &_model.Comment{
	//		Kind:           "anonymous",
	//		State:          "exists",
	//		AuthorUsername: "jinsu",
	//		StudyArticleID: null.Int{sql.NullInt64{Int64: 1}},
	//		Content:        "새로운 테스트 댓글",
	//	}
	//
	//	mockCommentRepository.EXPECT().Create(gomock.Any()).DoAndReturn(
	//		func(comment *_model.Comment) (*_model.Comment, error) {
	//			c := *comment
	//			c.Author = &_model.KhumuUserSimple{Username: c.AuthorUsername}
	//			return &c, nil
	//		},
	//	)
	//
	//	newComment, err := commentUseCase.Create(c)
	//	assert.Nil(t, err)
	//	assert.NotNil(t, newComment)
	//	assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
	//	assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
	//	assert.Equal(t, c.Content, newComment.Content)
	//})
}

func TestCommentUseCase_Get(t *testing.T) {
	BeforeCommentUseCaseTest(t)
	defer A(t)

	articles, err := repo.Article.Query().All(context.Background())
	assert.NoError(t, err)
	t.Log(articles)

	tmp, _ := commentUseCase.Create(&data.CommentInput{
		Author: test.UserJinsu.ID,
		Article: &test.Articles[0].ID,
		Content: "테스트 댓글",
	})

	comment, err := commentUseCase.Get(test.UserJinsu.ID, tmp.Id)
	assert.NoError(t, err)
	// 기본적으로는 익명 댓글임.
	assert.Equal(t, AnonymousCommentUsername, comment.Author.Username)
	assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)

}

//func TestLikeCommentUseCase_List(t *testing.T) {
//	BeforeCommentUseCaseTest(t)
//	defer A(t)
//	mockCommentRepository.EXPECT().List(gomock.Any()).Return(test.Comments).AnyTimes()
//	mockLikeCommentRepository.EXPECT().List(gomock.Any()).Return([]*_model.LikeComment{}).AnyTimes()
//	comments, err := commentUseCase.List("jinsu", &_repository.CommentQueryOption{})
//	assert.NoError(t, err)
//	for _, comment := range comments {
//
//		assert.Equal(t, comment.AuthorUsername, comment.Author.Username)
//		if comment.Kind == "named" {
//			assert.NotEqual(t, AnonymousCommentUsername, comment.Author.Username)
//			assert.NotEqual(t, AnonymousCommentNickname, comment.Author.Nickname)
//			assert.NotEqual(t, DeletedCommentNickname, comment.Author.Username)
//			assert.NotEqual(t, DeletedCommentUsername, comment.Author.Nickname)
//		} else if comment.Kind == "anonymous" {
//			assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
//		} else if comment.Content == "deleted" {
//			assert.Equal(t, DeletedCommentNickname, comment.Author.Username)
//			assert.Equal(t, DeletedCommentContent, comment.Content)
//		}
//		if comment.ID == test.Comment1JinsuAnnonymous.ID {
//			assert.True(t, comment.IsAuthor)
//		}
//	}
//}
//

//
//func TestCommentUseCase_Update(t *testing.T) {
//	BeforeCommentUseCaseTest(t)
//	defer A(t)
//	// Update는 대부분 repository 계층에서만 확인해도 될 듯.
//
//	//before := *test.Comment1JinsuAnnonymous
//	//updateData := map[string]interface{}{
//	//	"content": "수정된 1번 코멘트입니다.",
//	//	"kind": "named",
//	//}
//	//
//	//after, err := commentUseCase.Update("jinsu", before.ID, updateData)
//	//assert.NoError(t, err)
//	//assert.Equal(t, "수정된 1번 코멘트입니다.", after.Content)
//	//assert.Equal(t, "named", after.Kind)
//}
//
//// Delete 작업 자체는 Repository 계층에서 이루어져야할 듯하고
//// service 계층에서는 간단한 테스트만 수행한다.
//func TestCommentUseCase_삭제된_댓글에_대한_조회(t *testing.T) {
//	BeforeCommentUseCaseTest(t)
//	defer A(t)
//	test.Comment1JinsuAnnonymous.State = "deleted"
//	mockCommentRepository.EXPECT().Get(gomock.Any()).Return(test.Comment1JinsuAnnonymous, nil)
//	mockLikeCommentRepository.EXPECT().List(gomock.Any()).Return([]*_model.LikeComment{}).AnyTimes()
//	comment, err := commentUseCase.Get("Anything will be fine", 1)
//	assert.NoError(t, err)
//	assert.Equal(t, DeletedCommentUsername, comment.AuthorUsername)
//	assert.Equal(t, DeletedCommentUsername, comment.Author.Username)
//	assert.Equal(t, DeletedCommentNickname, comment.Author.Nickname)
//	assert.Equal(t, DeletedCommentContent, comment.Content)
//}
//
//func TestLikeCommentUseCase_Toggle(t *testing.T) {
//	BeforeCommentUseCaseTest(t)
//	defer A(t)
//	commentID := test.Comment1JinsuAnnonymous.ID
//	mockCommentRepository.EXPECT().Get(gomock.Any()).Return(test.Comment1JinsuAnnonymous, nil).AnyTimes()
//	likeComments := make([]*_model.LikeComment, 0)
//	mockLikeCommentRepository.EXPECT().Create(gomock.Any()).DoAndReturn(func(c *_model.LikeComment) (*_model.LikeComment, error) {
//		likeComments = append(likeComments, c)
//		return c, nil
//	}).AnyTimes()
//	// mock으로 그냥 한 칸 줄이기만함.
//	mockLikeCommentRepository.EXPECT().Delete(gomock.Any()).DoAndReturn(func(id int) (*_model.LikeComment, error) {
//		deleted := likeComments[id]
//		likeComments = append(likeComments[:id], likeComments[id+1:]...)
//		return deleted, nil
//	}).AnyTimes()
//	// 거의 로직을 구현해버렸네.....
//	mockLikeCommentRepository.EXPECT().List(gomock.Any()).DoAndReturn(func(option *_repository.LikeCommentQueryOption) []*_model.LikeComment {
//		answers := make([]*_model.LikeComment, 0)
//		// 쿼리 조건 판별
//		if option.Username != "" && option.CommentID != 0 {
//			for _, like := range likeComments {
//				if like.Username == option.Username && like.CommentID == option.CommentID {
//					answers = append(answers, like)
//				}
//			}
//		} else if option.Username != "" {
//			for _, like := range likeComments {
//				if like.Username == option.Username {
//					answers = append(answers, like)
//				}
//			}
//		} else if option.CommentID != 0 {
//			for _, like := range likeComments {
//				if like.CommentID == option.CommentID {
//					answers = append(answers, like)
//				}
//			}
//		} else {
//			answers = likeComments
//		}
//		return answers
//	}).AnyTimes()
//
//	t.Run("Somebody toggle(create&delete) jinsu's comment", func(t *testing.T) {
//
//		// toggle to create
//		func() {
//			created, err := likeCommentUseCase.Toggle(
//				&_model.LikeComment{
//					CommentID: commentID,
//					Username:  "somebody",
//				})
//			assert.Nil(t, err)
//			assert.True(t, created)
//		}()
//		// toggle to delete
//		func() {
//			deleted, err := likeCommentUseCase.Toggle(
//				&_model.LikeComment{
//					CommentID: commentID,
//					Username:  "somebody",
//				})
//			assert.Nil(t, err)
//			assert.False(t, deleted)
//		}()
//
//		// toggle to create again
//		func() {
//			created, err := likeCommentUseCase.Toggle(
//				&_model.LikeComment{
//					CommentID: commentID,
//					Username:  "somebody",
//				})
//			assert.Nil(t, err)
//			assert.True(t, created)
//		}()
//
//		// 자기 댓글은 좋아요 불가능.
//		func() {
//			created, err := likeCommentUseCase.Toggle(
//				&_model.LikeComment{
//					CommentID: test.Comment1JinsuAnnonymous.ID,
//					Username:  test.UserJinsu.Username,
//				})
//			assert.NotNil(t, err)
//			assert.False(t, created)
//		}()
//	})
//}
