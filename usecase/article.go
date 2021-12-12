package usecase

import (
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/repository"
	"math"
)

type ArticleUseCase interface {
	List(username string, body *data.GetCommentedArticlesReq) (*data.GetCommentedArticlesResp, error)
}

type articleUseCase struct {
	repo repository.ArticleRepository
}

func NewArticleUseCase(repo repository.ArticleRepository) ArticleUseCase {
	return &articleUseCase{repo: repo}
}

func (uc *articleUseCase) List(username string, body *data.GetCommentedArticlesReq) (*data.GetCommentedArticlesResp, error) {
	if body.Cursor == 0 {
		body.Cursor = math.MaxInt
	}
	if body.Size == 0 {
		body.Size = 20
	}

	articleIDs, err := uc.repo.FindAllIDByAuthorIDAndRecentlyCommented(username, body.Cursor, body.Size)
	if err != nil {
		return nil, err
	}

	return &data.GetCommentedArticlesResp{Data: articleIDs, Message: ""}, nil
}
