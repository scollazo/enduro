// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AipColumns holds the columns for the "aip" table.
	AipColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "aip_id", Type: field.TypeUUID, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unspecified", "in_review", "rejected", "stored", "moving"}},
		{Name: "object_key", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "location_id", Type: field.TypeInt, Nullable: true},
	}
	// AipTable holds the schema information for the "aip" table.
	AipTable = &schema.Table{
		Name:       "aip",
		Columns:    AipColumns,
		PrimaryKey: []*schema.Column{AipColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "aip_location_location",
				Columns:    []*schema.Column{AipColumns[6]},
				RefColumns: []*schema.Column{LocationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "aip_aip_id",
				Unique:  false,
				Columns: []*schema.Column{AipColumns[2]},
			},
			{
				Name:    "aip_object_key",
				Unique:  false,
				Columns: []*schema.Column{AipColumns[4]},
			},
		},
	}
	// LocationColumns holds the columns for the "location" table.
	LocationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "description", Type: field.TypeString, Size: 2048},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"unspecified", "minio", "sftp", "amss"}},
		{Name: "purpose", Type: field.TypeEnum, Enums: []string{"unspecified", "aip_store"}},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "config", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
	}
	// LocationTable holds the schema information for the "location" table.
	LocationTable = &schema.Table{
		Name:       "location",
		Columns:    LocationColumns,
		PrimaryKey: []*schema.Column{LocationColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "location_name",
				Unique:  false,
				Columns: []*schema.Column{LocationColumns[1]},
				Annotation: &entsql.IndexAnnotation{
					Prefix: 50,
				},
			},
			{
				Name:    "location_uuid",
				Unique:  false,
				Columns: []*schema.Column{LocationColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AipTable,
		LocationTable,
	}
)

func init() {
	AipTable.ForeignKeys[0].RefTable = LocationTable
	AipTable.Annotation = &entsql.Annotation{
		Table: "aip",
	}
	LocationTable.Annotation = &entsql.Annotation{
		Table: "location",
	}
}
