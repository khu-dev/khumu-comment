package config

import (
	"github.com/umi0410/ezconfig"
	"log"
	"os"
	"strings"
	"time"
)

var (
	Config   *KhumuConfig
	Location *time.Location
)

func init() {
	l, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatal(err)
	}
	Location = l

	os.Setenv("KHUMU_DB.MYSQL.HOST", "123123")
	Config = &KhumuConfig{}
	wd, _ := os.Getwd()
	ezconfig.LoadConfig("KHUMU", Config, []string{wd, strings.ToLower(os.Getenv("KHUMU_CONFIG_PATH"))})
}

type KhumuConfig struct {
	Host             string `yaml:"host"`
	RestRootEndpoint string `yaml:"restRootEndpoint"`
	Port             string
	DB               struct {
		Kind    string `yaml:"kind"`
		SQLite3 struct {
			FilePath string `yaml:"filePath"`
		} `yaml:"sqlite3"`
		MySQL struct {
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			DatabaseName string `yaml:"databaseName"`
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
		}
	}
	Redis struct {
		Address            string
		Password           string
		DB                 int
		CommentChannel     string
		LikeCommentChannel string
	}
	Sns struct {
		TopicArn string
	}
}
