package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	configPath = "config.yaml"

	defaultPort = 8080
)

type ConfigStruct struct {
	HostName string `yaml:"host_name"`
	HttpPort int    `yaml:"http_port"`
	GrpcPort int    `yaml:"grpc_port"`
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

type Communication uint8

const (
	HTTP Communication = iota
	GRPC
)

func (c ConfigStruct) GetCommunicationAddress(Communication Communication) string {
	switch Communication {
	case HTTP:
		return fmt.Sprintf("%s:%d", c.HostName, c.HttpPort)
	case GRPC:
		return fmt.Sprintf("%s:%d", c.HostName, c.GrpcPort)
	default:
		return fmt.Sprintf("%s:%d", c.HostName, defaultPort)
	}
}
