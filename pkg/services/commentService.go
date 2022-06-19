package services

import (
	"context"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/ent/comment"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/google/uuid"
)

type CommentService struct {
	Client *ent.CommentClient
}

func (c *CommentService) Create(newcomment utils.CommmentInput, usr *ent.User) (*ent.Comment, error) {

	comment, err := c.Client.Create().SetBody(newcomment.Body).
		SetNoteID(newcomment.NoteID).SetUser(usr).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentService) DeleteComment(id uuid.UUID) error {
	err := c.Client.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentService) GetByID(id uuid.UUID) (*ent.Comment, error) {
	comment, err := c.Client.Query().Where(comment.IDEQ(id)).
		WithUser().First(context.Background())
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentService) GetComments(noteid uuid.UUID) ([]*ent.Comment, error) {
	comments, err := c.Client.Query().Where(comment.NoteIDEQ(noteid)).
		WithUser().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return comments, nil
}
