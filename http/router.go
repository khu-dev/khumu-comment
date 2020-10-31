package http

import (
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
)

func NewEcho(userRepository repository.UserRepository, commentUC *usecase.CommentUseCase) *echo.Echo {
	e := echo.New()

	authenticator := &Authenticator{UserRepository: userRepository}
	e.Use(authenticator.Authenticate)

	e.GET("", func(c echo.Context) error { return c.Redirect(301, "/api") })

	apiRouterGroup := e.Group("/api")
	apiRouterGroup.GET("", serveHome)

	_ = NewCommentRouter(apiRouterGroup, commentUC)

	return e
}

func serveHome(c echo.Context) error {
	c.Path()
	return c.HTML(200, `
<h2>KHUMU comment REST APIs</h2>
<br/><br/>
<ul>		
	<li>Comment List <a href="http://localhost:9000/api/comments">http://localhost:9000/api/comments</a>
	`)
}
