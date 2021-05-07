package main

import (
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/container"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: false, DisableQuote: true, ForceColors: true})
	logrus.Println("Args: ", len(os.Args), os.Args)
	logrus.Printf("Default config. %#v\n", config.Config)
	cont := container.Build()
	err := cont.Invoke(func(e *echo.Echo) {
		e.Logger.Print("Started Server")
		e.Logger.Fatal(e.Start(config.Config.Host + ":" + config.Config.Port))
	})
	if err != nil {
		logrus.Panic(err)
	}
}
