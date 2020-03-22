package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	BindAddress           string `default:"localhost:3501" split_words:"true"`
	Cost                  int    `default:13`
	PostgresConnection    string `required:"true" split_words:"true"`
	SessionServiceAddress string `default:"localhost:3503" split_words:"true"`
}

func Load() *Config {
	var c Config
	err := envconfig.Process("USER_SERVICE", &c)
	if err != nil {
		zap.L().Fatal("Failed to create config.", zap.Error(err))
	}
	return &c
}
