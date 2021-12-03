package main

import (
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/container"
	"github.com/khu-dev/khumu-comment/infra/message"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

func init() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		DisableQuote:  true,
		ForceColors:   true,
		// line을 깔끔하게 보여줌.
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := strings.Replace(f.File, workingDir+"/", "", -1)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		FullTimestamp:   false,
		TimestampFormat: "2006/01/03 15:04:05",
	})
}
func main() {
	log.Println("Args: ", len(os.Args), os.Args)
	log.Printf("Default config. %#v\n", config.Config)

	termSig := make(chan os.Signal, 3)
	signal.Notify(termSig, syscall.SIGINT, syscall.SIGTERM)

	cont := container.Build(termSig)

	go func() {
		err := cont.Invoke(func(h message.MessageHandler) {
			h.Listen()
		})
		log.Fatal(err)
	}()

	err := cont.Invoke(func(e *echo.Echo) {
		e.Logger.Print("Started Server")
		e.Logger.Fatal(e.Start(config.Config.Host + ":" + config.Config.Port))
	})
	if err != nil {
		log.Panic(err)
	}
}
