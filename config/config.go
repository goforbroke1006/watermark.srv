package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Config struct {
	InputDir  string `yaml:"inputDir"`
	OutputDir string `yaml:"outputDir"`
}

func LoadConfig(path string) (Config, error) {
	bytes, err := ioutil.ReadFile(path)

	cfg := Config{}

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
