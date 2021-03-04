package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

var (
	Config   *KhumuConfig
	Location *time.Location
	// 개발 단계에서 편의상 개발자들의 home path를 설정
	devKhumuConfigPath []string = []string{
		"/home/jinsu/git/khumu/khumu-comment/config",
	}
)

func init() {
	Load()
	l, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatal(err)
	}
	Location = l
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: false, DisableQuote: true, ForceColors: true})
}

func Load() {
	for _, configPath := range devKhumuConfigPath {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath(os.Getenv("KHUMU_CONFIG_PATH"))
	viper.SetConfigType("yaml")
	khumuEnvironment := strings.ToLower(os.Getenv("KHUMU_ENVIRONMENT"))
	if khumuEnvironment == "" {
		khumuEnvironment = "default"
	}
	viper.SetConfigName(khumuEnvironment)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	Config = new(KhumuConfig)
	err = viper.Unmarshal(Config)
	if err != nil {
		logrus.Fatal(err)
	}
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
}
