package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Account struct {
	Name     string
	Password string
	Database string
	Table    string
	DBName   string `yaml:"dbname"`
}
type ConnectionInfo struct {
	MaxOpenConns int `yaml:"max_open_conns"`
	MaxIdleConns int `yaml:"max_idle_conns"`
}
type Database struct {
	Account        Account
	Master         []Server
	ConnectionInfo ConnectionInfo `yaml:"connection_info"`
}

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}
type SeederConfig struct {
	Database Database
	Server   Server
	Monitior Monitior
	Preload []string  `yaml:"preload"`
}
type Monitior struct {
	VigilantValue uint8 `yaml:"vigilant_value"`
}

func NewSeederConfig(yamlfile string) SeederConfig {

	seederConfig := SeederConfig{}
	content, err := ioutil.ReadFile(yamlfile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &seederConfig)
	if err != nil {
		log.Fatal(err)
	}
	return seederConfig
}
