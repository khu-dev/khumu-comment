// TODO: 현재는 거의 usecase level에서 repository 계층까지 테스트하는 셈인데
// 좀 더 순수한 service 계층의 logic을 테스트할 수 있도록 바뀌었으면 좋겠다.
package usecase

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentUseCase_Create(t *testing.T) {
	t.Run("jinsu의 익명 Article 댓글", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		tmp, err := commentUseCase.Create(test.UserJinsu.ID, &data.CommentInput{
			Author:  test.UserJinsu.ID,
			Article: &test.Articles[0].ID,
			Content: "테스트 댓글",
			Kind:    pointer.ToString("anonymous"),
		})
		assert.NoError(t, err)

		comment, err := commentUseCase.Get(test.UserJinsu.ID, tmp.ID)
		assert.NoError(t, err)
		// 기본적으로는 익명 댓글임.
		assert.Equal(t, AnonymousCommentUsername, comment.Author.Username)
		assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
		assert.Equal(t, "테스트 댓글", comment.Content)
	})

	t.Run("jinsu의 기명 Article 댓글", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		tmp, _ := commentUseCase.Create(test.UserJinsu.ID, &data.CommentInput{
			Author:  test.UserJinsu.ID,
			Article: &test.Articles[0].ID,
			Content: "테스트 기명 댓글",
			Kind:    pointer.ToString("named"),
		})

		comment, err := commentUseCase.Get(test.UserJinsu.ID, tmp.ID)
		assert.NoError(t, err)
		// 기본적으로는 익명 댓글임.
		assert.Equal(t, test.UserJinsu.ID, comment.Author.Username)
		assert.Equal(t, test.UserJinsu.Nickname, comment.Author.Nickname)
		assert.Equal(t, "테스트 기명 댓글", comment.Content)
	})

	t.Run("jinsu의 익명 Article 댓글에 대한 익명 대댓글", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		tmp, _ := commentUseCase.Create(test.UserJinsu.ID, &data.CommentInput{
			Author:  test.UserJinsu.ID,
			Article: &test.Articles[0].ID,
			Parent:  &test.Comment1JinsuAnonymous.ID,
			Content: "테스트 익명 댓글에 대한 익명 대댓글",
			Kind:    pointer.ToString("anonymous"),
		})

		comment, err := commentUseCase.Get(test.UserJinsu.ID, tmp.ID)
		assert.NoError(t, err)
		// 기본적으로는 익명 댓글임.
		assert.Equal(t, AnonymousCommentUsername, comment.Author.Username)
		assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
		assert.Equal(t, test.Comment1JinsuAnonymous.ID, *comment.Parent)
		assert.Equal(t, "테스트 익명 댓글에 대한 익명 대댓글", comment.Content)
	})

	t.Run("실패) 커뮤니티 article도 study article도 아닌 경우", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		newComment, err := commentUseCase.Create(test.UserJinsu.ID, &data.CommentInput{
			Author:  test.UserJinsu.ID,
			Content: "테스트 댓글",
		})
		assert.Error(t, err)
		assert.Nil(t, newComment)
	})
}

