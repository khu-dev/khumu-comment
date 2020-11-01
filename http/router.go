package http

import (
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
)

// 의존성 주입시에 root router를 판별하기 위해 임베딩
type RootRouter struct{*echo.Group}

func NewEcho(userRepository repository.UserRepositoryInterface, commentUC usecase.CommentUseCaseInterface) *echo.Echo {
	e := echo.New()

	authenticator := &Authenticator{UserRepository: userRepository}
	e.Use(authenticator.Authenticate)
	e.GET("", func(c echo.Context) error { return c.Redirect(301, "/api") })
	root := NewRootRouter(e)
	_ = NewCommentRouter(root, commentUC)
	return e
}

func NewRootRouter(echoServer *echo.Echo) *RootRouter{
	g := RootRouter{Group: echoServer.Group("/api")}
	g.GET("/", serveHome)
	return &g
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
