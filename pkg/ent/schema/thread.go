package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Thread holds the schema definition for the Thread entity.
type Thread struct {
	ent.Schema
}

// Fields of the Thread.
func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name"),
		field.Uint64("description_id"),
		field.Uint64("subcategory_id"),
	}
}

// Edges of the Thread.
func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subcategory", Subcategory.Type).Ref("threads").Unique().Required().Field("subcategory_id"),
		edge.From("description", Post.Type).Ref("described_thread").Unique().Required().Field("description_id"),
		edge.To("posts", Post.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
