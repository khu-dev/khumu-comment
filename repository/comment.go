package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	rcache "github.com/go-redis/cache/v8"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
	"github.com/khu-dev/khumu-comment/repository/cache"
	log "github.com/sirupsen/logrus"
)

type CommentRepository interface {
	Create(createInput *data.CommentInput) (com *ent.Comment, err error)
	FindAllParentCommentsByAuthorID(authorID string) (coms []*ent.Comment, err error)
	FindAllParentCommentsByArticleID(articleID int) (coms []*ent.Comment, err error)
	//FindAllParentCommentsByStudyArticleID(articleID int) (coms []*ent.Comment, err error)
	Get(id int) (com *ent.Comment, err error)
	Update(id int, updateInput map[string]interface{}) (com *ent.Comment, err error)
	Delete(id int) (err error)
}

type commentRepository struct {
	db    *ent.Client
	cache cache.CommentCacheRepository
	// synchronousCacheWrite 은 cache를 concurrent하게 write할 것인지 synchrnous하게 write할 것인지를 의미
	synchronousCacheWrite SynchronousCacheWrite
}

func NewCommentRepository(client *ent.Client, cache cache.CommentCacheRepository, synchronousCacheWrite SynchronousCacheWrite) CommentRepository {
	return &commentRepository{
		db:                    client,
		cache:                 cache,
		synchronousCacheWrite: synchronousCacheWrite,
	}
}

func (repo commentRepository) Create(createInput *data.CommentInput) (newComment *ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
	}()
	newComment, err = repo.db.Comment.Create().
		SetNillableArticleID(createInput.Article).
		SetNillableStudyArticleID(createInput.StudyArticle).
		SetNillableParentID(createInput.Parent).
		SetAuthorID(createInput.Author).
		SetContent(createInput.Content).
		SetState("exists").
		SetNillableKind(createInput.Kind).
		Save(context.TODO())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	newComment, err = repo.db.Comment.Query().
		// 댓글 작성자
		WithAuthor().
		WithArticle(func(query *ent.ArticleQuery) {
			query.WithAuthor()
		}).
		WithStudyArticle(func(query *ent.StudyArticleQuery) {
			query.WithAuthor()
		}).
		// 대댓글. 근데 어차피 새 댓글에는 대댓글이 없긴 하다.
		WithChildren(
			func(query *ent.CommentQuery) {
				query.
					// 대댓글의 작성자
					WithAuthor().
					// 대댓글의 게시글
					WithArticle(func(query *ent.ArticleQuery) {
						query.WithAuthor()
					}).
					// 대댓글의 게시글
					WithStudyArticle(func(query *ent.StudyArticleQuery) {
						query.WithAuthor()
					})
			},
		).
		Where(comment.ID(newComment.ID)).
		Only(context.TODO())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	repo.setCommentsCacheByArticleID(*createInput.Article)

	return newComment, nil
}

func (repo commentRepository) FindAllParentCommentsByAuthorID(authorID string) (coms []*ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	query := repo.db.Comment.Query().Where(comment.HasAuthorWith(khumuuser.ID(authorID)))
	parents, err := AppendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())
	return parents, err
}

func (repo commentRepository) FindAllParentCommentsByArticleID(articleID int) ([]*ent.Comment, error) {
	cached, err := repo.cache.FindAllParentCommentsByArticleID(articleID)
	if err != nil {
		if !errors.Is(err, rcache.ErrCacheMiss) {
			log.Error(err)
		}
		coms, err := repo.findParentCommentsByArticleWithoutCache(articleID)
		if err != nil {
			return nil, err
		}
		// 캐시 미스 발생 시 캐시를 기록
		go repo.cache.SetCommentsByArticleID(articleID, coms)
		return coms, err
	}

	return cached, nil
}

func (repo commentRepository) FindAllParentCommentsByStudyArticleID(articleID int) (coms []*ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	query := repo.db.Comment.Query().Where(comment.HasStudyArticleWith(studyarticle.ID(articleID)))
	parents, err := AppendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())
	return parents, err
}

