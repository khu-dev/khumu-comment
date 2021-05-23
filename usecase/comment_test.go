// 현재는 거의 usecase level에서 repository 계층까지 테스트하는 셈인데
// 좀 더 순수한 service 계층의 logic을 테스트할 수 있도록 바뀌었으면 좋겠다.
package usecase

import (
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

var (
	mockCommentRepository       *repository.MockCommentRepositoryInterface
	mockLikeCommentRepository   *repository.MockLikeCommentRepositoryInterface
	mockSnsClient               *external.MockSnsClient
	redisEventMessageRepository *repository.RedisEventMessageRepository
	commentUseCase              *CommentUseCase
	likeCommentUseCase          *LikeCommentUseCase
	ctrl                        *gomock.Controller
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
	test.SetUp()
	ctrl = gomock.NewController(t)

	mockCommentRepository = repository.NewMockCommentRepositoryInterface(ctrl)
	mockLikeCommentRepository = repository.NewMockLikeCommentRepositoryInterface(ctrl)
	mockSnsClient = external.NewMockSnsClient(ctrl)
	commentUseCase = &CommentUseCase{
		Repository:            mockCommentRepository,
		LikeCommentRepository: mockLikeCommentRepository,
		SnsClient:             mockSnsClient,
	}
	likeCommentUseCase = &LikeCommentUseCase{
		Repository:        mockLikeCommentRepository,
		CommentRepository: mockCommentRepository,
	}

	mockSnsClient.EXPECT().PublishMessage(gomock.Any()).DoAndReturn(
		func(message interface{}) {
			t.Log("그냥 테스트라서 푸시 알림 패스")
		}).AnyTimes()
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
	defer A(t)
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
	t.Run("My anonymous article comment", func(t *testing.T) {
		BeforeCommentUseCaseTest(t)
		defer A(t)
		c := &model.Comment{
			Kind:           "anonymous",
			State:          "exists",
			AuthorUsername: "jinsu",
			ArticleID:      null.Int{sql.NullInt64{1, true}},
			Content:        "새로운 테스트 댓글",
		}

		mockCommentRepository.EXPECT().Create(gomock.Any()).DoAndReturn(
			func(comment *model.Comment) (*model.Comment, error) {
				c := *comment
				c.Author = &model.KhumuUserSimple{Username: c.AuthorUsername}
				return &c, nil
			},
		)

		newComment, err := commentUseCase.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, c.AuthorUsername, newComment.AuthorUsername)
		assert.Equal(t, c.AuthorUsername, newComment.Author.Username)
		assert.Equal(t, c.Content, newComment.Content)
	})

	t.Run("My anonymous study article comment", func(t *testing.T) {
		BeforeCommentUseCaseTest(t)
		defer A(t)
		c := &model.Comment{
			Kind:           "anonymous",
			State:          "exists",
			AuthorUsername: "jinsu",
			StudyArticleID: null.Int{sql.NullInt64{Int64: 1}},
			Content:        "새로운 테스트 댓글",
		}

		mockCommentRepository.EXPECT().Create(gomock.Any()).DoAndReturn(
			func(comment *model.Comment) (*model.Comment, error) {
				c := *comment
				c.Author = &model.KhumuUserSimple{Username: c.AuthorUsername}
				return &c, nil
			},
		)

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
	// Update는 대부분 repository 계층에서만 확인해도 될 듯.

	//before := *test.Comment1JinsuAnnonymous
	//updateData := map[string]interface{}{
	//	"content": "수정된 1번 코멘트입니다.",
	//	"kind": "named",
	//}
	//
	//after, err := commentUseCase.Update("jinsu", before.ID, updateData)
	//assert.NoError(t, err)
	//assert.Equal(t, "수정된 1번 코멘트입니다.", after.Content)
	//assert.Equal(t, "named", after.Kind)
}

// Delete 작업 자체는 Repository 계층에서 이루어져야할 듯하고
// service 계층에서는 간단한 테스트만 수행한다.
func TestCommentUseCase_삭제된_댓글에_대한_조회(t *testing.T) {
	BeforeCommentUseCaseTest(t)
	defer A(t)
	test.Comment1JinsuAnnonymous.State = "deleted"
	mockCommentRepository.EXPECT().Get(gomock.Any()).Return(test.Comment1JinsuAnnonymous, nil)
	mockLikeCommentRepository.EXPECT().List(gomock.Any()).Return([]*model.LikeComment{}).AnyTimes()
	comment, err := commentUseCase.Get("Anything will be fine", 1)
	assert.NoError(t, err)
	assert.Equal(t, DeletedCommentUsername, comment.AuthorUsername)
	assert.Equal(t, DeletedCommentUsername, comment.Author.Username)
	assert.Equal(t, DeletedCommentNickname, comment.Author.Nickname)
	assert.Equal(t, DeletedCommentContent, comment.Content)
}

func TestLikeCommentUseCase_Toggle(t *testing.T) {
	BeforeCommentUseCaseTest(t)
	defer A(t)
	commentID := test.Comment1JinsuAnnonymous.ID
	mockCommentRepository.EXPECT().Get(gomock.Any()).Return(test.Comment1JinsuAnnonymous, nil).AnyTimes()
	likeComments := make([]*model.LikeComment, 0)
	mockLikeCommentRepository.EXPECT().Create(gomock.Any()).DoAndReturn(func(c *model.LikeComment) (*model.LikeComment, error) {
		likeComments = append(likeComments, c)
		return c, nil
	}).AnyTimes()
	// mock으로 그냥 한 칸 줄이기만함.
	mockLikeCommentRepository.EXPECT().Delete(gomock.Any()).DoAndReturn(func(id int) (*model.LikeComment, error) {
		deleted := likeComments[id]
		likeComments = append(likeComments[:id], likeComments[id+1:]...)
		return deleted, nil
	}).AnyTimes()
	// 거의 로직을 구현해버렸네.....
	mockLikeCommentRepository.EXPECT().List(gomock.Any()).DoAndReturn(func(option *repository.LikeCommentQueryOption) []*model.LikeComment {
		answers := make([]*model.LikeComment, 0)
		// 쿼리 조건 판별
		if option.Username != "" && option.CommentID != 0 {
			for _, like := range likeComments {
				if like.Username == option.Username && like.CommentID == option.CommentID {
					answers = append(answers, like)
				}
			}
		} else if option.Username != "" {
			for _, like := range likeComments {
				if like.Username == option.Username {
					answers = append(answers, like)
				}
			}
		} else if option.CommentID != 0 {
			for _, like := range likeComments {
				if like.CommentID == option.CommentID {
					answers = append(answers, like)
				}
			}
		} else {
			answers = likeComments
		}
		return answers
	}).AnyTimes()

	t.Run("Somebody toggle(create&delete) jinsu's comment", func(t *testing.T) {

		// toggle to create
		func() {
			created, err := likeCommentUseCase.Toggle(
				&model.LikeComment{
					CommentID: commentID,
					Username:  "somebody",
				})
			assert.Nil(t, err)
			assert.True(t, created)
		}()
		// toggle to delete
		func() {
			deleted, err := likeCommentUseCase.Toggle(
				&model.LikeComment{
					CommentID: commentID,
					Username:  "somebody",
				})
			assert.Nil(t, err)
			assert.False(t, deleted)
		}()

		// toggle to create again
		func() {
			created, err := likeCommentUseCase.Toggle(
				&model.LikeComment{
					CommentID: commentID,
					Username:  "somebody",
				})
			assert.Nil(t, err)
			assert.True(t, created)
		}()

		// 자기 댓글은 좋아요 불가능.
		func() {
			created, err := likeCommentUseCase.Toggle(
				&model.LikeComment{
					CommentID: test.Comment1JinsuAnnonymous.ID,
					Username:  test.UserJinsu.Username,
				})
			assert.NotNil(t, err)
			assert.False(t, created)
		}()
	})
}
