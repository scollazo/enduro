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
	// TaskColumns holds the columns for the "task" table.
	TaskColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unspecified", "in progress", "done", "error", "queued", "pending"}},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "note", Type: field.TypeString, Size: 2147483647},
		{Name: "workflow_id", Type: field.TypeInt},
	}
	// TaskTable holds the schema information for the "task" table.
	TaskTable = &schema.Table{
		Name:       "task",
		Columns:    TaskColumns,
		PrimaryKey: []*schema.Column{TaskColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_workflow_tasks",
				Columns:    []*schema.Column{TaskColumns[7]},
				RefColumns: []*schema.Column{WorkflowColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "task_uuid",
				Unique:  false,
				Columns: []*schema.Column{TaskColumns[1]},
			},
		},
	}
	// WorkflowColumns holds the columns for the "workflow" table.
	WorkflowColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "temporal_id", Type: field.TypeString, Size: 255},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"unspecified", "upload aip", "move aip", "delete aip"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unspecified", "in progress", "done", "error", "queued", "pending"}},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "aip_id", Type: field.TypeInt},
	}
	// WorkflowTable holds the schema information for the "workflow" table.
	WorkflowTable = &schema.Table{
		Name:       "workflow",
		Columns:    WorkflowColumns,
		PrimaryKey: []*schema.Column{WorkflowColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_aip_workflows",
				Columns:    []*schema.Column{WorkflowColumns[7]},
				RefColumns: []*schema.Column{AipColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workflow_uuid",
				Unique:  false,
				Columns: []*schema.Column{WorkflowColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AipTable,
		LocationTable,
		TaskTable,
		WorkflowTable,
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
	TaskTable.ForeignKeys[0].RefTable = WorkflowTable
	TaskTable.Annotation = &entsql.Annotation{
		Table: "task",
	}
	WorkflowTable.ForeignKeys[0].RefTable = AipTable
	WorkflowTable.Annotation = &entsql.Annotation{
		Table: "workflow",
	}
}
