package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/labstack/echo/v4"
)

type NoteController struct {
	Service         *services.NoteService
	UserService     *services.UserService
	ReactionService *services.ReactionService
}

func (n *NoteController) CreateNote(c echo.Context) error {
	if c.Request().Body != nil {
		var nn utils.NoteInput
		if err := c.Bind(&nn); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Msg: "JSON data missing",
			})
		}
		v := utils.NewValidator()
		err := v.Struct(nn)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.ToErrResponse(err))
		}
		u := c.Get("user").(*ent.User)
		if u == nil {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "Status Unauthorized",
			})
		}

		s := utils.GenSlug(nn.Title)
		newNote, err := n.Service.CreateNote(nn, *u, s)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, newNote)

	}
	return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
		Msg: "Json Data Missing",
	})
}

func (n *NoteController) UpdateNote(c echo.Context) error {
	if c.Request().Body != nil {

		var nn utils.NoteInput
		if err := c.Bind(&nn); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Msg: "JSON data missing",
			})
		}
		if nn.ID.String() == "" || *nn.Slug == "" {
			return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{})
		}
		v := utils.NewValidator()
		err := v.Struct(nn)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.ToErrResponse(err))
		}
		u := c.Get("user").(*ent.User)
		if u == nil {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "Status Unauthorized",
			})
		}

		note, err := n.Service.GetNoteByID(*nn.ID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "No Note Found by The ID",
			})
		}
		if note.Edges.Author.ID != u.ID {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
				Msg: "You do not have the permission to edit",
			})
		}
		n2, err2 := n.Service.UpdateNote(note.ID, nn)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, err2.Error())
		}
		return c.JSON(http.StatusCreated, n2)

	}
	return c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
		Msg: "Missing Data",
	})
}

func (n *NoteController) GetAllNote(c echo.Context) error {
	currentPage := c.Param("page")
	newpage, err := strconv.Atoi(currentPage)
	newpage2 := (newpage - 1) * 10
	count, notes, err := n.Service.AllNotes(newpage2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	lastpage := int(math.Ceil(float64(count) / float64(10)))
	return c.JSON(http.StatusOK, utils.MultipleNotes{
		Lastpage:    &lastpage,
		CurrentPage: &newpage,
		Notes:       notes,
	})
}

func (n *NoteController) GetSingleNote(c echo.Context) error {
	slug := c.Param("slug")
	note, err := n.Service.GetNoteBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "Not Found",
		})
	}
	type NoteResp struct {
		LikeCount int `json:"likes"`
		Likes     []*ent.Love
		Note      *ent.Notes `json:"note"`
	}
	likes, err := n.ReactionService.GetReactions(note.ID)

	return c.JSON(http.StatusOK, &NoteResp{
		LikeCount: len(likes),
		Likes:     likes,
		Note:      note,
	})
}

func (n *NoteController) MyNotes(c echo.Context) error {
	u := c.Get("user").(*ent.User)

	currentPage := c.Param("page")
	newpage, err := strconv.Atoi(currentPage)
	newpage2 := (newpage - 1) * 10
	count, notes, err := n.Service.NoteByUserID(u.ID, newpage2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	lastpage := int(math.Ceil(float64(count) / float64(10)))
	return c.JSON(http.StatusOK, utils.MultipleNotes{
		Lastpage:    &lastpage,
		CurrentPage: &newpage,
		Notes:       notes,
	})

}

func (n *NoteController) NoteByUser(c echo.Context) error {
	username := c.Param("username")
	user, err := n.UserService.FindUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "User Not Found",
		})
	}

	currentPage := c.Param("page")
	newpage, err := strconv.Atoi(currentPage)
	newpage2 := (newpage - 1) * 10
	count, notes, err := n.Service.NoteByUserID(user.ID, newpage2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	lastpage := int(math.Ceil(float64(count) / float64(10)))
	return c.JSON(http.StatusOK, utils.MultipleNotes{
		Lastpage:    &lastpage,
		CurrentPage: &newpage,
		Notes:       notes,
	})

}

func (n *NoteController) DeleteNote(c echo.Context) error {
	user := c.Get("user").(*ent.User)
	slug := c.Param("slug")

	note, err := n.Service.GetNoteBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Msg: "Not Found",
		})
	}
	if note.Edges.Author.ID != user.ID {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse{
			Msg: "Unauthorized",
		})
	}
	err = n.Service.DeletenoteByID(note.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"msg": "deleted!!",
	})

}