func TestCommentUseCase_List(t *testing.T) {
	t.Run("게시글에 대한 댓글 리스트", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		var err error
		tmpArticle, err := db.Article.Create().SetAuthor(test.UserPuppy).Save(context.TODO())
		assert.NoError(t, err)

		correctComment, err := db.Comment.Create().
			// 1번 article은 이미 많은 댓
			SetArticleID(tmpArticle.ID).
			SetAuthorID(test.UserPuppy.ID).
			SetContent("테스트 댓글").
			Save(context.TODO())
		assert.NoError(t, err)

		results, err := commentUseCase.List(test.UserPuppy.ID, &CommentQueryOption{ArticleID: tmpArticle.ID})
		assert.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, results[0].ID, correctComment.ID)
	})

	// 스터디 게시글은 개발 중지
	//t.Run("스터디 게시글에 대한 댓글 리스트", func(t *testing.T) {
	//	BeforeEach(t)
	//	defer AfterEach(t)
	//	var err error
	//	correctComment1, err := db.Comment.Create().
	//		SetStudyArticleID(1).
	//		SetAuthorID(test.UserPuppy.ID).
	//		SetContent("테스트 댓글").
	//		Save(context.TODO())
	//	assert.NoError(t, err)
	//	_, err = db.Comment.Create().
	//		SetStudyArticleID(2).
	//		SetAuthorID(test.UserPuppy.ID).
	//		SetContent("테스트 댓글").
	//		Save(context.TODO())
	//	assert.NoError(t, err)
	//
	//	results, err := commentUseCase.List(test.UserPuppy.ID, &CommentQueryOption{StudyArticleID: 1})
	//	assert.NoError(t, err)
	//	assert.Len(t, results, 1)
	//	assert.Equal(t, results[0].ID, correctComment1.ID)
	//})
}
func TestCommentUseCase_Get(t *testing.T) {
	t.Run("기명 댓글과 그 대댓글들", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		deletedAnonymousCommentFromComment1, err := db.Comment.Create().
			SetArticle(test.Comment1JinsuAnonymous.QueryArticle().OnlyX(context.TODO())).
			SetAuthorID(test.UserPuppy.ID).
			SetContent("테스트 댓글").
			SetKind("anonymous").
			SetState("deleted").
			Save(context.TODO())
		assert.NoError(t, err)

		deletedNamedCommentFromComment1, err := db.Comment.Create().
			SetArticle(test.Comment1JinsuAnonymous.QueryArticle().OnlyX(context.TODO())).
			SetAuthorID(test.UserPuppy.ID).
			SetContent("테스트 댓글").
			SetKind("named").
			SetState("deleted").
			Save(context.TODO())
		assert.NoError(t, err)

		comment, err := commentUseCase.Get(test.UserJinsu.ID, test.Comment1JinsuAnonymous.ID)
		assert.Greater(t, len(comment.Children), 0)
		for _, child := range comment.Children {
			switch child.ID {
			// 본인의 익명 대댓글
			case test.Comment5JinsuAnonymousFromComment1.ID:
				assert.True(t, child.IsAuthor)
				assert.Equal(t, AnonymousCommentUsername, child.Author.Username)
				assert.Equal(t, AnonymousCommentNickname, child.Author.Nickname)
			// 본인의 기명 대댓글
			case test.Comment6JinsuNamedFromComment1.ID:
				assert.True(t, child.IsAuthor)
				assert.Equal(t, test.UserJinsu.ID, child.Author.Username)
				assert.Equal(t, test.UserJinsu.Nickname, child.Author.Nickname)
			// 타인의 익명 대댓글
			case test.Comment7SomebodyAnonymousFromComment1.ID:
				assert.False(t, child.IsAuthor)
				assert.Equal(t, AnonymousCommentUsername, child.Author.Username)
				assert.Equal(t, AnonymousCommentNickname, child.Author.Nickname)
			// 타인의 삭제된 익명 대댓글
			case deletedAnonymousCommentFromComment1.ID:
				assert.False(t, child.IsAuthor)
				assert.Equal(t, DeletedCommentUsername, child.Author.Username)
				assert.Equal(t, DeletedCommentNickname, child.Author.Nickname)
				assert.Equal(t, child.Content, DeletedCommentContent)
			// 타인의 삭제된 기명 대댓글
			case deletedNamedCommentFromComment1.ID:
				assert.False(t, child.IsAuthor)
				assert.Equal(t, DeletedCommentUsername, child.Author.Username)
				assert.Equal(t, DeletedCommentNickname, child.Author.Nickname)
				assert.Equal(t, child.Content, DeletedCommentContent)
			}
		}
		assert.NoError(t, err)
		// 기본적으로는 익명 댓글임.
		assert.Equal(t, AnonymousCommentUsername, comment.Author.Username)
		assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
	})

	t.Run("익명 댓글", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		comment, err := commentUseCase.Get(test.UserJinsu.ID, test.Comment1JinsuAnonymous.ID)
		assert.NoError(t, err)
		// 기본적으로는 익명 댓글임.
		assert.Equal(t, AnonymousCommentUsername, comment.Author.Username)
		assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
	})

	t.Run("삭제된 댓글", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		tmp, err := db.Comment.Create().
			SetArticleID(1).
			SetAuthor(test.UserJinsu).
			SetContent("삭제되어 보여져라!").
			SetState("deleted").
			Save(context.TODO())
		assert.NoError(t, err)
		comment, err := commentUseCase.Get(test.UserJinsu.ID, tmp.ID)
		assert.NoError(t, err)

		// 기본적으로는 익명 댓글임.
		assert.Equal(t, DeletedCommentUsername, comment.Author.Username)
		assert.Equal(t, DeletedCommentNickname, comment.Author.Nickname)
		assert.Equal(t, DeletedCommentContent, comment.Content)
	})
}

