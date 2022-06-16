package utils

import "github.com/google/uuid"

type UserInput struct {
	Name             string `json:"name" validate:"required"`
	Username         string `json:"username" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=8" `
	Confirm_password string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserSigninInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
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
