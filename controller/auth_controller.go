package controller

import (
	"net/http"
	"strings"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	UserService *services.UserService
	AuthService *services.AuthService
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
		a.AuthService.DeleteSessionByID(user.ID)

		newSession, err := a.AuthService.CreateSession(user.ID)
		if err != nil {
			return e.JSON(500, err)
		}
		_, rfToken := a.AuthToken.RfTokenWithSession(newSession.ID)

		_, token := a.AuthToken.TokenWithUser(user.ID)

		return e.JSON(http.StatusAccepted, utils.SuccessResponse{
			AccessToken:  &token,
			RefreshToken: &rfToken,
			Data: echo.Map{
				"user": user,
			},
		})

	}
	return e.JSON(http.StatusBadRequest, utils.ErrorResponse{
		Msg: "OOps not Allowed",
	})
}

func (a *AuthController) Refresh(e echo.Context) error {
	token := e.Request().Header.Get("token")
	if token != "" {
		err, id := a.AuthToken.VerifyToken(token)
		if err != nil {
			return e.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "Invalid Token",
			})
		}

		session, err := a.AuthService.GetSessionByID(id)

		if session.IsBlocked != true {
			_, newtoken := a.AuthToken.TokenWithUser(session.Sessionid)
			return e.JSON(http.StatusAccepted, utils.SuccessResponse{
				AccessToken: &newtoken,
			})
		}
		return e.JSON(http.StatusAccepted, utils.ErrorResponse{
			Msg: "Invalid Token",
		})
	}
	return e.JSON(http.StatusUnauthorized, utils.ErrorResponse{
		Msg: "Status Unauthorized",
	})
}
func (a *AuthController) Logout(e echo.Context) error {
	id := e.Get("id").(*ent.User)
	err := a.AuthService.UpdateSession(id.ID)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "Unauthorized",
		})
	}
	return e.JSON(http.StatusAccepted, echo.Map{
		"msg": "logged out user",
	})

}
