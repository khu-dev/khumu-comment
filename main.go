package main

import (
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/container"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

func init(){
	workingDir, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		DisableQuote: true,
		ForceColors: true,
		// line을 깔끔하게 보여줌.
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := strings.Replace(f.File, workingDir + "/", "", -1)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
}
func main() {

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
