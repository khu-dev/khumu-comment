package container

import (
	"github.com/khu-dev/khumu-comment/infra"
	"github.com/khu-dev/khumu-comment/infra/khumu"
	"github.com/khu-dev/khumu-comment/infra/message"
	"github.com/khu-dev/khumu-comment/infra/rest"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/repository/cache"
	"github.com/khu-dev/khumu-comment/usecase"
	"go.uber.org/dig"
	"log"
	"os"
)

func Build(termSig <-chan os.Signal) *dig.Container {
	c := dig.New()

	// os의 signal을 받아서 종료할 수 있게해주는 channel
	err := c.Provide(func() <-chan os.Signal {
		return termSig
	})

	// Provide DB Connection
	err = c.Provide(repository.NewEnt)
	if err != nil {
		log.Panic(err)
	}

	// 캐시에 동기적으로 write할 지 비동기적으로 write할지
	err = c.Provide(func() repository.SynchronousCacheWrite { return false })
	if err != nil {
		log.Panic(err)
	}
	// sns
	err = c.Provide(message.NewSnsMessagePublisher)
	if err != nil {
		log.Panic(err)
	}

	// sqs
	err = c.Provide(message.NewSqsMessageHandler)
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(infra.NewRedisAdapter) //dig.Group("LikeCommentCacheRepository"), dig.Group("CommentCacheRepository")
	if err != nil {
		log.Panic(err)
	}
	err = c.Provide(func(adapter infra.RedisAdapter) (cache.CommentCacheRepository, cache.LikeCommentCacheRepository) {
		return cache.CommentCacheRepository(adapter), cache.LikeCommentCacheRepository(adapter)
	})
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(khumu.NewKhumuAPIAdapter)
	if err != nil {
		log.Panic(err)
	}

	// Provide repository
	err = c.Provide(repository.NewCommentRepository)
	if err != nil {
		log.Panic(err)
	}
	err = c.Provide(repository.NewLikeCommentRepository)
	if err != nil {
		log.Panic(err)
	}

	// Provide usecases
	err = c.Provide(usecase.NewCommentUseCase)
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(usecase.NewLikeCommentUseCase)
	if err != nil {
		log.Panic(err)
	}

	// Provide Echo and routers
	// http내에서 echo와 router 그룹등은 의존성이 없기때문에 한 번에
	// NewEcho에서 Group 생성등까지 처리한다.
	err = c.Provide(rest.NewEcho)
	if err != nil {
		log.Panic(err)
	}
	return c
}
