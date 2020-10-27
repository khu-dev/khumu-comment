package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)



func Load() *KhumuConfig{
    filename, err := filepath.Rel("", "config/default.yaml")
    if err != nil{
    	log.Panic(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)
    if err != nil{
    	log.Panic(err)
	}

	var config *KhumuConfig = &KhumuConfig{}
    err = yaml.Unmarshal(yamlFile, config)
    if err != nil {
        panic(err)
    }
    log.Printf("%#v\n",*config)

    return config
}

type KhumuConfig struct {
    Host string `yaml:"host"`
    Port string
    DB struct{
    	Type string
    	SQLite3 struct{
    		FilePath string `yaml:"filePath"`
		} `yaml:"sqlite3"`
	}
}