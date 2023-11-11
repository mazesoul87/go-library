// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TestsColumns holds the columns for the "tests" table.
	TestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TestsTable holds the schema information for the "tests" table.
	TestsTable = &schema.Table{
		Name:       "tests",
		Columns:    TestsColumns,
		PrimaryKey: []*schema.Column{TestsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TestsTable,
	}
)

func init() {
}
