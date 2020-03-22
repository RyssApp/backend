package main

import (
	"github.com/ryssapp/backend/src/go/common/env"
	"github.com/ryssapp/backend/src/go/common/log"
	"github.com/ryssapp/backend/src/go/user-service/service"
)

func main() {
	log.Init()
	env.Init()
	service.Start()
}
