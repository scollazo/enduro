package package_

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/jmoiron/sqlx"
	temporalsdk_client "go.temporal.io/sdk/client"

	goapackage "github.com/artefactual-labs/enduro/internal/api/gen/package_"
)

type Service interface {
	// Goa returns an implementation of the goapackage Service.
	Goa() goapackage.Service
	Create(context.Context, *Package) error
	UpdateWorkflowStatus(ctx context.Context, ID uint, name string, workflowID, runID, aipID string, status Status, storedAt time.Time) error
	SetStatus(ctx context.Context, ID uint, status Status) error
	SetStatusInProgress(ctx context.Context, ID uint, startedAt time.Time) error
	SetStatusPending(ctx context.Context, ID uint) error
	SetLocation(ctx context.Context, ID uint, location string) error
	CreatePreservationAction(ctx context.Context, pa *PreservationAction) error
}

type packageImpl struct {
	logger logr.Logger
	db     *sqlx.DB
	tc     temporalsdk_client.Client

	// Destination for events to be published.
	events EventService
}

var _ Service = (*packageImpl)(nil)

func NewService(logger logr.Logger, db *sql.DB, tc temporalsdk_client.Client) *packageImpl {
	return &packageImpl{
		logger: logger,
		db:     sqlx.NewDb(db, "mysql"),
		tc:     tc,
		events: NewEventService(),
	}
}

func (svc *packageImpl) Goa() goapackage.Service {
	return &goaWrapper{
		packageImpl: svc,
	}
}

func (svc *packageImpl) Create(ctx context.Context, pkg *Package) error {
	query := `INSERT INTO package (name, workflow_id, run_id, aip_id, location, status) VALUES ((?), (?), (?), (?), (?), (?))`
	args := []interface{}{
		pkg.Name,
		pkg.WorkflowID,
		pkg.RunID,
		pkg.AIPID,
		pkg.Location,
		pkg.Status,
	}

	query = svc.db.Rebind(query)
	res, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting package: %w", err)
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return fmt.Errorf("error retrieving insert ID: %w", err)
	}

	pkg.ID = uint(id)

	if pkg, err := svc.Goa().Show(ctx, &goapackage.ShowPayload{ID: uint(id)}); err == nil {
		event := &goapackage.EnduroPackageCreatedEvent{ID: uint(id), Item: pkg}
		publishEvent(ctx, svc.events, event)
	}

	return nil
}

func (svc *packageImpl) UpdateWorkflowStatus(ctx context.Context, ID uint, name string, workflowID, runID, aipID string, status Status, storedAt time.Time) error {
	// Ensure that storedAt is reset during retries.
	completedAt := &storedAt
	if status == StatusInProgress {
		completedAt = nil
	}
	if completedAt != nil && completedAt.IsZero() {
		completedAt = nil
	}

	query := `UPDATE package SET name = (?), workflow_id = (?), run_id = (?), aip_id = (?), status = (?), completed_at = (?) WHERE id = (?)`
	args := []interface{}{
		name,
		workflowID,
		runID,
		aipID,
		status,
		completedAt,
		ID,
	}

	if _, err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	if pkg, err := svc.Goa().Show(ctx, &goapackage.ShowPayload{ID: ID}); err == nil {
		event := &goapackage.EnduroPackageUpdatedEvent{ID: uint(ID), Item: pkg}
		publishEvent(ctx, svc.events, event)
	}

	return nil
}

func (svc *packageImpl) SetStatus(ctx context.Context, ID uint, status Status) error {
	query := `UPDATE package SET status = (?) WHERE id = (?)`
	args := []interface{}{
		status,
		ID,
	}

	if _, err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	event := &goapackage.EnduroPackageStatusUpdatedEvent{ID: uint(ID), Status: status.String()}
	publishEvent(ctx, svc.events, event)

	return nil
}

func (svc *packageImpl) SetStatusInProgress(ctx context.Context, ID uint, startedAt time.Time) error {
	var query string
	args := []interface{}{StatusInProgress}

	if !startedAt.IsZero() {
		query = `UPDATE package SET status = (?), started_at = (?) WHERE id = (?)`
		args = append(args, startedAt, ID)
	} else {
		query = `UPDATE package SET status = (?) WHERE id = (?)`
		args = append(args, ID)
	}

	if _, err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	event := &goapackage.EnduroPackageStatusUpdatedEvent{ID: uint(ID), Status: StatusInProgress.String()}
	publishEvent(ctx, svc.events, event)

	return nil
}

func (svc *packageImpl) SetStatusPending(ctx context.Context, ID uint) error {
	query := `UPDATE package SET status = (?), WHERE id = (?)`
	args := []interface{}{
		StatusPending,
		ID,
	}

	if _, err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	event := &goapackage.EnduroPackageStatusUpdatedEvent{ID: uint(ID), Status: StatusPending.String()}
	publishEvent(ctx, svc.events, event)

	return nil
}

func (svc *packageImpl) SetLocation(ctx context.Context, ID uint, location string) error {
	query := `UPDATE package SET location = (?) WHERE id = (?)`
	args := []interface{}{
		location,
		ID,
	}

	if _, err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	event := &goapackage.EnduroPackageLocationUpdatedEvent{ID: uint(ID), Location: location}
	publishEvent(ctx, svc.events, event)

	return nil
}

func (svc *packageImpl) updateRow(ctx context.Context, query string, args []interface{}) (int64, error) {
	query = svc.db.Rebind(query)
	res, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("error updating package: %v", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error retrieving rows affected: %v", err)
	}

	return n, nil
}

func (svc *packageImpl) read(ctx context.Context, ID uint) (*Package, error) {
	query := "SELECT id, name, workflow_id, run_id, aip_id, location, status, CONVERT_TZ(created_at, @@session.time_zone, '+00:00') AS created_at, CONVERT_TZ(started_at, @@session.time_zone, '+00:00') AS started_at, CONVERT_TZ(completed_at, @@session.time_zone, '+00:00') AS completed_at FROM package WHERE id = (?)"
	args := []interface{}{ID}
	c := Package{}

	query = svc.db.Rebind(query)
	if err := svc.db.GetContext(ctx, &c, query, args...); err != nil {
		return nil, err
	}

	return &c, nil
}
