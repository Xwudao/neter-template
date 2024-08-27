package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// DataList holds the schema definition for the DataList entity.
type DataList struct {
	ent.Schema
}

// Fields of the DataList.
func (DataList) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("label").NotEmpty().Comment("标签"),
		field.String("kind").NotEmpty().Comment("分类"),
		field.String("key").Comment("和kind组合成唯一索引"),
		field.Text("value").NotEmpty().Comment("内容，一般为json格式"),

		field.Int("item_order").Default(1).Min(1).Max(1000).Comment("排序"),
	}
}

// Edges of the DataList.
func (DataList) Edges() []ent.Edge {
	return nil
}

// Mixin of the DataList.
func (DataList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the DataList.
func (DataList) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key", "kind").Unique(),
	}
}
