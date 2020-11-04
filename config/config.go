package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const DefaultKhumuHome string = "/home/jinsu/git/khumu/khumu-comment"

var Config *KhumuConfig

func init() {
	Load()
}

func readConfigFileFromKhumuHome(relPath string) *[]byte {
	if Config == nil {
		Config = &KhumuConfig{}
	}

	khumuHome := os.Getenv("KHUMU_HOME")
	if khumuHome == "" {
		khumuHome = DefaultKhumuHome
	}
	log.Print("KHUMU_HOME is ", khumuHome)

	filename, err := filepath.Abs(path.Join(khumuHome, relPath))
	log.Println("Open config file ", filename)
	if err != nil {
		log.Panic(err)
	}

	yamlFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
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
		log.Panic(err)
	}
}

func LoadTestConfig() {
	yamlDataPtr := readConfigFileFromKhumuHome("config/test.yaml")

	err := yaml.Unmarshal(*yamlDataPtr, Config)
	if err != nil {
		panic(err)
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