func TestLikeCommentUseCase_List(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	deletedAnonymousCommentFromComment1, err := db.Comment.Create().
		SetArticle(test.Comment1JinsuAnonymous.QueryArticle().OnlyX(context.TODO())).
		SetAuthorID(test.UserPuppy.ID).
		SetContent("테스트 댓글").
		SetKind("anonymous").
		SetState("deleted").
		Save(context.TODO())
	assert.NoError(t, err)

	deletedNamedCommentFromComment1, err := db.Comment.Create().
		SetArticle(test.Comment1JinsuAnonymous.QueryArticle().OnlyX(context.TODO())).
		SetAuthorID(test.UserPuppy.ID).
		SetContent("테스트 댓글").
		SetKind("named").
		SetState("deleted").
		Save(context.TODO())
	assert.NoError(t, err)

	comments, err := commentUseCase.List(test.UserJinsu.ID, &CommentQueryOption{})
	assert.NoError(t, err)
	for _, comment := range comments {
		if comment.State == "deleted" {
			assert.Equal(t, DeletedCommentNickname, comment.Author.Username)
			assert.Equal(t, DeletedCommentContent, comment.Content)
		} else {
			if comment.Kind == "named" {
				assert.NotEqual(t, AnonymousCommentUsername, comment.Author.Username)
				assert.NotEqual(t, AnonymousCommentNickname, comment.Author.Nickname)
				assert.NotEqual(t, DeletedCommentNickname, comment.Author.Username)
				assert.NotEqual(t, DeletedCommentUsername, comment.Author.Nickname)
			} else if comment.Kind == "anonymous" {
				assert.Equal(t, AnonymousCommentNickname, comment.Author.Nickname)
			}
		}

		// jinsu의 댓글들
		if comment.ID == test.Comment1JinsuAnonymous.ID ||
			comment.ID == test.Comment2JinsuNamed.ID ||
			comment.ID == test.Comment6JinsuNamedFromComment1.ID ||
			comment.ID == test.Comment5JinsuAnonymousFromComment1.ID {
			assert.True(t, comment.IsAuthor)
		}

		// 대표적으로 Comment1의 Children들 테스트
		if comment.ID == test.Comment1JinsuAnonymous.ID {
			for _, child := range comment.Children {
				switch child.ID {
				// 본인의 익명 대댓글
				case test.Comment5JinsuAnonymousFromComment1.ID:
					assert.True(t, child.IsAuthor)
					assert.Equal(t, AnonymousCommentUsername, child.Author.Username)
					assert.Equal(t, AnonymousCommentNickname, child.Author.Nickname)
				// 본인의 기명 대댓글
				case test.Comment6JinsuNamedFromComment1.ID:
					assert.True(t, child.IsAuthor)
					assert.Equal(t, test.UserJinsu.ID, child.Author.Username)
					assert.Equal(t, test.UserJinsu.Nickname, child.Author.Nickname)
				// 타인의 익명 대댓글
				case test.Comment7SomebodyAnonymousFromComment1.ID:
					assert.False(t, child.IsAuthor)
					assert.Equal(t, AnonymousCommentUsername, child.Author.Username)
					assert.Equal(t, AnonymousCommentNickname, child.Author.Nickname)
				// 타인의 삭제된 익명 대댓글
				case deletedAnonymousCommentFromComment1.ID:
					assert.False(t, child.IsAuthor)
					assert.Equal(t, DeletedCommentUsername, child.Author.Username)
					assert.Equal(t, DeletedCommentNickname, child.Author.Nickname)
					assert.Equal(t, child.Content, DeletedCommentContent)
				// 타인의 삭제된 기명 대댓글
				case deletedNamedCommentFromComment1.ID:
					assert.False(t, child.IsAuthor)
					assert.Equal(t, DeletedCommentUsername, child.Author.Username)
					assert.Equal(t, DeletedCommentNickname, child.Author.Nickname)
					assert.Equal(t, child.Content, DeletedCommentContent)
				}
			}
		}
	}
}

