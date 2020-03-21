package main

import (
	"github.com/ryssapp/backend/src/go/common/env"
	"github.com/ryssapp/backend/src/go/common/log"
)

func main() {
	log.Init()
	env.Init()
	// TODO: Start service here
}
