package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
	log "github.com/sirupsen/logrus"
)

type CommentRepository interface {
	Create(createInput *data.CommentInput) (*ent.Comment, error)
	FindAllParentsByAuthorID(authorID string) ([]*ent.Comment, error)
	FindAllParentsByArticleID(articleID int) ([]*ent.Comment, error)
	FindAllParentsByStudyArticleID(articleID int) ([]*ent.Comment, error)
	Get(id int) (*ent.Comment, error)
	Update(id int, updateInput map[string]interface{}) (*ent.Comment, error)
	Delete(id int) error
}

type commentRepository struct {
	db *ent.Client
}

func NewCommentRepository(client *ent.Client) CommentRepository {
	return &commentRepository{
		db: client,
	}
}

func (c commentRepository) Create(createInput *data.CommentInput) (newComment *ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
	}()
	newComment, err = c.db.Comment.Create().
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

	newComment, err = c.db.Comment.Query().
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

	return newComment, nil
}

func (c commentRepository) FindAllParentsByAuthorID(authorID string) ([]*ent.Comment, error) {
	query := c.db.Comment.Query().Where(comment.HasAuthorWith(khumuuser.ID(authorID)))
	parents, err := AppendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())
	return parents, err
}

func (c commentRepository) FindAllParentsByArticleID(articleID int) ([]*ent.Comment, error) {
	query := c.db.Comment.Query().Where(comment.HasArticleWith(article.ID(articleID)))
	parents, err := AppendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())
	return parents, err
}

func (c commentRepository) FindAllParentsByStudyArticleID(articleID int) ([]*ent.Comment, error) {
	query := c.db.Comment.Query().Where(comment.HasStudyArticleWith(studyarticle.ID(articleID)))
	parents, err := AppendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())
	return parents, err
}

func (c commentRepository) Get(id int) (com *ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	query := c.db.Comment.Query().Where(comment.ID(id))
	com, err = AppendQueryForComment(query).
		Only(context.TODO())
	return com, err
}

func (c commentRepository) Update(id int, updateInput map[string]interface{}) (com *ent.Comment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	ctx := context.TODO()
	query := c.db.Comment.Update().Where(comment.ID(id))
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

	return c.Get(id)
}

func (c commentRepository) Delete(id int) (err error) {
	ctx := context.TODO()
	log.Info("부모 댓글이 없어 댓글 자체를 삭제하는 작업을 시작합니다.")
	tx, err := c.db.BeginTx(ctx, new(sql.TxOptions))
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

	err = tx.Comment.DeleteOneID(id).Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Comment(id=%d)를 삭제했습니다.", id)

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
