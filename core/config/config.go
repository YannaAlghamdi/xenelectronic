package config

import (
	"io/ioutil"

	"go.uber.org/zap"

	yaml "gopkg.in/yaml.v2"
)

var logger, _ = zap.NewProduction()

func init() {
	defer logger.Sync()
}

// NewConfig parses the config file from the given path
func NewConfig(path string) (*Config, error) {
	cfg := &Config{}
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if err := yaml.Unmarshal(contents, cfg); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return cfg, nil
}

// Config holds app config
type Config struct {
	HTTP struct {
		Port string
	} `yaml:"http"`
	Postgres struct {
		Host     string
		Port     string
		Db       string
		User     string
		Password string
		SSLMode  string `yaml:"sslMode"`
	} `yaml:"postgres"`
}
