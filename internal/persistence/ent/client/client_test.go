package entclient_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"gotest.tools/v3/assert"

	"github.com/artefactual-sdps/enduro/internal/enums"
	"github.com/artefactual-sdps/enduro/internal/persistence"
	entclient "github.com/artefactual-sdps/enduro/internal/persistence/ent/client"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/enttest"
)

func setUpClient(t *testing.T, logger logr.Logger) (*db.Client, persistence.Service) {
	t.Helper()

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name())
	entc := enttest.Open(t, "sqlite3", dsn)
	t.Cleanup(func() { entc.Close() })

	c := entclient.New(logger, entc)

	return entc, c
}

func createSIP(
	entc *db.Client,
	name string,
	status enums.SIPStatus,
) (*db.SIP, error) {
	runID := uuid.MustParse("aee9644d-6397-4b34-92f7-442ad3dd3b13")
	aipID := uuid.MustParse("30223842-0650-4f79-80bd-7bf43b810656")

	return entc.SIP.Create().
		SetName(name).
		SetWorkflowID("12345").
		SetRunID(runID).
		SetAipID(aipID).
		SetStatus(int8(status)). // #nosec G115 -- constrained value.
		Save(context.Background())
}

func createPreservationAction(
	entc *db.Client,
	sipID int,
	status enums.PreservationActionStatus,
) (*db.PreservationAction, error) {
	return entc.PreservationAction.Create().
		SetWorkflowID("12345").
		SetType(int8(enums.PreservationActionTypeCreateAip)).
		SetStatus(int8(status)). // #nosec G115 -- constrained value.
		SetSipID(sipID).
		Save(context.Background())
}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("Returns a working ent DB client", func(t *testing.T) {
		t.Parallel()

		entc, _ := setUpClient(t, logr.Discard())
		s, err := createSIP(
			entc,
			"testing 1-2-3",
			enums.SIPStatusInProgress,
		)
		assert.NilError(t, err)

		assert.Equal(t, s.Name, "testing 1-2-3")
		assert.Equal(t, s.WorkflowID, "12345")
		assert.Equal(t, s.RunID, uuid.MustParse("aee9644d-6397-4b34-92f7-442ad3dd3b13"))
		assert.Equal(t, s.AipID, uuid.MustParse("30223842-0650-4f79-80bd-7bf43b810656"))
		assert.Equal(t, s.Status, int8(enums.SIPStatusInProgress))
	})
}
