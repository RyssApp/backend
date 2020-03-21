package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	BindAddress        string `default:"localhost:3501"`
	RedisConnection string
}

func Load() *Config {
	var c Config
	err := envconfig.Process("session_service", &c)
	if err != nil {
		zap.L().Fatal("Failed to create config.", zap.Error(err))
	}
	return &c
}
