package repository

import (
	"context"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	log "github.com/sirupsen/logrus"
)

type LikeCommentRepository interface {
	Create(createInput *data.LikeCommentInput) (like *ent.LikeComment, err error)
	FindAllByUserIDAndCommentID(authorID string, commentID int) (likes []*ent.LikeComment, err error)
	CountByCommentID(commentID int) (int, error)
	Delete(id int) error
	DeleteAllByCommentID(commentID int) error
}

type likeCommentRepository struct {
	db *ent.Client
}

func NewLikeCommentRepository(client *ent.Client) LikeCommentRepository {
	return &likeCommentRepository{
		db: client,
	}
}
func (l likeCommentRepository) Create(createInput *data.LikeCommentInput) (like *ent.LikeComment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	like, err = l.db.LikeComment.Create().SetLikedByID(createInput.User).SetAboutID(createInput.Comment).Save(context.TODO())
	return
}

func (l likeCommentRepository) FindAllByUserIDAndCommentID(userID string, commentID int) (likes []*ent.LikeComment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	likes, err = l.db.LikeComment.Query().
		Where(likecomment.HasLikedByWith(khumuuser.ID(userID)), likecomment.HasAboutWith(comment.ID(commentID))).
		All(context.TODO())
	return
}

func (l likeCommentRepository) CountByCommentID(commentID int) (n int, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	likes, err := l.db.LikeComment.Query().Select("id").WithAbout(func(query *ent.CommentQuery) {
		query.Where(comment.ID(commentID))
	}).All(context.TODO())
	if err != nil {
		return 0, err
	}

	return len(likes), nil
}

func (l likeCommentRepository) Delete(id int) (err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	err = l.db.LikeComment.DeleteOneID(id).Exec(context.TODO())
	return
}

func (l likeCommentRepository) DeleteAllByCommentID(commentID int) (err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	n, err := l.db.LikeComment.Delete().Where(likecomment.HasAboutWith(comment.ID(commentID))).Exec(context.TODO())
	log.Infof("Comment(id=%d)에 대한 좋아요를 %d개 삭제했습니다.", commentID, n)

	return
}
