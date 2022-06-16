package controller

import (
	"net/http"
	"strings"

	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	UserService *services.UserService
	AuthToken   *auth.Token
}

func (a *AuthController) Signup(e echo.Context) error {
	if e.Request().Body != nil {
		var user utils.UserInput

		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, utils.ErrorResponse{
				Msg: "JSOn data missing",
			})
		}
		validator := utils.NewValidator()
		err := validator.Struct(user)
		if err != nil {
			return e.JSON(http.StatusUnprocessableEntity, utils.ToErrResponse(err))

		}
		newUser, err := a.UserService.SaveUser(user)
		if err != nil {
			return e.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Msg: "Email or Password already exists",
			})
		}
		return e.JSON(http.StatusCreated, newUser)

	}
	return e.JSON(http.StatusBadRequest, utils.ErrorResponse{
		Msg: "Json data missing",
	})
}

func (a *AuthController) Signin(e echo.Context) error {
	if e.Request().Body != nil {
		var credentials utils.UserSigninInput
		if err := e.Bind(&credentials); err != nil {
			return e.JSON(http.StatusBadRequest, utils.ErrorResponse{
				Msg: "JSON Data Missing",
			})
		}
		validator := utils.NewValidator()
		err := validator.Struct(credentials)
		if err != nil {
			return e.JSON(http.StatusUnprocessableEntity, utils.ToErrResponse(err))

		}
		user, err := a.UserService.FindUserByEmail(strings.TrimSpace(credentials.Email))
		if err != nil {
			return e.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "Email and Password Doesnt match",
			})
		}
		if user != nil {
			err := utils.VerifyPass(user.Password, credentials.Password)
			if err != nil {
				return e.JSON(http.StatusUnauthorized, utils.ErrorResponse{
					Msg: "Email or Password doesnt match",
				})
			}
		}

		_, token := a.AuthToken.TokenWithUser(user.ID)

		return e.JSON(http.StatusAccepted, utils.SuccessResponse{
			AccessToken: &token,
			Data: echo.Map{
				"user": user,
			},
		})

	}
	return e.JSON(http.StatusBadRequest, utils.ErrorResponse{
		Msg: "OOps not Allowed",
	})
}
