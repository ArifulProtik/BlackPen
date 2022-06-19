package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Love holds the schema definition for the Love entity.
type Love struct {
	ent.Schema
}

// Fields of the Love.
func (Love) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("noteid", uuid.UUID{}),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Love.
func (Love) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("loves").
			Unique(),
	}
}
