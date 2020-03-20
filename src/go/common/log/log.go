package log

import (
	"go.uber.org/zap"
	"log"
)

func Init() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(l)
}
