package repository

import (
	"context"
	rcache "github.com/go-redis/cache/v8"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/repository/cache"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type LikeCommentRepository interface {
	Create(createInput *data.LikeCommentInput) (like *ent.LikeComment, err error)
	FindAllByCommentID(commentID int) (likes []*ent.LikeComment, err error)
	FindAllByUserIDAndCommentID(authorID string, commentID int) (likes []*ent.LikeComment, err error)
	Delete(id int) error
	DeleteAllByCommentID(commentID int) error
}

type likeCommentRepository struct {
	db    *ent.Client
	cache cache.LikeCommentCacheRepository `name:"LikeCommentCacheRepository"`
	// synchronousCacheWrite 은 cache를 concurrent하게 write할 것인지 synchrnous하게 write할 것인지를 의미
	synchronousCacheWrite SynchronousCacheWrite
}

func NewLikeCommentRepository(
	client *ent.Client,
	cache cache.LikeCommentCacheRepository,
	synchronousCacheWrite SynchronousCacheWrite) LikeCommentRepository {
	return &likeCommentRepository{
		db:                    client,
		cache:                 cache,
		synchronousCacheWrite: synchronousCacheWrite,
	}
}
func (l likeCommentRepository) Create(createInput *data.LikeCommentInput) (like *ent.LikeComment, err error) {
	like, err = l.db.LikeComment.Create().SetLikedByID(createInput.User).SetAboutID(createInput.Comment).Save(context.TODO())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	l.setLikesCacheByCommentID(createInput.Comment)

	return
}

func (l likeCommentRepository) FindAllByCommentID(commentID int) (likes []*ent.LikeComment, err error) {
	cached, err := l.cache.FindAllByCommentID(commentID)
	if err != nil {
		if !errors.Is(err, rcache.ErrCacheMiss) {
			log.Error(err)
		}

		likes, err := l.findAllByCommentIDWithoutCache(commentID)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// 캐시 미스 발생 시 캐시를 기록
		l.cache.SetLikesByCommentID(commentID, likes)

		return likes, nil
	}

	return cached, nil
}

func (l likeCommentRepository) findAllByCommentIDWithoutCache(commentID int) (likes []*ent.LikeComment, err error) {
	likes, err = l.db.LikeComment.Query().
		Where(likecomment.HasAboutWith(comment.ID(commentID))).
		WithLikedBy().
		All(context.TODO())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return likes, nil
}

func (l likeCommentRepository) FindAllByUserIDAndCommentID(userID string, commentID int) (likes []*ent.LikeComment, err error) {
	likes, err = l.db.LikeComment.Query().
		Where(likecomment.HasLikedByWith(khumuuser.ID(userID)), likecomment.HasAboutWith(comment.ID(commentID))).
		All(context.TODO())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return likes, nil
}

func (l likeCommentRepository) Delete(id int) (err error) {
	like, err := l.db.LikeComment.Query().WithAbout().Where(likecomment.ID(id)).First(context.TODO())
	err = l.db.LikeComment.DeleteOneID(id).Exec(context.TODO())
	if err != nil {
		return errors.WithStack(err)
	}
	l.setLikesCacheByCommentID(like.Edges.About.ID)

	return nil
}

func (l likeCommentRepository) DeleteAllByCommentID(commentID int) (err error) {
	n, err := l.db.LikeComment.Delete().Where(likecomment.HasAboutWith(comment.ID(commentID))).Exec(context.TODO())
	if err != nil {
		return errors.WithStack(err)
	}

	log.Infof("Comment(id=%d)에 대한 좋아요를 %d개 삭제했습니다.", commentID, n)

	l.setLikesCacheByCommentID(commentID)

	return nil
}

// invalidate 는 부모 댓글에 대한 캐시를 invalidate 합니다.
func (l *likeCommentRepository) setLikesCacheByCommentID(commentID int) {
	var done chan struct{}
	if l.synchronousCacheWrite {
		done = make(chan struct{})
	} else {
		done = make(chan struct{}, 1)
	}
	go func() {
		defer func() {
			<-done
		}()

		likes, err := l.findAllByCommentIDWithoutCache(commentID)
		if err != nil {
			log.Errorf("%+v", err)
		} else {
			l.cache.SetLikesByCommentID(commentID, likes)
		}
	}()
	// synchronous write이 false이면 buffered chan이라 바로 값을 넣을 수 있다.
	done <- struct{}{}

	return
}
