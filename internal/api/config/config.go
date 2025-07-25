package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env string "yaml:ENV"
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
