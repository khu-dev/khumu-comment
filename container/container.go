package container

import (
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/http"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"go.uber.org/dig"
	"log"
)

func Build() *dig.Container {
	c := dig.New()

	// Provide DB Connection
	err := c.Provide(repository.NewGorm)
	if err != nil {
		log.Panic(err)
	}

	// sns
	err = c.Provide(external.NewSnsClient)
	if err != nil {
		log.Panic(err)
	}

	// Provide repositories
	err = c.Provide(repository.NewCommentRepositoryGorm)
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(repository.NewLikeCommentRepositoryGorm)
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(repository.NewUserRepositoryGorm)
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(repository.NewRedisEventMessageRepository)
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
	err = c.Provide(http.NewEcho)
	if err != nil {
		log.Panic(err)
	}
	return c
}
