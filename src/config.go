package src

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type ApiConfig struct {
	Prefix string `yaml:"prefix"`
}
type Config struct {
	Port string    `yaml:"port"`
	Api  ApiConfig `api:"metadata"`
}

func ParseYaml() (*Config, error) {
	filename, _ := filepath.Abs("../config.yml")
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
