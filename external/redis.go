package external

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/khu-dev/khumu-comment/config"
	log "github.com/sirupsen/logrus"
)

type RedisAdapter interface {
	//InvalidateComment(commentID int)
	InvalidateCommentsOfArticle(articleID int)
}

type RedisAdapterImpl struct {
	client *redis.Client
}

func NewRedisAdapter() RedisAdapter {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisAdapterImpl{
		client: client,
	}
}

// 댓글 생성 시 게시글의 캐시쪽에 댓글 내용을 invalidate함.
func (a *RedisAdapterImpl) InvalidateCommentsOfArticle(articleID int) {
	log.Infof("Article(id=%d)에 대한 댓글들의 캐시를 invalidate합니다.", articleID)
	stringCmd := a.client.SMembers(context.Background(), fmt.Sprintf("conj:comment_comment:article_id=%d", articleID))
	cacheKeys, err := stringCmd.Result()
	if err != nil {
		log.Error(err)
		return
	}
	for _, cacheKey := range cacheKeys {
		// cacheKey는 q:abcd123abc123abc.. 의 형태이다.
		log.Infof("redis에서 %s를 Delete합니다.", cacheKey)
		intCmd := a.client.Del(context.Background(), cacheKey)
		err = intCmd.Err()
		if err != nil {
			log.Error(err)
		}
	}
}
