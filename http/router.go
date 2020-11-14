package http

import (
	_ "github.com/khu-dev/khumu-comment/docs"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// 의존성 주입시에 root router를 판별하기 위해 임베딩
type RootRouter struct{*echo.Group}

func NewEcho(userRepository repository.UserRepositoryInterface,
	commentUC usecase.CommentUseCaseInterface,
	likeUC usecase.LikeCommentUseCaseInterface) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("", func(c echo.Context) error { return c.Redirect(301, "/api") })
	e.GET("/healthz", func(c echo.Context) error { return c.String(200, "OK") })
	e.GET("/api/comments/swagger/*", echoSwagger.WrapHandler)
	root := NewRootRouter(e, userRepository)
	_ = NewCommentRouter(root, commentUC)
	_ = NewLikeCommentRouter(root, likeUC)
	return e
}

func NewRootRouter(echoServer *echo.Echo, userRepository repository.UserRepositoryInterface) *RootRouter{
	g := RootRouter{Group: echoServer.Group("/api")}
	authenticator := &Authenticator{UserRepository: userRepository}
	g.Use(authenticator.Authenticate)
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

