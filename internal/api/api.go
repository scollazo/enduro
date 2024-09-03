/*
Package api contains the API server.

HTTP is the only transport supported at the moment.

The design package is the Goa design package while the gen package contains all
the generated code produced with goa gen.
*/
package api

import (
	"context"
	"embed"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-logr/logr"
	"github.com/gorilla/websocket"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	goahttp "goa.design/goa/v3/http"
	goahttpmwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	packagesvr "github.com/artefactual-sdps/enduro/internal/api/gen/http/package_/server"
	storagesvr "github.com/artefactual-sdps/enduro/internal/api/gen/http/storage/server"
	swaggersvr "github.com/artefactual-sdps/enduro/internal/api/gen/http/swagger/server"
	"github.com/artefactual-sdps/enduro/internal/api/gen/package_"
	"github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	intpkg "github.com/artefactual-sdps/enduro/internal/package_"
	intstorage "github.com/artefactual-sdps/enduro/internal/storage"
	"github.com/artefactual-sdps/enduro/internal/version"
)

//go:embed gen/http/openapi.json
var openAPIJSON embed.FS

func HTTPServer(
	logger logr.Logger,
	tp trace.TracerProvider,
	config *Config,
	pkgsvc intpkg.Service,
	storagesvc intstorage.Service,
) *http.Server {
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	websocketUpgrader := &websocket.Upgrader{
		HandshakeTimeout: time.Second,
		CheckOrigin:      sameOriginChecker(logger),
	}

	// Package service.
	packageEndpoints := package_.NewEndpoints(pkgsvc.Goa())
	packageErrorHandler := errorHandler(logger, "Package error.")
	packageServer := packagesvr.New(packageEndpoints, mux, dec, enc, packageErrorHandler, nil, websocketUpgrader, nil)
	packagesvr.Mount(mux, packageServer)

	// Storage service.
	storageEndpoints := storage.NewEndpoints(storagesvc)
	storageErrorHandler := errorHandler(logger, "Storage error.")
	storageServer := storagesvr.New(storageEndpoints, mux, dec, enc, storageErrorHandler, nil)
	storageServer.Download = writeTimeout(intstorage.Download(storagesvc, mux, dec), 0)
	storagesvr.Mount(mux, storageServer)

	// Swagger service.
	swaggerService := swaggersvr.New(nil, nil, nil, nil, nil, nil, http.FS(openAPIJSON))
	swaggersvr.Mount(mux, swaggerService)

	// Global middlewares.
	var handler http.Handler = mux
	handler = recoverMiddleware(logger)(handler)
	handler = otelhttp.NewHandler(handler, "api", otelhttp.WithTracerProvider(tp))
	handler = versionHeaderMiddleware(version.Short)(handler)
	if config.Debug {
		handler = goahttpmwr.Log(loggerAdapter(logger))(handler) //nolint SA1019: deprecated - use OpenTelemetry.
		handler = goahttpmwr.Debug(mux, os.Stdout)(handler)
	}

	return &http.Server{
		Addr:         config.Listen,
		Handler:      handler,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 120,
	}
}

type errorMessage struct {
	RequestID string
	Error     error
}

// errorHandler returns a function that writes and logs the given error
// including the request ID.
func errorHandler(logger logr.Logger, msg string) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		reqID, ok := ctx.Value(middleware.RequestIDKey).(string)
		if !ok {
			reqID = "unknown"
		}

		// Only write the error if the connection is not hijacked.
		var ws bool
		if _, err := w.Write(nil); err == http.ErrHijacked {
			ws = true
		} else {
			_ = json.NewEncoder(w).Encode(&errorMessage{RequestID: reqID})
		}

		logger.Error(err, "Service error.", "reqID", reqID, "ws", ws, "msg", msg)
	}
}
