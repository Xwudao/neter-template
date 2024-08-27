package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SiteConfig holds the schema definition for the SiteConfig entity.
type SiteConfig struct {
	ent.Schema
}

// Fields of the SiteConfig.
func (SiteConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty().Unique(),
		field.Text("config").Optional(),
	}
}

// Edges of the SiteConfig.
func (SiteConfig) Edges() []ent.Edge {
	return nil
}

// Mixin of the SiteConfig.
func (SiteConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
