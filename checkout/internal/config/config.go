package config

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
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

type DBInfo struct {
	DSN                string `yaml:"dsn"`
	MaxOpenConnections int32  `yaml:"max_open_connections"`
}

type ConfigData struct {
	ProductToken string `yaml:"product_token"`
	HostName     string `yaml:"host_name"`
	HttpPort     int    `yaml:"http_port"`
	GrpcPort     int    `yaml:"grpc_port"`

	Services struct {
		Loms    ServiceInfo `yaml:"loms"`
		Product ServiceInfo `yaml:"product"`
	} `yaml:"services"`

	DB DBInfo `yaml:"db"`
}

var Data ConfigData

func Init() error {
	rawYAML, err := os.ReadFile(configPath)
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &Data)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}

func (c *ConfigData) GetToken() string {
	return c.ProductToken
}

type Communication uint8

const (
	HTTP Communication = iota
	GRPC
)

func (c *ConfigData) GetCommunicationAddress(communication Communication) string {
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

func (c *ConfigData) GetServiceAddress(service Service) string {
	switch service {
	case Loms:
		return fmt.Sprintf("%s:%d", c.Services.Loms.HostName, c.Services.Loms.Port)
	case Product:
		return fmt.Sprintf("%s:%d", c.Services.Product.HostName, c.Services.Product.Port)
	default:
		return ""
	}
}

func (c *ConfigData) GetDBConfig() (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(c.DB.DSN)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}
