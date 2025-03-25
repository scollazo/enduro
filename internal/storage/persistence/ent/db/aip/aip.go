// Code generated by ent, DO NOT EDIT.

package aip

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artefactual-sdps/enduro/internal/storage/enums"
)

const (
	// Label holds the string label denoting the aip type in the database.
	Label = "aip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAipID holds the string denoting the aip_id field in the database.
	FieldAipID = "aip_id"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldObjectKey holds the string denoting the object_key field in the database.
	FieldObjectKey = "object_key"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeWorkflows holds the string denoting the workflows edge name in mutations.
	EdgeWorkflows = "workflows"
	// Table holds the table name of the aip in the database.
	Table = "aip"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "aip"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "location"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location_id"
	// WorkflowsTable is the table that holds the workflows relation/edge.
	WorkflowsTable = "workflow"
	// WorkflowsInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowsInverseTable = "workflow"
	// WorkflowsColumn is the table column denoting the workflows relation/edge.
	WorkflowsColumn = "aip_id"
)

// Columns holds all SQL columns for aip fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAipID,
	FieldLocationID,
	FieldStatus,
	FieldObjectKey,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.AIPStatus) error {
	switch s.String() {
	case "unspecified", "in_review", "rejected", "stored", "moving":
		return nil
	default:
		return fmt.Errorf("aip: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the AIP queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAipID orders the results by the aip_id field.
func ByAipID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAipID, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByObjectKey orders the results by the object_key field.
func ByObjectKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldObjectKey, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLocationField orders the results by location field.
func ByLocationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkflowsCount orders the results by workflows count.
func ByWorkflowsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowsStep(), opts...)
	}
}

// ByWorkflows orders the results by workflows terms.
func ByWorkflows(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLocationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, LocationTable, LocationColumn),
	)
}
func newWorkflowsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowsTable, WorkflowsColumn),
	)
}
