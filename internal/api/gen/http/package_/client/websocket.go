// Code generated by goa v3.7.7, DO NOT EDIT.
//
// package WebSocket client streaming
//
// Command:
// $ goa-v3.7.7 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"io"

	package_ "github.com/artefactual-labs/enduro/internal/api/gen/package_"
	package_views "github.com/artefactual-labs/enduro/internal/api/gen/package_/views"
	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "package" service.
type ConnConfigurer struct {
	MonitorFn goahttp.ConnConfigureFunc
}

// MonitorClientStream implements the package_.MonitorClientStream interface.
type MonitorClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "package" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		MonitorFn: fn,
	}
}

// Recv reads instances of "package_.EnduroMonitorEvent" from the "monitor"
// endpoint websocket connection.
func (s *MonitorClientStream) Recv() (*package_.EnduroMonitorEvent, error) {
	var (
		rv   *package_.EnduroMonitorEvent
		body MonitorResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewMonitorEnduroMonitorEventOK(&body)
	vres := &package_views.EnduroMonitorEvent{res, "default"}
	if err := package_views.ValidateEnduroMonitorEvent(vres); err != nil {
		return rv, goahttp.ErrValidationError("package", "monitor", err)
	}
	return package_.NewEnduroMonitorEvent(vres), nil
}
