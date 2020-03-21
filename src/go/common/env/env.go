package env

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		zap.L().Warn("Starting without .env configuration.", zap.Error(err))
	}
}
