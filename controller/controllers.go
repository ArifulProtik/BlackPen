package controller

import (
	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/services"
)

type Controller struct {
	Auth *AuthController
}

func New(services *services.Service, AuthToken *auth.Token) *Controller {
	return &Controller{
		Auth: &AuthController{
			UserService: services.User,
			AuthService: services.Auth,
			AuthToken:   AuthToken,
		},
	}
}
