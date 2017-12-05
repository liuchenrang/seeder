package config


import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type Account struct{
	Name string
	Password string
	Database string
	Table string
	DBName string `yaml:"dbname"`
}
type ConnectionInfo struct{
	Max int
	Idle int
}
type Database struct{
	Account Account
	Master  []Server
	ConnectionInfo ConnectionInfo `yaml:"connection_info"`
}
type Logger struct{
	Path string `yaml:"path"`
	Level string `yaml:"level"`
	File string `yaml:"file"`
}
type Server struct{
	Port int `yaml:"port"`
	Host string `yaml:"host"`
}
type SeederConfig struct{
	Database Database
	Logger Logger
	Server Server
}

func NewSeederConfig(yamlfile string) SeederConfig{

	seederConfig := SeederConfig{}
	content, err := ioutil.ReadFile(yamlfile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &seederConfig)
	if err != nil {
		log.Fatal(err)
	}
	return  seederConfig
}
