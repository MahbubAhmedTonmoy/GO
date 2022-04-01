package service

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Port             string `yaml:"port"`
		ConnectionString string `yaml:"connectionString"`
		DbName           string `yaml:"dbname"`
		CollcetionName   string `yaml:"collectionName"`
	} `yaml:"database"`
}


func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
