package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Configuration struct {
	Postgres *Postgres `yaml:"postgres"`
	Server   string    `yaml:"server"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Post     int    `yaml:"post"`
	Dbname   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

var Config *Configuration

func LoadConfiguration(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error: %v", err)
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("error: %v", err)
	}
	Config = &config
	log.Printf("config load succeessfully:%v", Config)
}
