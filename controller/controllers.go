package controller

import (
	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/services"
)

type Controller struct {
	Auth    *AuthController
	Note    *NoteController
	Comment *CommentController
}

func New(services *services.Service, AuthToken *auth.Token) *Controller {
	return &Controller{
		Auth: &AuthController{
			UserService: services.User,
			AuthService: services.Auth,
			AuthToken:   AuthToken,
		},
		Note: &NoteController{
			UserService: services.User,
			Service:     services.Note,
		},
		Comment: &CommentController{
			Service: services.Comment,
		},
	}
}
