// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DataListsColumns holds the columns for the "data_lists" table.
	DataListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "label", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString, Size: 2147483647},
		{Name: "item_order", Type: field.TypeInt, Default: 1},
	}
	// DataListsTable holds the schema information for the "data_lists" table.
	DataListsTable = &schema.Table{
		Name:       "data_lists",
		Columns:    DataListsColumns,
		PrimaryKey: []*schema.Column{DataListsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "datalist_key_kind",
				Unique:  true,
				Columns: []*schema.Column{DataListsColumns[5], DataListsColumns[4]},
			},
		},
	}
	// SiteConfigsColumns holds the columns for the "site_configs" table.
	SiteConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "config", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// SiteConfigsTable holds the schema information for the "site_configs" table.
	SiteConfigsTable = &schema.Table{
		Name:       "site_configs",
		Columns:    SiteConfigsColumns,
		PrimaryKey: []*schema.Column{SiteConfigsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"user", "admin"}, Default: "user"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DataListsTable,
		SiteConfigsTable,
		UsersTable,
	}
)

func init() {
}
