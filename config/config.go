package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Config struct {
	WatermarkFile string   `yaml:"watermark"`
	Directories   []string `yaml:"directories"`
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
