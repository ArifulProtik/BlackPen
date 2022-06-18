package services

import (
	"context"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/ent/auth"
	"github.com/google/uuid"
)

type AuthService struct {
	Client *ent.AuthClient
}

func (a *AuthService) CreateSession(s uuid.UUID) (*ent.Auth, error) {
	sessionn, err := a.Client.Create().SetSessionid(s).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return sessionn, nil
}

func (a *AuthService) GetSessionByID(id uuid.UUID) (*ent.Auth, error) {
	sessionn, err := a.Client.Query().Where(auth.IDEQ(id)).
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return sessionn, nil
}

func (a *AuthService) GetSessionBySessionID(id uuid.UUID) (*ent.Auth, error) {
	sessionn, err := a.Client.Query().Where(auth.SessionidEQ(id)).
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return sessionn, nil
}

func (a *AuthService) DeleteSessionByID(id uuid.UUID) error {
	_, err := a.Client.Delete().Where(auth.SessionidEQ(id)).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) UpdateSession(id uuid.UUID) error {
	session, err := a.Client.Query().
		Where(auth.SessionidEQ(id)).First(context.Background())
	if err != nil {
		return err
	}
	session.Update().SetIsBlocked(true).Save(context.Background())
	return nil
}
