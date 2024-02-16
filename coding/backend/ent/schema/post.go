package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	// TODO: Update the fields to allow expiration filtering.
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("id"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Text("content").NotEmpty(),
		field.Text("title").NotEmpty(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
