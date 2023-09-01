package entclient

import (
	"github.com/go-logr/logr"

	"github.com/artefactual-sdps/enduro/internal/persistence"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db"
)

type client struct {
	logger logr.Logger
	ent    *db.Client
}

var _ persistence.Service = (*client)(nil)

func New(logger logr.Logger, ent *db.Client) persistence.Service {
	return &client{
		logger: logger,
		ent:    ent,
	}
}
