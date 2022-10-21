package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserMetadata holds the schema definition for the UserMetadata entity.
type UserMetadata struct {
	ent.Schema
}

// Fields of the UserMetadata.
func (UserMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name_color"), // hex
		field.String("title"),
		field.String("gender"),
		field.Time("last_online"),
		field.Uint64("user_id"),
	}
}

// Edges of the UserMetadata.
func (UserMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("metadata").Unique().Required().Field("user_id"),
	}
}