func TestCommentUseCase_Update(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	// Update는 대부분 repository 계층에서만 확인해도 될 듯.

	before := *test.Comment2JinsuNamed
	// named => anonymous로 변경
	// content 내용 변경
	updateData := map[string]interface{}{
		"content": "수정된 1번 코멘트입니다.",
		"kind":    "anonymous",
	}

	after, err := commentUseCase.Update("jinsu", before.ID, updateData)
	assert.NoError(t, err)
	assert.Equal(t, "수정된 1번 코멘트입니다.", after.Content)
	assert.Equal(t, "anonymous", after.Kind)
}

func TestCommentUseCase_Delete(t *testing.T) {
	t.Run("대댓글이 없는 부모 댓글 삭제", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		parentComment := test.Comment4PuppyAnonymous
		// 좋아요가 있으면 삭제해야함. 그 기능 테스트를 위해 좋아요도 만들어봄.
		for i := 0; i < 3; i++ {
			_, err := db.LikeComment.Create().
				SetAbout(parentComment).
				SetLikedBy(test.UserPuppy).
				Save(context.TODO())
			assert.NoError(t, err)
		}
		assert.NotNil(t, parentComment)
		err := commentUseCase.Delete(test.UserPuppy.ID, parentComment.ID)
		assert.NoError(t, err)
		_, err = db.Comment.Get(context.TODO(), parentComment.ID)
		assert.IsType(t, &ent.NotFoundError{}, err)
	})
	t.Run("대댓글 삭제", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		toDelete := test.Comment5JinsuAnonymousFromComment1
		err := commentUseCase.Delete(test.UserJinsu.ID, toDelete.ID)
		assert.NoError(t, err)
		_, err = db.Comment.Get(context.TODO(), toDelete.ID)
		assert.IsType(t, &ent.NotFoundError{}, err)
	})
	t.Run("대댓글이 있는 부모 댓글 삭제", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)
		toDelete := test.Comment1JinsuAnonymous
		err := commentUseCase.Delete(test.UserJinsu.ID, toDelete.ID)
		assert.NoError(t, err)
		deleted, err := db.Comment.Query().WithAuthor().Where(comment.ID(toDelete.ID)).Only(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, "deleted", deleted.State)
	})

}

func TestLikeCommentUseCase_Toggle(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	commentID := test.Comment1JinsuAnonymous.ID
	// mock으로 그냥 한 칸 줄이기만함.
	t.Run("Somebody toggle(create&delete) jinsu's comment", func(t *testing.T) {
		// toggle to create
		func() {
			created, err := likeCommentUseCase.Toggle(&data.LikeCommentInput{
				Comment: commentID,
				User:    "somebody",
			})
			assert.Nil(t, err)
			assert.True(t, created)
		}()
		// toggle to delete
		func() {
			created, err := likeCommentUseCase.Toggle(&data.LikeCommentInput{
				Comment: commentID,
				User:    "somebody",
			})
			assert.Nil(t, err)
			assert.False(t, created)
		}()

		// toggle to create again
		func() {
			created, err := likeCommentUseCase.Toggle(&data.LikeCommentInput{
				Comment: commentID,
				User:    "somebody",
			})
			assert.Nil(t, err)
			assert.True(t, created)
		}()

		// 자기 댓글은 좋아요 불가능.
		func() {
			created, err := likeCommentUseCase.Toggle(&data.LikeCommentInput{
				Comment: commentID,
				User:    test.UserJinsu.ID,
			})
			assert.NotNil(t, err)
			assert.False(t, created)
		}()
	})
}
