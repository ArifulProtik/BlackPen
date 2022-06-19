package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) Run() {
	s.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	s.Echo.Logger.Fatal(s.Echo.Start(s.Cfg.Port))
}
