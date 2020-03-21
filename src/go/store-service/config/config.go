package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	BindAddress        string `default:"localhost:3502" split_words:"true"`
	PostgresConnection string `split_words:"true"`
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	var c Config
	err = envconfig.Process("STORE_SERVICE", &c)
	if err != nil {
		log.Fatal(err)
	}
	return &c
}
