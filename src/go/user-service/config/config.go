package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	BindAddress string `default:"localhost:3501"`
	Cost    int    `default:13`
}

func Load() *Config {
	var c Config
	err := envconfig.Process("user_service", &c)
	if err != nil {
		zap.L().Fatal("Failed to create config.", zap.Error(err))
	}
	return &c
}
