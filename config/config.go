package config

import (
	"github.com/umi0410/ezconfig"
	"log"
	"time"
)

var (
	Config   *KhumuConfig
	Location *time.Location
	// 개발 단계에서 편의상 개발자들의 home path를 설정
	devKhumuConfigPath []string = []string{
		"./config",
		"/home/jinsu/workspace/khumu/khumu-comment/config",
	}
)

func init() {
	l, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatal(err)
	}
	Location = l

	Config =  &KhumuConfig{}
	ezconfig.LoadConfig("KHUMU", Config)
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
	Sns struct{
		TopicArn string
	}
}
