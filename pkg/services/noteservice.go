package services

import (
	"context"
	"time"

	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/ent/notes"
	"github.com/ArifulProtik/BlackPen/ent/user"
	"github.com/ArifulProtik/BlackPen/pkg/utils"
	"github.com/google/uuid"
)

type NoteService struct {
	Client *ent.NotesClient
}

func (n *NoteService) CreateNote(NewNote utils.NoteInput, u ent.User, slug string) (*ent.Notes, error) {
	note, err := n.Client.Create().SetAuthor(&u).SetAuthorID(u.ID).SetTitle(NewNote.Title).
		SetBody(NewNote.Body).SetTags(NewNote.Tags).SetFImage(NewNote.Featured_Image).SetSlug(slug).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NoteService) UpdateNote(id uuid.UUID, note utils.NoteInput) (*ent.Notes, error) {
	nn, err := n.Client.UpdateOneID(id).SetTitle(note.Title).SetBody(note.Body).
		SetFImage(note.Featured_Image).SetTags(note.Tags).SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return nn, nil
}

func (n *NoteService) NoteByUserID(u uuid.UUID, page int) (int, []*ent.Notes, error) {
	count, _ := n.Client.Query().Where(notes.HasAuthorWith(user.IDEQ(u))).Count(context.Background())
	notes, err := n.Client.Query().Order(ent.Desc(notes.FieldCreatedAt)).
		Where(notes.HasAuthorWith(user.IDEQ(u))).
		Select(
			notes.FieldTitle,
			notes.FieldFImage,
			notes.FieldSlug,
			notes.FieldTags,
			notes.FieldCreatedAt,
		).WithAuthor().Limit(10).Offset(page).
		All(context.Background())
	if err != nil {
		return 0, nil, err
	}
	return count, notes, err
}

func (n *NoteService) AllNotes(page int) (int, []*ent.Notes, error) {
	count, _ := n.Client.Query().Count(context.Background())
	notes, err := n.Client.Query().Limit(10).Offset(page).
		Select(

			notes.FieldTitle,
			notes.FieldFImage,
			notes.FieldSlug,
			notes.FieldTags,
			notes.FieldCreatedAt,
		).WithAuthor().All(context.Background())
	if err != nil {
		return 0, nil, err
	}
	return count, notes, nil
}

func (n *NoteService) GetNoteBySlug(s string) (*ent.Notes, error) {
	note, err := n.Client.Query().Where(notes.SlugEQ(s)).WithAuthor().
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NoteService) GetNoteByID(id uuid.UUID) (*ent.Notes, error) {
	note, err := n.Client.Query().Where(notes.IDEQ(id)).WithAuthor().
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return note, nil
}
func (n *NoteService) DeletenoteByID(id uuid.UUID) error {
	err := n.Client.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
