package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("message"),
		field.Time("date").Immutable(),
		field.Uint64("user_id"),
		field.Uint64("thread_id"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("thread", Thread.Type).Ref("posts").Unique().Required().Field("thread_id"),
		edge.From("user", User.Type).Ref("posts").Unique().Required().Field("user_id"),
		edge.To("attachments", Attachment.Type),
	}
}
