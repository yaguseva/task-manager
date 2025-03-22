package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	FileName string `yaml:"file_name"`
	DBPath   string `yaml:"db_path"`
	Server   struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

var Config Configuration

func Initialize() error {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	return err
}
