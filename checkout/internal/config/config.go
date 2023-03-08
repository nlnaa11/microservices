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

type ServiceInfo struct {
	HostName string `yaml:"host_name"`
	Port     int    `yaml:"port"`
}

type ConfigStruct struct {
	ProductToken string `yaml:"product_token"`
	HostName     string `yaml:"host_name"`
	HttpPort     int    `yaml:"http_port"`
	GrpcPort     int    `yaml:"grpc_port"`
	Services     struct {
		Loms    ServiceInfo `yaml:"loms"`
		Product ServiceInfo `yaml:"product"`
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
	return c.ProductToken
}

type Communication uint8

const (
	HTTP Communication = iota
	GRPC
)

func (c ConfigStruct) GetCommunicationAddress(communication Communication) string {
	switch communication {
	case HTTP:
		return fmt.Sprintf("%s:%d", c.HostName, c.HttpPort)
	case GRPC:
		return fmt.Sprintf("%s:%d", c.HostName, c.GrpcPort)
	default:
		return fmt.Sprintf("%s:%d", c.HostName, defaultPort)
	}
}

type Service uint8

const (
	Loms Service = iota
	Product
)

func (c ConfigStruct) GetServiceAddress(service Service) string {
	switch service {
	case Loms:
		return fmt.Sprintf("%s:%d", c.Services.Loms.HostName, c.Services.Loms.Port)
	case Product:
		return fmt.Sprintf("%s:%d", c.Services.Product.HostName, c.Services.Product.Port)
	default:
		return ""
	}
}
