package external

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/repository"
	log "github.com/sirupsen/logrus"
)

type RedisAdapter interface {
	//InvalidateComment(commentID int)
	InvalidateCommentsOfArticle(articleID int)
	Refresh(articleID int)
	GetAllByArticle(articleID int) data.CommentEntities
}

type RedisAdapterImpl struct {
	cache             *cache.Cache
	commentRepository repository.CommentRepository
}

func NewRedisAdapter(commentRepository repository.CommentRepository) RedisAdapter {
	ring := redis.NewRing(&redis.RingOptions{Addrs: map[string]string{
		"server_1": config.Config.Redis.Addr,
	}})
	cacheCli := cache.New(&cache.Options{Redis: ring})

	return &RedisAdapterImpl{
		cache:             cacheCli,
		commentRepository: commentRepository,
	}
}

// 댓글 생성 시 게시글의 캐시쪽에 댓글 내용을 invalidate함.
func (a *RedisAdapterImpl) InvalidateCommentsOfArticle(articleID int) {
	//log.Infof("Article(id=%d)에 대한 댓글들의 캐시를 invalidate합니다.", articleID)
	//stringCmd := a.client.SMembers(context.Background(), fmt.Sprintf("conj:comment_comment:article_id=%d", articleID))
	//cacheKeys, err := stringCmd.Result()
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//for _, cacheKey := range cacheKeys {
	//	// cacheKey는 q:abcd123abc123abc.. 의 형태이다.
	//	log.Infof("redis에서 %s를 Delete합니다.", cacheKey)
	//	intCmd := a.client.Del(context.Background(), cacheKey)
	//	err = intCmd.Err()
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}
}

func (a *RedisAdapterImpl) Refresh(articleID int) {
	var (
		key = fmt.Sprintf("khumu-comment:comments:article=%d", articleID)
		val = data.CommentEntities(make([]*ent.Comment, 0))
	)
	val, err := a.commentRepository.FindAllParentsByArticleID(articleID)
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("캐시를 refresh합니다. key=%s", key)
	err = a.cache.Set(&cache.Item{
		Ctx:   context.TODO(),
		Key:   key,
		Value: &val,
	})
	if err != nil {
		log.Error(err)
	}
}

func (a *RedisAdapterImpl) GetAllByArticle(articleID int) data.CommentEntities {
	var (
		key = fmt.Sprintf("khumu-comment:comments:article=%d", articleID)
		val = data.CommentEntities(make([]*ent.Comment, 0))
	)
	log.Infof("Article에 대한 댓글 캐시를 조회합니다.key=%s", key)
	err := a.cache.Get(context.TODO(), key, &val)
	if err != nil {
		// 캐시 미스
		if errors.Is(err, cache.ErrCacheMiss) {
			a.Refresh(articleID)
			err = a.cache.Get(context.TODO(), key, &val)
			if err != nil {
				log.Errorf("캐시 미스 후 Refresh 했지만 에러 발생: %v", err)
			}
		} else {
			log.Error(err)
			return []*ent.Comment{}
		}
	}

	return val
}
