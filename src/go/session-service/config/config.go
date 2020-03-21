package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	BindAddress   string `default:"localhost:3503"`
	RedisAddress  string `default:"localhost:6379"`
	RedisPassword string `default:""`
	RedisDatabase int    `default:0`
	Secret        string `required:"true"`
}

func Load() *Config {
	var c Config
	err := envconfig.Process("SESSION_SERVICE", &c)
	if err != nil {
		zap.L().Fatal("Failed to create config.", zap.Error(err))
	}
	return &c
}