func (repo commentRepository) Get(id int) (com *ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	query := repo.db.Comment.Query().Where(comment.ID(id))
	com, err = AppendQueryForComment(query).
		Only(context.TODO())
	return com, err
}

func (repo commentRepository) Update(id int, updateInput map[string]interface{}) (com *ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	ctx := context.TODO()
	query := repo.db.Comment.Update().Where(comment.ID(id))
	if val, ok := updateInput["state"]; ok {
		log.Infof("Comment(id=%d)의 state를 %s로 변경합니다.", id, val)
		query = query.SetState(val.(string))
	}
	if val, ok := updateInput["content"]; ok {
		log.Infof("Comment(id=%d)의 content를 %s로 변경합니다.", id, val)
		query = query.SetContent(val.(string))
	}
	if val, ok := updateInput["kind"]; ok {
		log.Infof("Comment(id=%d)의 state를 %s로 변경합니다.", id, val)
		query = query.SetKind(val.(string))
	}

	_, err = query.Save(ctx)
	if err != nil {
		return nil, err
	}

	updated, err := repo.Get(id)
	if err != nil {
		return nil, err
	}

	go repo.setCommentsCacheByArticleID(updated.Edges.Article.ID)

	return updated, nil
}

func (repo commentRepository) Delete(id int) (err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	ctx := context.TODO()
	log.Info("부모 댓글이 없어 댓글 자체를 삭제하는 작업을 시작합니다.")
	tx, err := repo.db.BeginTx(ctx, new(sql.TxOptions))
	defer func() {
		if err = tx.Commit(); err != nil {
			log.Error(err)
		}
	}()
	if err != nil {
		log.Error(err)
		return err
	}

	n, err := tx.LikeComment.Delete().
		Where(likecomment.HasAboutWith(comment.ID(id))).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Comment(id=%d)에 대한 좋아요를 %d개 삭제했습니다.", id, n)

	com, err := repo.Get(id)
	if err != nil {
		return err
	}

	err = tx.Comment.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	log.Infof("Comment(id=%d)를 삭제했습니다.", id)

	repo.setCommentsCacheByArticleID(com.Edges.Article.ID)

	return nil
}

// 댓글 조회 시 기본적으로 필요한 query문을 추가한다.
// 댓글 작성자
// 게시물과 게시물의 작성자
// 스터디 게시물과 스터디 게시물의 작성자
// 대댓글도 마찬가지 과정
func AppendQueryForComment(query *ent.CommentQuery) *ent.CommentQuery {
	query.
		// 댓글 작성자
		WithAuthor().
		WithArticle(func(query *ent.ArticleQuery) {
			query.WithAuthor()
		}).
		WithStudyArticle(func(query *ent.StudyArticleQuery) {
			query.WithAuthor()
		}).
		// 대댓글
		WithChildren(
			func(query *ent.CommentQuery) {
				query.
					WithAuthor().
					WithArticle(func(query *ent.ArticleQuery) {
						query.WithAuthor()
					}).
					WithStudyArticle(func(query *ent.StudyArticleQuery) {
						query.WithAuthor()
					})
			},
		).
		// 부모 댓글
		WithParent()

	return query
}

func (repo *commentRepository) findParentCommentsByArticleWithoutCache(articleID int) (coms []*ent.Comment, err error) {
	query := repo.db.Comment.Query().Where(comment.HasArticleWith(article.ID(articleID)))
	parents, err := AppendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())
	return parents, err
}

// invalidate 는 부모 댓글에 대한 캐시를 invalidate 합니다.
func (repo *commentRepository) setCommentsCacheByArticleID(articleID int) {
	var done chan struct{}
	if repo.synchronousCacheWrite {
		done = make(chan struct{})
	} else {
		done = make(chan struct{}, 1)
	}
	go func() {
		defer func() {
			<-done
		}()

		coms, queryErr := repo.findParentCommentsByArticleWithoutCache(articleID)
		if queryErr != nil {
			log.Error(queryErr)
		} else {
			repo.cache.SetCommentsByArticleID(articleID, coms)
		}
	}()
	// synchronous write이 false이면 buffered chan이라 바로 값을 넣을 수 있다.
	done <- struct{}{}

	return
}
