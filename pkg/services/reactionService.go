package services

import (
	"context"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/ent/love"
	"github.com/google/uuid"
)

type ReactionService struct {
	Client *ent.LoveClient
}

func (r *ReactionService) CreateReaction(noteid uuid.UUID, usr ent.User) error {
	_, err := r.Client.Create().SetNoteid(noteid).
		SetUser(&usr).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (r *ReactionService) DeleteReaction(id int) error {
	_, err := r.Client.Delete().Where(love.IDEQ(id)).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (r *ReactionService) GetReactions(noteid uuid.UUID) ([]*ent.Love, error) {
	ls, err := r.Client.Query().Where(love.NoteidEQ(noteid)).WithUser().
		All(context.Background())
	if err != nil {
		return nil, nil
	}
	return ls, nil
}

func (r *ReactionService) GetReactionByID(id int) (*ent.Love, error) {
	rr, err := r.Client.Query().Where(love.IDEQ(id)).WithUser().
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return rr, nil

}
