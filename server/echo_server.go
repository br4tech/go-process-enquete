package server

import (
	"fmt"

	"github.com/br4tech/go-process-enquete/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	app *echo.Echo
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config) Server {
	return &EchoServer{
		app: echo.New(),
		cfg: cfg,
	}
}

func (server *EchoServer) Start() {
	server.app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${time_rfc3339} ${status} ${method} ${host} ${path}",
		},
	))

	serverUrl := fmt.Sprintf(":%d", server.cfg.App.Port)
	server.app.Logger.Fatal(server.app.Start(serverUrl))
}
