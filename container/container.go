package container

import (
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/external/khumu"
	"github.com/khu-dev/khumu-comment/http"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"go.uber.org/dig"
	"log"
)

func Build() *dig.Container {
	c := dig.New()

	// Provide DB Connection

	err := c.Provide(repository.NewEnt)
	if err != nil {
		log.Panic(err)
	}

	// sns
	err = c.Provide(external.NewSnsClient)
	if err != nil {
		log.Panic(err)
	}

	err = c.Provide(external.NewRedisAdapter)
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
	err = c.Provide(http.NewEcho)
	if err != nil {
		log.Panic(err)
	}
	return c
}
