package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subcategory holds the schema definition for the Subcategory entity.
type Subcategory struct {
	ent.Schema
}

// Fields of the Subcategory.
func (Subcategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name"),
		field.String("description"),
		field.Uint64("category_id"),
	}
}

// Edges of the Subcategory.
func (Subcategory) Edges() []ent.Edge {
	// add edge LinkedSubcategories for creating chains of subcategories
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("subcategories").Unique().Required().Field("category_id"),
		edge.To("threads", Thread.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
