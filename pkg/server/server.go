package server

import (
	"github.com/ArifulProtik/BlackPen/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Cfg  *config.ServerConfig
	Echo *echo.Echo
}

func New(cfg *config.ServerConfig) *Server {
	return &Server{
		Cfg:  cfg,
		Echo: echo.New(),
	}
}
