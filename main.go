// @title Docs::KHUMU Comment
// @version 1.0
// @description KHUMU의 Comment와 Comment-Like에 대한 RESTful API server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://github.com/khu-dev
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

package main
import (
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/container"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	fmt.Println("Args: ", len(os.Args), os.Args)
	if len(os.Args) == 1{
		Run()
	} else{
		if os.Args[1] == "run"{
			Run()
		} else if os.Args[1] == "migrate"{
			//config.Load()
			db := repository.NewGorm()
			err := repository.MigrateMinimum(db)
			if err != nil{
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Successfully migrated db.")
		}
	}
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
