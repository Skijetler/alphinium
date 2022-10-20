package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("link"),
		field.String("name"),
		field.String("size"),
		field.String("type"),
		field.String("hash"),
		field.Uint64("user_id"),
		field.Uint64("post_id").Optional(),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("attachments").Unique().Required().Field("user_id"),
		edge.From("post", Post.Type).Ref("attachments").Unique().Field("post_id"),
	}
}
