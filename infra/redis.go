package infra

import (
	"context"
	"fmt"

	rcache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/repository/cache"
)

type RedisAdapter interface {
	cache.CommentCacheRepository
	cache.LikeCommentCacheRepository
}

type RedisAdapterImpl struct {
	client *rcache.Cache `name:"CommentCacheRepository,LikeCommentCacheRepository"`
}

func NewRedisAdapter() RedisAdapter {
	ring := redis.NewRing(&redis.RingOptions{Addrs: map[string]string{
		"server_1": config.Config.Redis.Addr,
	}})
	client := rcache.New(&rcache.Options{Redis: ring})

	return &RedisAdapterImpl{
		client: client,
	}
}

func (a *RedisAdapterImpl) SetCommentsByArticleID(articleID int, coms data.CommentEntities) {
	var (
		key = fmt.Sprintf("khumu-comment:comments:article=%d", articleID)
	)

	log.Infof("캐시를 수정합니다. key=%s", key)
	err := a.client.Set(&rcache.Item{
		Ctx:   context.TODO(),
		Key:   key,
		Value: &coms,
	})
	if err != nil {
		log.Error(err)
	}
}

//func (a *RedisAdapterImpl) GetCommentsByArticle(articleID int) data.CommentEntities {
//	var (
//		key = fmt.Sprintf("khumu-comment:comments:article=%d", articleID)
//		val = data.CommentEntities(make([]*ent.Comment, 0))
//	)
//	log.Infof("Article에 대한 댓글 캐시를 조회합니다.key=%s", key)
//	err := a.cache.Get(context.TODO(), key, &val)
//	if err != nil {
//		// 캐시 미스
//		if errors.Is(err, cache.ErrCacheMiss) {
//			a.RefreshCommentsByArticle(articleID)
//			err = a.cache.Get(context.TODO(), key, &val)
//			if err != nil {
//				log.Errorf("캐시 미스 후 RefreshCommentsByArticle 했지만 에러 발생: %v", err)
//			}
//		} else {
//			log.Error(err)
//			return []*ent.Comment{}
//		}
//	}
//
//	return val
//}

func (a *RedisAdapterImpl) FindAllParentCommentsByArticleID(articleID int) (data.CommentEntities, error) {
	var (
		key = fmt.Sprintf("khumu-comment:comments:article=%d", articleID)
		val = data.CommentEntities(make([]*ent.Comment, 0))
	)
	log.Infof("Article에 대한 댓글 캐시를 조회합니다.key=%s", key)
	err := a.client.Get(context.TODO(), key, &val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

//func (a *RedisAdapterImpl) GetLikeCommentsByUserAndComment(username string, commentID int) data.LikeCommentEntities {
//	var (
//		key = fmt.Sprintf("khumu-comment:comments:user=%s&comment=%d", username, commentID)
//		val = data.LikeCommentEntities(make([]*ent.LikeComment, 0))
//	)
//	log.Infof("username=%s, comment=%d 대한 댓글 좋아요를 조회합니다.key=%s", username, commentID, key)
//	err := a.cache.Get(context.TODO(), key, &val)
//	if err != nil {
//		// 캐시 미스
//		if errors.Is(err, cache.ErrCacheMiss) {
//			a.RefreshLikeCommentsByUserAndComment(username, commentID)
//			err = a.cache.Get(context.TODO(), key, &val)
//			if err != nil {
//				log.Errorf("캐시 미스 후 RefreshCommentsByArticle 했지만 에러 발생: %v", err)
//			}
//		} else {
//			log.Error(err)
//			return []*ent.LikeComment{}
//		}
//	}
//
//	return val
//}

func (a *RedisAdapterImpl) SetLikesByCommentID(commentID int, likes data.LikeCommentEntities) {
	var (
		key = fmt.Sprintf("khumu-comment:likeComments:comment=%d", commentID)
	)

	log.Infof("캐시를 수정합니다. key=%s", key)
	err := a.client.Set(&rcache.Item{
		Ctx:   context.TODO(),
		Key:   key,
		Value: &likes,
	})
	if err != nil {
		log.Error(err)
	}
}

func (a *RedisAdapterImpl) FindAllByCommentID(commentID int) (data.LikeCommentEntities, error) {
	var (
		key = fmt.Sprintf("khumu-comment:likeComments:comment=%d", commentID)
		val = data.LikeCommentEntities(make([]*ent.LikeComment, 0))
	)
	log.Infof("comment=%d 대한 댓글 좋아요를 조회합니다.key=%s", commentID, key)
	err := a.client.Get(context.TODO(), key, &val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (a *RedisAdapterImpl) SetCommentCountByArticleID(articleID, cnt int) {
	key := fmt.Sprintf("khumu-comment:comment_counts:article=%d:", articleID)
	log.Infof("article=%d 의 comment count 개수를 저장합니다. cnt=%d", articleID, cnt)
	if err := a.client.Set(&rcache.Item{
		Ctx:   context.TODO(),
		Key:   key,
		Value: cnt,
	}); err != nil {
		log.Error(err)
	}
}

func (a *RedisAdapterImpl) Count(articleID int) (int, error) {
	key := fmt.Sprintf("khumu-comment:comment_counts:article=%d:", articleID)
	log.Infof("article=%d 의 comment count 개수를 조회합니다.", articleID)
	cnt := 0
	if err := a.client.Get(context.TODO(), key, &cnt); err != nil {
		return 0, errors.WithStack(err)
	}

	return cnt, nil
}
