package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const configPath = "config.yaml"

type ConfigStruct struct {
	Token    string `yaml:"token"`
	Port     string `yaml:"port"`
	Services struct {
		Product string `yaml:"product"`
	} `yaml:"services"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYAML, err := os.ReadFile(configPath)
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}

func (c ConfigStruct) GetToken() string {
	return c.Token
}

func (c ConfigStruct) GetPort() string {
	return c.Port
}
