package controller

import (
	"log"
	"net/http"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CommentController struct {
	Service *services.CommentService
}

func (s *CommentController) CreateComment(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	var comment utils.CommmentInput
	if c.Request().Body == nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Msg: "Missing JSON data",
		})
	}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Msg: "Invalid JSON",
		})
	}
	v := utils.NewValidator()
	err := v.Struct(comment)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.ToErrResponse(err))
	}
	var newComment utils.CommmentInput
	newComment = utils.CommmentInput{
		NoteID: comment.NoteID,
		Body:   comment.Body,
	}
	nc, err := s.Service.Create(newComment, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, nc)
}

func (s *CommentController) DeleteComment(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Msg: "comment doesnt exist",
		})
	}
	log.Println(uid)

	comment, err := s.Service.GetByID(uid)
	log.Println(comment)
	if err != nil {

		return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
			Msg: "comment doesnt exist",
		})
	}
	if comment.Edges.User.ID != user.ID {

		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "Status Unauthorized",
		})
	}

	err = s.Service.DeleteComment(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"msg": "deleted the comment",
	})

}

func (s *CommentController) GetALLComment(c echo.Context) error {
	id := c.Param("noteid")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "Note Not Found",
		})
	}
	comments, err := s.Service.GetComments(uid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, comments)
}
