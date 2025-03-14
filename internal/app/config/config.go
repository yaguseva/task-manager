package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	FileName string `yaml:"fileName"`
	Server   struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

var Config Configuration

func InitConfig() error {
	data, err := os.ReadFile("internal/app/config/config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	return err
}
