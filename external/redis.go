package external

import (
	"context"
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
	repository.CommentCacheRepository
	repository.LikeCommentCacheRepository
	//RefreshCommentsByArticle(articleID int)
	//GetCommentsByArticle(articleID int) (data.CommentEntities, error)
	//RefreshLikeCommentsByUserAndComment(username string, commentID int)
	//GetLikeCommentsByUserAndComment(username string, commentID int) data.LikeCommentEntities
	//RefreshLikeCommentsByComment(commentID int)
	//GetLikeCommentsByComment(commentID int) (data.LikeCommentEntities, error)
}

type RedisAdapterImpl struct {
	client                *cache.Cache
	commentRepository     repository.CommentRepository
	likeCommentRepository repository.LikeCommentRepository
}

func NewRedisAdapter(
	commentRepository repository.CommentRepository,
	likeCommentRepository repository.LikeCommentRepository) RedisAdapter {
	ring := redis.NewRing(&redis.RingOptions{Addrs: map[string]string{
		"server_1": config.Config.Redis.Addr,
	}})
	client := cache.New(&cache.Options{Redis: ring})

	return &RedisAdapterImpl{
		client:                client,
		commentRepository:     commentRepository,
		likeCommentRepository: likeCommentRepository,
	}
}

//func (a *RedisAdapterImpl) RefreshCommentsByArticle(articleID int) {
//	var (
//		key = fmt.Sprintf("khumu-comment:comments:article=%d", articleID)
//		val = data.CommentEntities(make([]*ent.Comment, 0))
//	)
//	val, err := a.commentRepository.FindAllParentsByArticleID(articleID)
//	if err != nil {
//		log.Error(err)
//		return
//	}
//
//	log.Infof("캐시를 refresh합니다. key=%s", key)
//	err = a.cache.Set(&cache.Item{
//		Ctx:   context.TODO(),
//		Key:   key,
//		Value: &val,
//	})
//	if err != nil {
//		log.Error(err)
//	}
//}

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
		// 캐시 미스
		//if errors.Is(err, cache.ErrCacheMiss) {
		//	err = a.cache.Get(context.TODO(), key, &val)
		//	if err != nil {
		//		log.Errorf("캐시 미스 후 RefreshCommentsByArticle 했지만 에러 발생: %v", err)
		//	}
		//} else {
		//	log.Error(err)
		//	return []*ent.Comment{}
		//}
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

//func (a *RedisAdapterImpl) RefreshLikeCommentsByComment(commentID int) {
//	var (
//		key = fmt.Sprintf("khumu-comment:likeComments:comment=%d", commentID)
//		val = data.LikeCommentEntities(make([]*ent.LikeComment, 0))
//	)
//	val, err := a.likeCommentRepository.FindAllByCommentID(commentID)
//	if err != nil {
//		log.Error(err)
//		return
//	}
//
//	log.Infof("캐시를 refresh합니다. key=%s", key)
//	err = a.cache.Set(&cache.Item{
//		Ctx:   context.TODO(),
//		Key:   key,
//		Value: &val,
//	})
//	if err != nil {
//		log.Error(err)
//	}
//}

func (a *RedisAdapterImpl) FindAllByCommentID(commentID int) (data.LikeCommentEntities, error) {
	var (
		key = fmt.Sprintf("khumu-comment:likeComments:comment=%d", commentID)
		val = data.LikeCommentEntities(make([]*ent.LikeComment, 0))
	)
	log.Infof("comment=%d 대한 댓글 좋아요를 조회합니다.key=%s", commentID, key)
	err := a.client.Get(context.TODO(), key, &val)
	if err != nil {
		return nil, err
		// 캐시 미스
		//if errors.Is(err, cache.ErrCacheMiss) {
		//	a.RefreshLikeCommentsByComment(commentID)
		//	err = a.cache.Get(context.TODO(), key, &val)
		//	if err != nil {
		//		log.Errorf("캐시 미스 후 RefreshLikeCommentsByComment 했지만 에러 발생: %v", err)
		//	}
		//} else {
		//	log.Error(err)
		//	return []*ent.LikeComment{}
		//}
	}

	return val, nil
}
