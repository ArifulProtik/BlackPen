package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReactionController struct {
	Service *services.ReactionService
}

func (r *ReactionController) CreateReaction(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	noteid := c.Param("noteid")
	uid, err := uuid.Parse(noteid)
	log.Println(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "Note doesnt exsits",
		})
	}

	err = r.Service.CreateReaction(uid, *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg": "Created!!!!!!!!!!",
	})
}

func (r *ReactionController) DeleteReaction(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	id := c.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"msg": "Not Found",
		})
	}
	reaction, err := r.Service.GetReactionByID(iid)
	if err != nil {

		return c.JSON(http.StatusNotFound, echo.Map{
			"msg": "Not Found",
		})
	}
	if reaction.Edges.User.ID != user.ID {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "Status Unauthorized",
		})
	}
	err = r.Service.DeleteReaction(reaction.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"msg": "deleted",
	})

}
