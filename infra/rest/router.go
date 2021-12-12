package rest

import (
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// 의존성 주입시에 root router를 판별하기 위해 임베딩
type RootRouter struct{ *echo.Group }

func NewEcho(commentUC usecase.CommentUseCaseInterface,
	likeUC usecase.LikeCommentUseCaseInterface,
	articleUC usecase.ArticleUseCase,
	repo *ent.Client) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = CustomHTTPErrorHandler
	e.Pre(middleware.RemoveTrailingSlash()) // 이거 안하면 꼭 끝에 /를 붙여야할 수도 있음.
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${status} uri=${uri} latency=${latency}\n",
		Skipper: func(context echo.Context) bool {
			// health check log는 너무 verbose함.
			if context.Request().URL.RequestURI() == "/healthz" {
				return true
			}
			return false
		},
	}))
	e.Use(KhumuRequestLog)
	e.GET("", func(c echo.Context) error { return c.Redirect(301, "/api") })
	e.GET("/healthz", func(c echo.Context) error { return c.String(200, "OK") })
	e.GET("/docs/comment/*", echoSwagger.WrapHandler)
	root := NewRootRouter(e, repo)
	_ = NewCommentRouter(root, commentUC, likeUC, articleUC)
	return e
}

func NewRootRouter(echoServer *echo.Echo, repo *ent.Client) *RootRouter {
	g := RootRouter{Group: echoServer.Group("/api")}
	authenticator := &Authenticator{Repo: repo}
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
