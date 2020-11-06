package main

import (
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/container"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	for i:=0; i<1000; i++{
		fmt.Println("kkkk")
	}

	Run()
}

func Run() {
	log.Printf("Default config. %#v\n", config.Config)
	cont := container.Build()
	err := cont.Invoke(func(e *echo.Echo){
		e.Logger.Print("Started Server")
		e.Logger.Fatal(e.Start(config.Config.Host + ":" + config.Config.Port))
	})
	if err != nil{
		log.Panic(err)
	}
}
