package utils

import (
	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/google/uuid"
)

type UserInput struct {
	Name             string `json:"name" validate:"required"`
	Username         string `json:"username" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=8" `
	Confirm_password string `json:"confirm_password" validate:"required,eqfield=Password"`
}
type CommmentInput struct {
	NoteID uuid.UUID `json:"noteid" validate:"required"`
	Body   string    `json:"body" validate:"required"`
	User   ent.User
}
type UserSigninInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type NoteInput struct {
	ID             *uuid.UUID `json:"id"`
	Slug           *string    `json:"slug"`
	Title          string     `json:"title" validate:"required"`
	Body           string     `json:"body" validate:"required"`
	Tags           []string   `json:"tags" validate:"required"`
	Featured_Image string     `json:"img" validate:"required"`
}

type MultipleNotes struct {
	Lastpage    *int         `json:"lastpage"`
	CurrentPage *int         `json:"currentpage"`
	Notes       []*ent.Notes `json:"notes"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}
type SuccessResponse struct {
	AccessToken  *string                `json:"access_token"`
	RefreshToken *string                `json:"refresh_token"`
	Data         map[string]interface{} `json:"data"`
}
type SessionData struct {
	ID uuid.UUID
	Ip string
}
