package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	DefaultKhumuHome string = "/home/jinsu/git/khumu/khumu-comment"
)

var (
	Config *KhumuConfig
	Location *time.Location
)

func init() {
	Load()
	l, err := time.LoadLocation("Asia/Seoul")
	if err != nil{
		log.Fatal(err)
	}
	Location = l
}

// env.KHUMU_HOME의 경로로부터 config/local.yaml or config/dev.yaml을 읽어옵니다.
func readConfigFileFromKhumuHome(relPath string) *[]byte {
	if Config == nil {
		Config = &KhumuConfig{}
	}

	khumuHome := os.Getenv("KHUMU_HOME")
	if khumuHome == "" {
		khumuHome = DefaultKhumuHome
	}
	logrus.Print("KHUMU_HOME is ", khumuHome)

	filename, err := filepath.Abs(path.Join(khumuHome, relPath))
	logrus.Println("Open config file ", filename)
	if err != nil {
		logrus.Panic(err)
	}

	yamlFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Panic(err)
	}

	return &yamlFileData
}
func Load() {
	var configRelPath string = "config/local.yaml" // 기본은 local.yaml
	if os.Getenv("KHUMU_ENVIRONMENT") == "DEV"{
		// KHUMU_HOME을 기준으로한 대한 상대 경로
		configRelPath = "config/dev.yaml"
	}
	yamlDataPtr := readConfigFileFromKhumuHome(configRelPath)
	Config = &KhumuConfig{}
	err := yaml.Unmarshal(*yamlDataPtr, Config)
	if err != nil {
		logrus.Panic(err)
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
		MySQL struct{
			Host string `yaml:"host"`
			Port int `yaml:"port"`
			DatabaseName string `yaml:"databaseName"`
			User string `yaml:"user"`
			Password string `yaml:"password"`
		}
	}
}
