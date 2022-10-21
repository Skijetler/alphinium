package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Unique(),
		field.String("email").Unique(),
		field.String("password"),
		field.Time("registration_date").Immutable(),
		field.Bool("disabled").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", UserMetadata.Type).Unique(),
		edge.To("posts", Post.Type),
		edge.To("attachments", Attachment.Type),
	}
}
