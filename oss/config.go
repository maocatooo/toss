package oss

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	AK       string `yaml:"AK"`
	SK       string `yaml:"SK"`
	Region   string `yaml:"region"`
	Bucket   string `yaml:"bucket"`
	Endpoint string `yaml:"endpoint"`
	Source   string `yaml:"source"`
	Target   string `yaml:"target"`
}

func ConfigFromFile(path string) Config {
	var c Config
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &c)
	if err != nil {
		panic(err)
	}
	return c
}
