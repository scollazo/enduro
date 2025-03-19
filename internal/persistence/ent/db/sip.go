// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/sip"
	"github.com/google/uuid"
)

// SIP is the model entity for the SIP schema.
type SIP struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// WorkflowID holds the value of the "workflow_id" field.
	WorkflowID string `json:"workflow_id,omitempty"`
	// RunID holds the value of the "run_id" field.
	RunID uuid.UUID `json:"run_id,omitempty"`
	// AipID holds the value of the "aip_id" field.
	AipID uuid.UUID `json:"aip_id,omitempty"`
	// LocationID holds the value of the "location_id" field.
	LocationID uuid.UUID `json:"location_id,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt time.Time `json:"started_at,omitempty"`
	// CompletedAt holds the value of the "completed_at" field.
	CompletedAt time.Time `json:"completed_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SIPQuery when eager-loading is set.
	Edges        SIPEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SIPEdges holds the relations/edges for other nodes in the graph.
type SIPEdges struct {
	// Workflows holds the value of the workflows edge.
	Workflows []*Workflow `json:"workflows,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkflowsOrErr returns the Workflows value or an error if the edge
// was not loaded in eager-loading.
func (e SIPEdges) WorkflowsOrErr() ([]*Workflow, error) {
	if e.loadedTypes[0] {
		return e.Workflows, nil
	}
	return nil, &NotLoadedError{edge: "workflows"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SIP) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sip.FieldID, sip.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sip.FieldName, sip.FieldWorkflowID:
			values[i] = new(sql.NullString)
		case sip.FieldCreatedAt, sip.FieldStartedAt, sip.FieldCompletedAt:
			values[i] = new(sql.NullTime)
		case sip.FieldRunID, sip.FieldAipID, sip.FieldLocationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SIP fields.
func (s *SIP) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sip.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sip.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case sip.FieldWorkflowID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_id", values[i])
			} else if value.Valid {
				s.WorkflowID = value.String
			}
		case sip.FieldRunID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field run_id", values[i])
			} else if value != nil {
				s.RunID = *value
			}
		case sip.FieldAipID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field aip_id", values[i])
			} else if value != nil {
				s.AipID = *value
			}
		case sip.FieldLocationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value != nil {
				s.LocationID = *value
			}
		case sip.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = int8(value.Int64)
			}
		case sip.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sip.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				s.StartedAt = value.Time
			}
		case sip.FieldCompletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field completed_at", values[i])
			} else if value.Valid {
				s.CompletedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SIP.
// This includes values selected through modifiers, order, etc.
func (s *SIP) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryWorkflows queries the "workflows" edge of the SIP entity.
func (s *SIP) QueryWorkflows() *WorkflowQuery {
	return NewSIPClient(s.config).QueryWorkflows(s)
}

// Update returns a builder for updating this SIP.
// Note that you need to call SIP.Unwrap() before calling this method if this SIP
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *SIP) Update() *SIPUpdateOne {
	return NewSIPClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the SIP entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *SIP) Unwrap() *SIP {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("db: SIP is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *SIP) String() string {
	var builder strings.Builder
	builder.WriteString("SIP(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("workflow_id=")
	builder.WriteString(s.WorkflowID)
	builder.WriteString(", ")
	builder.WriteString("run_id=")
	builder.WriteString(fmt.Sprintf("%v", s.RunID))
	builder.WriteString(", ")
	builder.WriteString("aip_id=")
	builder.WriteString(fmt.Sprintf("%v", s.AipID))
	builder.WriteString(", ")
	builder.WriteString("location_id=")
	builder.WriteString(fmt.Sprintf("%v", s.LocationID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(s.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("completed_at=")
	builder.WriteString(s.CompletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SIPs is a parsable slice of SIP.
type SIPs []*SIP
