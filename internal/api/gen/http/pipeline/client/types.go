// Code generated by goa v3.2.2, DO NOT EDIT.
//
// pipeline HTTP client types
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	pipeline "github.com/artefactual-labs/enduro/internal/api/gen/pipeline"
	pipelineviews "github.com/artefactual-labs/enduro/internal/api/gen/pipeline/views"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "pipeline" service "list" endpoint HTTP
// response body.
type ListResponseBody []*EnduroStoredPipelineResponse

// ShowResponseBody is the type of the "pipeline" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// Name of the collection
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the collection
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "pipeline" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Identifier of missing collection
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// EnduroStoredPipelineResponse is used to define fields on response body types.
type EnduroStoredPipelineResponse struct {
	// Name of the collection
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the collection
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewListEnduroStoredPipelineOK builds a "pipeline" service "list" endpoint
// result from a HTTP "OK" response.
func NewListEnduroStoredPipelineOK(body []*EnduroStoredPipelineResponse) []*pipeline.EnduroStoredPipeline {
	v := make([]*pipeline.EnduroStoredPipeline, len(body))
	for i, val := range body {
		v[i] = unmarshalEnduroStoredPipelineResponseToPipelineEnduroStoredPipeline(val)
	}
	return v
}

// NewShowEnduroStoredPipelineOK builds a "pipeline" service "show" endpoint
// result from a HTTP "OK" response.
func NewShowEnduroStoredPipelineOK(body *ShowResponseBody) *pipelineviews.EnduroStoredPipelineView {
	v := &pipelineviews.EnduroStoredPipelineView{
		ID:   body.ID,
		Name: body.Name,
	}

	return v
}

// NewShowNotFound builds a pipeline service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *pipeline.NotFound {
	v := &pipeline.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateEnduroStoredPipelineResponse runs the validations defined on
// EnduroStored-PipelineResponse
func ValidateEnduroStoredPipelineResponse(body *EnduroStoredPipelineResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}
