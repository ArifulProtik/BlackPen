package router

import (
	"github.com/ArifulProtik/BlackPen/controller"
	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Group, s *services.Service, auth *auth.Token) {
	handler := controller.New(s, auth)

	r.POST("/signup", handler.Auth.Signup)
	r.POST("/signin", handler.Auth.Signin)
}
