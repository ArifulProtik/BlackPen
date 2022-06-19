package controller

import (
	"net/http"

	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (a *AuthController) IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authheader := c.Request().Header.Get("Authorization")
		if authheader == "" {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "Status Unauthorized",
			})
		}
		bearer := "Bearer "
		token := authheader[len(bearer):]
		err, id := a.AuthToken.VerifyToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: err.Error(),
			})
		}
		session, err := a.AuthService.GetSessionBySessionID(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: err.Error(),
			})
		}
		if session.IsBlocked == true {

			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: err.Error(),
			})
		}
		user, _ := a.UserService.FindUserByID(session.Sessionid)
		c.Set("user", user)
		return next(c)

	}
}
