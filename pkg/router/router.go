package router

import (
	"github.com/ArifulProtik/BlackPen/controller"
	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/services"
	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Group, s *services.Service, auth *auth.Token, key string) {
	handler := controller.New(s, auth)

	r.POST("/signup", handler.Auth.Signup)
	r.POST("/signin", handler.Auth.Signin)

	r.GET("/refresh", handler.Auth.Refresh)
	r.GET("/notes/:page", handler.Note.GetAllNote)
	r.GET("/:slug", handler.Note.GetSingleNote)
	r.GET("/:username/:page", handler.Note.NoteByUser)
	r.GET("/comments/:noteid", handler.Comment.GetALLComment)

	r.Use(handler.Auth.IsAuth)
	r.POST("/note/create", handler.Note.CreateNote)
	r.PUT("/note/update", handler.Note.UpdateNote)
	r.GET("/mynotes/:page", handler.Note.MyNotes)
	r.DELETE("/note/:slug", handler.Note.DeleteNote)

	r.POST("/comment/create", handler.Comment.CreateComment)
	r.DELETE("/comment/:id", handler.Comment.DeleteComment)
	r.GET("/reaction/:noteid", handler.Reaction.CreateReaction)
	r.DELETE("/reaction/delete/:id", handler.Reaction.DeleteReaction)
	r.GET("/logout", handler.Auth.Logout)
}
