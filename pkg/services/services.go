package services

import (
	"github.com/ArifulProtik/BlackPen/ent"
)

type Service struct {
	Auth    *AuthService
	User    *UserService
	Note    *NoteService
	Comment *CommentService
}

func New(dbclient *ent.Client) *Service {
	return &Service{
		Auth: &AuthService{
			Client: dbclient.Auth,
		},
		User: &UserService{
			Client: dbclient.User,
		},
		Note: &NoteService{
			Client: dbclient.Notes,
		},
		Comment: &CommentService{
			Client: dbclient.Comment,
		},
	}
}
