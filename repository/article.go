package repository

import (
	"context"
	gosql "database/sql"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/pkg/errors"
)

type ArticleRepository interface {
	FindAllIDByAuthorIDAndRecentlyCommented(authorID string, cursor, size int) ([]int, error)
}

type articleRepository struct {
	db    *gosql.DB
	entDB *ent.Client
}

func NewArticleRepository(db *gosql.DB, entDB *ent.Client) ArticleRepository {
	return &articleRepository{
		db:    db,
		entDB: entDB,
	}
}

func (repo *articleRepository) FindAllIDByAuthorIDAndRecentlyCommented(authorID string, cursor, size int) ([]int, error) {
	res, err := repo.db.QueryContext(context.Background(),
		"select a.id from articles a inner join comments c on a.id = c.article_comments where c.author_id=? and a.id < ? group by a.id, a.created_at order by max(c.created_at) desc limit ?",
		authorID, cursor, size)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	articleIDs := make([]int, 0)
	for res.Next() {
		var articleID int
		if err := res.Scan(&articleID); err != nil {
			return nil, errors.WithStack(err)
		}
		articleIDs = append(articleIDs, articleID)
	}

	return articleIDs, nil

	return articleIDs, nil
}
