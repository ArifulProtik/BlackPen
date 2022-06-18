package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("sessionid", uuid.UUID{}),
		field.Bool("is_blocked").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return nil
}
