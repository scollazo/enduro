// Code generated by goa v3.11.3, DO NOT EDIT.
//
// upload client HTTP transport
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the upload service endpoint HTTP clients.
type Client struct {
	// Upload Doer is the HTTP client used to make requests to the upload endpoint.
	UploadDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the upload service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		UploadDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Upload returns an endpoint that makes HTTP requests to the upload service
// upload server.
func (c *Client) Upload() goa.Endpoint {
	var (
		encodeRequest  = EncodeUploadRequest(c.encoder)
		decodeResponse = DecodeUploadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUploadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UploadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("upload", "upload", err)
		}
		return decodeResponse(resp)
	}
}
