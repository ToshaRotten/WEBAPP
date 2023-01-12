package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		LogLevel int    `yaml:"logLevel"`
	} `yaml:"ServerConfig"`
}

func New() *Config {
	return &Config{}
}

func TakeFromFile(path string) (*Config, error) {
	c := Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, err
	}
	return &c, err
}
