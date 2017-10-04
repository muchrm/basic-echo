package http

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type APIConfig struct {
	Prefix string `yaml:"prefix"`
}
type Config struct {
	Port string    `yaml:"port"`
	API  APIConfig `api:"metadata"`
}

func ParseYaml(configPath string) (*Config, error) {
	filename, _ := filepath.Abs(configPath)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
