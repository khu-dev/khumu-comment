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
	yamlDataPtr := readConfigFileFromKhumuHome("config/default.yaml")
	Config = &KhumuConfig{}
	err := yaml.Unmarshal(*yamlDataPtr, Config)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Default Config: %#v\n", *Config)
}

func LoadTestConfig() {
	yamlDataPtr := readConfigFileFromKhumuHome("config/test.yaml")

	err := yaml.Unmarshal(*yamlDataPtr, Config)
	if err != nil {
		panic(err)
	}
	log.Printf("Test Config: %#v\n", *Config)
}

type KhumuConfig struct {
	Host             string `yaml:"host"`
	RestRootEndpoint string `yaml:"restRootEndpoint"`
	Port             string
	DB               struct {
		Type    string
		SQLite3 struct {
			FilePath string `yaml:"filePath"`
		} `yaml:"sqlite3"`
	}
}
