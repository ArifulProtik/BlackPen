package services

import (
	"context"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/ent/auth"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/google/uuid"
)

type AuthService struct {
	Client *ent.AuthClient
}

func (a *AuthService) CreateSession(data utils.SessionData) (*ent.Auth, error) {
	newsession, err := a.Client.Create().SetUserID(data.ID).SetIP(data.Ip).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newsession, nil
}
func (a *AuthService) GetSessionByUserID(id uuid.UUID) (*ent.Auth, error) {
	sessionn, err := a.Client.Query().Where(auth.SessionIDEQ(id)).
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return sessionn, nil
}

func (a *AuthService) DeleteSessionByID(id uuid.UUID) error {
	_, err := a.Client.Delete().Where(auth.SessionIDEQ(id)).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *AuthService) UpdateSession(id uuid.UUID) error {
	session, err := a.Client.Query().
		Where(auth.SessionIDEQ(id)).First(context.Background())
	if err != nil {
		return err
	}
	session.Update().SetIsBlocked(true).Save(context.Background())
	return nil
}
