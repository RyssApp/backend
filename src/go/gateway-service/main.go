package main

import (
	"github.com/ryssapp/backend/src/go/common/env"
	"github.com/ryssapp/backend/src/go/common/log"
	"github.com/ryssapp/backend/src/go/gateway-service/service"
)

func main() {
	env.Init()
	log.Init()
	service.Start()
}
