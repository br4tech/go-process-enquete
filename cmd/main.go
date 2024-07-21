package main

import (
	"github.com/br4tech/go-process-enquete/config"
	"github.com/br4tech/go-process-enquete/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewEchoServer(
		&cfg,
	).Start()
}
