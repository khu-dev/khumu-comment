package repository

import (
	"context"
	"errors"
	rcache "github.com/go-redis/cache/v8"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/repository/cache"
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
	synchronousCacheWrite bool `optional:"true"` // dig에서 optional하게 주입받도록 설정 => zero value로 주입받을 수 있음
}

func NewLikeCommentRepository(
	client *ent.Client,
	cache cache.LikeCommentCacheRepository,
	synchronousCacheWrite bool) LikeCommentRepository {
	return &likeCommentRepository{
		db:                    client,
		cache:                 cache,
		synchronousCacheWrite: synchronousCacheWrite,
	}
}
func (l likeCommentRepository) Create(createInput *data.LikeCommentInput) (like *ent.LikeComment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	like, err = l.db.LikeComment.Create().SetLikedByID(createInput.User).SetAboutID(createInput.Comment).Save(context.TODO())

	l.setLikesCacheByCommentID(createInput.Comment)

	return
}

func (l likeCommentRepository) FindAllByCommentID(commentID int) (likes []*ent.LikeComment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	cached, err := l.cache.FindAllByCommentID(commentID)
	if err != nil {
		if !errors.Is(err, rcache.ErrCacheMiss) {
			log.Error(err)
		}

		likes, err := l.findAllByCommentIDWithoutCache(commentID)
		if err != nil {
			return nil, err
		}

		// 캐시 미스 발생 시 캐시를 기록
		l.cache.SetLikesByCommentID(commentID, likes)

		return likes, nil
	}

	return cached, nil
}

func (l likeCommentRepository) findAllByCommentIDWithoutCache(commentID int) (likes []*ent.LikeComment, err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	likes, err = l.db.LikeComment.Query().
		Where(likecomment.HasAboutWith(comment.ID(commentID))).
		WithLikedBy().
		All(context.TODO())
	if err != nil {
		return nil, err
	}

	return likes, nil
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

func (l likeCommentRepository) Delete(id int) (err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	like, err := l.db.LikeComment.Query().WithAbout().Where(likecomment.ID(id)).First(context.TODO())
	err = l.db.LikeComment.DeleteOneID(id).Exec(context.TODO())

	l.setLikesCacheByCommentID(like.Edges.About.ID)

	return
}

func (l likeCommentRepository) DeleteAllByCommentID(commentID int) (err error) {
	defer func() {
		err = WrapEntError(err)
		//err = nil
	}()
	n, err := l.db.LikeComment.Delete().Where(likecomment.HasAboutWith(comment.ID(commentID))).Exec(context.TODO())
	log.Infof("Comment(id=%d)에 대한 좋아요를 %d개 삭제했습니다.", commentID, n)

	l.setLikesCacheByCommentID(commentID)

	return
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
			log.Error(err)
		} else {
			l.cache.SetLikesByCommentID(commentID, likes)
		}
	}()
	// synchronous write이 false이면 buffered chan이라 바로 값을 넣을 수 있다.
	done <- struct{}{}

	return
}
