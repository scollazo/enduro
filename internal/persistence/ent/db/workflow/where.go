// Code generated by ent, DO NOT EDIT.

package workflow

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldID, id))
}

// WorkflowID applies equality check predicate on the "workflow_id" field. It's identical to WorkflowIDEQ.
func WorkflowID(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldWorkflowID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldType, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldStatus, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldStartedAt, v))
}

// CompletedAt applies equality check predicate on the "completed_at" field. It's identical to CompletedAtEQ.
func CompletedAt(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldCompletedAt, v))
}

// SipID applies equality check predicate on the "sip_id" field. It's identical to SipIDEQ.
func SipID(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldSipID, v))
}

// WorkflowIDEQ applies the EQ predicate on the "workflow_id" field.
func WorkflowIDEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldWorkflowID, v))
}

// WorkflowIDNEQ applies the NEQ predicate on the "workflow_id" field.
func WorkflowIDNEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldWorkflowID, v))
}

// WorkflowIDIn applies the In predicate on the "workflow_id" field.
func WorkflowIDIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldWorkflowID, vs...))
}

// WorkflowIDNotIn applies the NotIn predicate on the "workflow_id" field.
func WorkflowIDNotIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldWorkflowID, vs...))
}

// WorkflowIDGT applies the GT predicate on the "workflow_id" field.
func WorkflowIDGT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldWorkflowID, v))
}

// WorkflowIDGTE applies the GTE predicate on the "workflow_id" field.
func WorkflowIDGTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldWorkflowID, v))
}

// WorkflowIDLT applies the LT predicate on the "workflow_id" field.
func WorkflowIDLT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldWorkflowID, v))
}

// WorkflowIDLTE applies the LTE predicate on the "workflow_id" field.
func WorkflowIDLTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldWorkflowID, v))
}

// WorkflowIDContains applies the Contains predicate on the "workflow_id" field.
func WorkflowIDContains(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContains(FieldWorkflowID, v))
}

// WorkflowIDHasPrefix applies the HasPrefix predicate on the "workflow_id" field.
func WorkflowIDHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasPrefix(FieldWorkflowID, v))
}

// WorkflowIDHasSuffix applies the HasSuffix predicate on the "workflow_id" field.
func WorkflowIDHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasSuffix(FieldWorkflowID, v))
}

// WorkflowIDEqualFold applies the EqualFold predicate on the "workflow_id" field.
func WorkflowIDEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEqualFold(FieldWorkflowID, v))
}

// WorkflowIDContainsFold applies the ContainsFold predicate on the "workflow_id" field.
func WorkflowIDContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContainsFold(FieldWorkflowID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldStatus, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldStartedAt, v))
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldIsNull(FieldStartedAt))
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldNotNull(FieldStartedAt))
}

// CompletedAtEQ applies the EQ predicate on the "completed_at" field.
func CompletedAtEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldCompletedAt, v))
}

// CompletedAtNEQ applies the NEQ predicate on the "completed_at" field.
func CompletedAtNEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldCompletedAt, v))
}

// CompletedAtIn applies the In predicate on the "completed_at" field.
func CompletedAtIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldCompletedAt, vs...))
}

// CompletedAtNotIn applies the NotIn predicate on the "completed_at" field.
func CompletedAtNotIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldCompletedAt, vs...))
}

// CompletedAtGT applies the GT predicate on the "completed_at" field.
func CompletedAtGT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldCompletedAt, v))
}

// CompletedAtGTE applies the GTE predicate on the "completed_at" field.
func CompletedAtGTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldCompletedAt, v))
}

// CompletedAtLT applies the LT predicate on the "completed_at" field.
func CompletedAtLT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldCompletedAt, v))
}

// CompletedAtLTE applies the LTE predicate on the "completed_at" field.
func CompletedAtLTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldCompletedAt, v))
}

// CompletedAtIsNil applies the IsNil predicate on the "completed_at" field.
func CompletedAtIsNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldIsNull(FieldCompletedAt))
}

// CompletedAtNotNil applies the NotNil predicate on the "completed_at" field.
func CompletedAtNotNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldNotNull(FieldCompletedAt))
}

// SipIDEQ applies the EQ predicate on the "sip_id" field.
func SipIDEQ(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldSipID, v))
}

// SipIDNEQ applies the NEQ predicate on the "sip_id" field.
func SipIDNEQ(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldSipID, v))
}

// SipIDIn applies the In predicate on the "sip_id" field.
func SipIDIn(vs ...int) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldSipID, vs...))
}

// SipIDNotIn applies the NotIn predicate on the "sip_id" field.
func SipIDNotIn(vs ...int) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldSipID, vs...))
}

// HasSip applies the HasEdge predicate on the "sip" edge.
func HasSip() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SipTable, SipColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSipWith applies the HasEdge predicate on the "sip" edge with a given conditions (other predicates).
func HasSipWith(preds ...predicate.SIP) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := newSipStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := newTasksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(sql.NotPredicates(p))
}
