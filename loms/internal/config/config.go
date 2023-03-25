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

type DBInfo struct {
	DSN                string `yaml:"dsn"`
	MaxOpenConnections int32  `yaml:"max_open_connections"`
}

type ConfigData struct {
	HostName string `yaml:"host_name"`
	HttpPort int    `yaml:"http_port"`
	GrpcPort int    `yaml:"grpc_port"`

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

type Communication uint8

const (
	HTTP Communication = iota
	GRPC
)

func (c *ConfigData) GetCommunicationAddress(Communication Communication) string {
	switch Communication {
	case HTTP:
		return fmt.Sprintf("%s:%d", c.HostName, c.HttpPort)
	case GRPC:
		return fmt.Sprintf("%s:%d", c.HostName, c.GrpcPort)
	default:
		return fmt.Sprintf("%s:%d", c.HostName, defaultPort)
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
