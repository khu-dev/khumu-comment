package repository

import (
    "context"
    "github.com/khu-dev/khumu-comment/config"
    "github.com/khu-dev/khumu-comment/model"
	"github.com/go-redis/redis/v8"
    "github.com/sirupsen/logrus"
)

type (
	EventMessageRepository interface{
        PublishCommentEvent(message *model.EventMessage)
        publishCommentEvent(message *model.EventMessage) error
        //PublishLikeCommentEvent(message *model.EventMessage)
    }
    RedisEventMessageRepository struct{
        client *redis.Client
        ctx context.Context
    }
)

func NewRedisEventMessageRepository() EventMessageRepository{
    h := &RedisEventMessageRepository{
        client: redis.NewClient(&redis.Options{
            Addr: config.Config.Redis.Address,
            Password: config.Config.Redis.Password,
            DB: config.Config.Redis.DB,
        }),
        ctx: context.Background(),
    }

    return h
}

// 외부에서 handler을 사용할 때에는 error을 신경쓰지 않았으면 좋겠어서
// unexported method를 이용해 error을 리턴하도록 구현.
func (h RedisEventMessageRepository) PublishCommentEvent(message *model.EventMessage) {
    err := h.publishCommentEvent(message)
    if err != nil{
        logrus.Error(err)
    }
}

// 실질적인 publish 작업
func (h RedisEventMessageRepository) publishCommentEvent(message *model.EventMessage) error{
    err := h.client.Publish(h.ctx, config.Config.Redis.CommentChannel, message).Err()
    logrus.Println(config.Config.Redis.CommentChannel)
    return err
}