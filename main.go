// @title Docs::KHUMU Comment
// @version 1.0
// @description KHUMU의 Comment와 Comment-Like에 대한 RESTful API server
// @description <h3>KHUMU API Documentations</h3>
// @description <ul>
// @description <li><a href='https://api.khumu.jinsu.me/docs/command-center'>command-center</a>: 인증, 유저, 게시판, 게시물, 게시물 좋아요, 게시물 북마크 등 전반적인 쿠뮤의 API</li>
// @description <li><a href='https://api.khumu.jinsu.me/docs/comment/index.html'>comment</a>: 댓글, 댓글 좋아요와 관련된 API</li>
// @description </ul>

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://github.com/khu-dev
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

package main

import (
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/container"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: false, DisableQuote: true, ForceColors: true})
}
func main() {
	logrus.Println("Args: ", len(os.Args), os.Args)
	Run()
}

func Run() {
	logrus.Printf("Default config. %#v\n", config.Config)
	cont := container.Build()
	err := cont.Invoke(func(e *echo.Echo) {
		e.Logger.Print("Started Server")
		e.Logger.Fatal(e.Start(config.Config.Host + ":" + config.Config.Port))
	})
	if err != nil {
		log.Panic(err)
	}
}
