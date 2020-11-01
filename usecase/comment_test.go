// repository 계층과 분리된 채 usecase 계층만의 테스트를 진행
// repository를 간단히 mock 한 뒤 기명, 익명 kind의 comment 들이
// desired한 방향으로 얻어지는 지 확인해본다.
package usecase

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CommentRepositoryMock struct{}

var (
	commentsMock   []*model.Comment
	commentUseCase CommentUseCaseInterface
)

func (r *CommentRepositoryMock) Create(comment *model.Comment) error {
	commentsMock = append(commentsMock, comment)
	return nil
}

// QyeryOption기능은 제외하고 mock
func (r *CommentRepositoryMock) List(opt *repository.CommentQueryOption) []*model.Comment {
	return commentsMock
}

func (r *CommentRepositoryMock) Get(id int) *model.Comment {
	for _, comment := range commentsMock {
		if int(comment.ArticleID) == id {
			return comment
		}
	}
	return nil
}

func TestInit(t *testing.T) {
	// build container
	cont := dig.New()
	err := cont.Provide(repository.NewGorm)
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
}

func TestCommentUseCase_List(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	context := e.NewContext(req, nil)
	context.Set("user_id", "jinsu")
	resultComments := commentUseCase.List(context)
	log.Println(resultComments)
	t.Run("My anonymous comment", func(t *testing.T) {
		c := resultComments[0]
		assert.Equal(t, c.Kind, "anonymous")
		assert.Equal(t, "jinsu", c.Author.Username)
	})
	t.Run("My named comment", func(t *testing.T) {
		c := resultComments[1]
		assert.Equal(t, "named", c.Kind)
		assert.Equal(t, "jinsu", c.Author.Username)
	})
	t.Run("Others anonymous comment", func(t *testing.T) {
		c := resultComments[2]
		assert.Equal(t, "anonymous", c.Kind)
		assert.Equal(t, "익명", c.Author.Username)
		assert.Equal(t, "someone", c.AuthorUsername)
	})

}

func _mockup(t *testing.T){
	var id uint = 1
	myAnonymousComment := &model.Comment{
		Kind:           "anonymous",
		AuthorUsername: "jinsu",
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 익명 코멘트",
		ParentID:       nil,
	}
	id++
	myNamedComment := &model.Comment{
		Kind:           "named",
		AuthorUsername: "jinsu",
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 기명 코멘트",
		ParentID:       nil,
	}
	id++
	othersAnonymousComment := &model.Comment{
		Kind:           "anonymous",
		AuthorUsername: "somebody",
		ArticleID:      1,
		Content:        "테스트로 작성한 somebody의 익명 코멘트",
		ParentID:       nil,
	}
	commentsMock = append(commentsMock, myAnonymousComment)
	commentsMock = append(commentsMock, myNamedComment)
	commentsMock = append(commentsMock, othersAnonymousComment)

	assert.Len(t, commentsMock, 3)
}
