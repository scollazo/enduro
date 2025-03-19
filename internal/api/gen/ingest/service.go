// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ingest service
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package ingest

import (
	"context"
	"io"

	ingestviews "github.com/artefactual-sdps/enduro/internal/api/gen/ingest/views"
	"github.com/google/uuid"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The ingest service manages ingested SIPs.
type Service interface {
	// Request access to the /monitor WebSocket
	MonitorRequest(context.Context, *MonitorRequestPayload) (res *MonitorRequestResult, err error)
	// Obtain access to the /monitor WebSocket
	Monitor(context.Context, *MonitorPayload, MonitorServerStream) (err error)
	// List all ingested SIPs
	ListSips(context.Context, *ListSipsPayload) (res *SIPs, err error)
	// Show SIP by ID
	ShowSip(context.Context, *ShowSipPayload) (res *SIP, err error)
	// List all workflows for a SIP
	ListSipWorkflows(context.Context, *ListSipWorkflowsPayload) (res *SIPWorkflows, err error)
	// Signal the SIP has been reviewed and accepted
	ConfirmSip(context.Context, *ConfirmSipPayload) (err error)
	// Signal the SIP has been reviewed and rejected
	RejectSip(context.Context, *RejectSipPayload) (err error)
	// Move a SIP to a permanent storage location
	MoveSip(context.Context, *MoveSipPayload) (err error)
	// Retrieve the status of a permanent storage location move of the SIP
	MoveSipStatus(context.Context, *MoveSipStatusPayload) (res *MoveStatusResult, err error)
	// Upload a SIP to trigger an ingest workflow
	UploadSip(context.Context, *UploadSipPayload, io.ReadCloser) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "enduro"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "ingest"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [10]string{"monitor_request", "monitor", "list_sips", "show_sip", "list_sip_workflows", "confirm_sip", "reject_sip", "move_sip", "move_sip_status", "upload_sip"}

// MonitorServerStream is the interface a "monitor" endpoint server stream must
// satisfy.
type MonitorServerStream interface {
	// Send streams instances of "MonitorEvent".
	Send(*MonitorEvent) error
	// Close closes the stream.
	Close() error
}

// MonitorClientStream is the interface a "monitor" endpoint client stream must
// satisfy.
type MonitorClientStream interface {
	// Recv reads instances of "MonitorEvent" from the stream.
	Recv() (*MonitorEvent, error)
}

// ConfirmSipPayload is the payload type of the ingest service confirm_sip
// method.
type ConfirmSipPayload struct {
	// Identifier of SIP to look up
	ID uint
	// Identifier of storage location
	LocationID uuid.UUID
	Token      *string
}

// Page represents a subset of search results.
type EnduroPage struct {
	// Maximum items per page
	Limit int
	// Offset from first result to start of page
	Offset int
	// Total result count before paging
	Total int
}

// ListSipWorkflowsPayload is the payload type of the ingest service
// list_sip_workflows method.
type ListSipWorkflowsPayload struct {
	// Identifier of SIP to look up
	ID    uint
	Token *string
}

// ListSipsPayload is the payload type of the ingest service list_sips method.
type ListSipsPayload struct {
	Name *string
	// Identifier of AIP
	AipID               *string
	EarliestCreatedTime *string
	LatestCreatedTime   *string
	// Identifier of storage location
	LocationID *string
	Status     *string
	// Limit number of results to return
	Limit *int
	// Offset from the beginning of the found set
	Offset *int
	Token  *string
}

// MonitorEvent is the result type of the ingest service monitor method.
type MonitorEvent struct {
	Event interface {
		eventVal()
	}
}

// MonitorPayload is the payload type of the ingest service monitor method.
type MonitorPayload struct {
	Ticket *string
}

type MonitorPingEvent struct {
	Message *string
}

// MonitorRequestPayload is the payload type of the ingest service
// monitor_request method.
type MonitorRequestPayload struct {
	Token *string
}

// MonitorRequestResult is the result type of the ingest service
// monitor_request method.
type MonitorRequestResult struct {
	Ticket *string
}

// MoveSipPayload is the payload type of the ingest service move_sip method.
type MoveSipPayload struct {
	// Identifier of SIP to move
	ID uint
	// Identifier of storage location
	LocationID uuid.UUID
	Token      *string
}

// MoveSipStatusPayload is the payload type of the ingest service
// move_sip_status method.
type MoveSipStatusPayload struct {
	// Identifier of SIP to move
	ID    uint
	Token *string
}

// MoveStatusResult is the result type of the ingest service move_sip_status
// method.
type MoveStatusResult struct {
	Done bool
}

// RejectSipPayload is the payload type of the ingest service reject_sip method.
type RejectSipPayload struct {
	// Identifier of SIP to look up
	ID    uint
	Token *string
}

// SIP is the result type of the ingest service show_sip method.
type SIP struct {
	// Identifier of SIP
	ID uint
	// Name of the SIP
	Name *string
	// Identifier of storage location
	LocationID *uuid.UUID
	// Status of the SIP
	Status string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of AIP
	AipID *string
	// Creation datetime
	CreatedAt string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

type SIPCollection []*SIP

type SIPCreatedEvent struct {
	// Identifier of SIP
	ID   uint
	Item *SIP
}

type SIPLocationUpdatedEvent struct {
	// Identifier of SIP
	ID uint
	// Identifier of storage location
	LocationID uuid.UUID
}

// SIP not found.
type SIPNotFound struct {
	// Message of error
	Message string
	// Identifier of missing SIP
	ID uint
}

type SIPStatusUpdatedEvent struct {
	// Identifier of SIP
	ID     uint
	Status string
}

// SIPTask describes a SIP workflow task.
type SIPTask struct {
	ID          uint
	TaskID      string
	Name        string
	Status      string
	StartedAt   string
	CompletedAt *string
	Note        *string
	WorkflowID  *uint
}

type SIPTaskCollection []*SIPTask

type SIPTaskCreatedEvent struct {
	// Identifier of task
	ID   uint
	Item *SIPTask
}

type SIPTaskUpdatedEvent struct {
	// Identifier of task
	ID   uint
	Item *SIPTask
}

type SIPUpdatedEvent struct {
	// Identifier of SIP
	ID   uint
	Item *SIP
}

// SIPWorkflow describes a workflow of a SIP.
type SIPWorkflow struct {
	ID          uint
	WorkflowID  string
	Type        string
	Status      string
	StartedAt   string
	CompletedAt *string
	Tasks       SIPTaskCollection
	SipID       *uint
}

type SIPWorkflowCollection []*SIPWorkflow

type SIPWorkflowCreatedEvent struct {
	// Identifier of workflow
	ID   uint
	Item *SIPWorkflow
}

type SIPWorkflowUpdatedEvent struct {
	// Identifier of workflow
	ID   uint
	Item *SIPWorkflow
}

// SIPWorkflows is the result type of the ingest service list_sip_workflows
// method.
type SIPWorkflows struct {
	Workflows SIPWorkflowCollection
}

// SIPs is the result type of the ingest service list_sips method.
type SIPs struct {
	Items SIPCollection
	Page  *EnduroPage
}

// ShowSipPayload is the payload type of the ingest service show_sip method.
type ShowSipPayload struct {
	// Identifier of SIP to show
	ID    uint
	Token *string
}

// UploadSipPayload is the payload type of the ingest service upload_sip method.
type UploadSipPayload struct {
	// Content-Type header, must define value for multipart boundary.
	ContentType string
	Token       *string
}

// Forbidden
type Forbidden string

// Unauthorized
type Unauthorized string

// Error returns an error description.
func (e *SIPNotFound) Error() string {
	return "SIP not found."
}

// ErrorName returns "SIPNotFound".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *SIPNotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "SIPNotFound".
func (e *SIPNotFound) GoaErrorName() string {
	return "not_found"
}

// Error returns an error description.
func (e Forbidden) Error() string {
	return "Forbidden"
}

// ErrorName returns "forbidden".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e Forbidden) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "forbidden".
func (e Forbidden) GoaErrorName() string {
	return "forbidden"
}

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Unauthorized"
}

// ErrorName returns "unauthorized".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e Unauthorized) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "unauthorized".
func (e Unauthorized) GoaErrorName() string {
	return "unauthorized"
}
func (*MonitorPingEvent) eventVal()        {}
func (*SIPCreatedEvent) eventVal()         {}
func (*SIPLocationUpdatedEvent) eventVal() {}
func (*SIPStatusUpdatedEvent) eventVal()   {}
func (*SIPTaskCreatedEvent) eventVal()     {}
func (*SIPTaskUpdatedEvent) eventVal()     {}
func (*SIPUpdatedEvent) eventVal()         {}
func (*SIPWorkflowCreatedEvent) eventVal() {}
func (*SIPWorkflowUpdatedEvent) eventVal() {}

// MakeNotAvailable builds a goa.ServiceError from an error.
func MakeNotAvailable(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_available", false, false, false)
}

// MakeNotValid builds a goa.ServiceError from an error.
func MakeNotValid(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_valid", false, false, false)
}

// MakeFailedDependency builds a goa.ServiceError from an error.
func MakeFailedDependency(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "failed_dependency", false, false, false)
}

// MakeInvalidMediaType builds a goa.ServiceError from an error.
func MakeInvalidMediaType(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_media_type", false, false, false)
}

// MakeInvalidMultipartRequest builds a goa.ServiceError from an error.
func MakeInvalidMultipartRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_multipart_request", false, false, false)
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal_error", false, false, false)
}

// NewSIPs initializes result type SIPs from viewed result type SIPs.
func NewSIPs(vres *ingestviews.SIPs) *SIPs {
	return newSIPs(vres.Projected)
}

// NewViewedSIPs initializes viewed result type SIPs from result type SIPs
// using the given view.
func NewViewedSIPs(res *SIPs, view string) *ingestviews.SIPs {
	p := newSIPsView(res)
	return &ingestviews.SIPs{Projected: p, View: "default"}
}

// NewSIP initializes result type SIP from viewed result type SIP.
func NewSIP(vres *ingestviews.SIP) *SIP {
	return newSIP(vres.Projected)
}

// NewViewedSIP initializes viewed result type SIP from result type SIP using
// the given view.
func NewViewedSIP(res *SIP, view string) *ingestviews.SIP {
	p := newSIPView(res)
	return &ingestviews.SIP{Projected: p, View: "default"}
}

// NewSIPWorkflows initializes result type SIPWorkflows from viewed result type
// SIPWorkflows.
func NewSIPWorkflows(vres *ingestviews.SIPWorkflows) *SIPWorkflows {
	return newSIPWorkflows(vres.Projected)
}

// NewViewedSIPWorkflows initializes viewed result type SIPWorkflows from
// result type SIPWorkflows using the given view.
func NewViewedSIPWorkflows(res *SIPWorkflows, view string) *ingestviews.SIPWorkflows {
	p := newSIPWorkflowsView(res)
	return &ingestviews.SIPWorkflows{Projected: p, View: "default"}
}

// newSIPs converts projected type SIPs to service type SIPs.
func newSIPs(vres *ingestviews.SIPsView) *SIPs {
	res := &SIPs{}
	if vres.Items != nil {
		res.Items = newSIPCollection(vres.Items)
	}
	if vres.Page != nil {
		res.Page = newEnduroPage(vres.Page)
	}
	return res
}

// newSIPsView projects result type SIPs to projected type SIPsView using the
// "default" view.
func newSIPsView(res *SIPs) *ingestviews.SIPsView {
	vres := &ingestviews.SIPsView{}
	if res.Items != nil {
		vres.Items = newSIPCollectionView(res.Items)
	}
	if res.Page != nil {
		vres.Page = newEnduroPageView(res.Page)
	}
	return vres
}

// newSIPCollection converts projected type SIPCollection to service type
// SIPCollection.
func newSIPCollection(vres ingestviews.SIPCollectionView) SIPCollection {
	res := make(SIPCollection, len(vres))
	for i, n := range vres {
		res[i] = newSIP(n)
	}
	return res
}

// newSIPCollectionView projects result type SIPCollection to projected type
// SIPCollectionView using the "default" view.
func newSIPCollectionView(res SIPCollection) ingestviews.SIPCollectionView {
	vres := make(ingestviews.SIPCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSIPView(n)
	}
	return vres
}

// newSIP converts projected type SIP to service type SIP.
func newSIP(vres *ingestviews.SIPView) *SIP {
	res := &SIP{
		Name:        vres.Name,
		LocationID:  vres.LocationID,
		WorkflowID:  vres.WorkflowID,
		RunID:       vres.RunID,
		AipID:       vres.AipID,
		StartedAt:   vres.StartedAt,
		CompletedAt: vres.CompletedAt,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.Status == nil {
		res.Status = "new"
	}
	return res
}

// newSIPView projects result type SIP to projected type SIPView using the
// "default" view.
func newSIPView(res *SIP) *ingestviews.SIPView {
	vres := &ingestviews.SIPView{
		ID:          &res.ID,
		Name:        res.Name,
		LocationID:  res.LocationID,
		Status:      &res.Status,
		WorkflowID:  res.WorkflowID,
		RunID:       res.RunID,
		AipID:       res.AipID,
		CreatedAt:   &res.CreatedAt,
		StartedAt:   res.StartedAt,
		CompletedAt: res.CompletedAt,
	}
	return vres
}

// newEnduroPage converts projected type EnduroPage to service type EnduroPage.
func newEnduroPage(vres *ingestviews.EnduroPageView) *EnduroPage {
	res := &EnduroPage{}
	if vres.Limit != nil {
		res.Limit = *vres.Limit
	}
	if vres.Offset != nil {
		res.Offset = *vres.Offset
	}
	if vres.Total != nil {
		res.Total = *vres.Total
	}
	return res
}

// newEnduroPageView projects result type EnduroPage to projected type
// EnduroPageView using the "default" view.
func newEnduroPageView(res *EnduroPage) *ingestviews.EnduroPageView {
	vres := &ingestviews.EnduroPageView{
		Limit:  &res.Limit,
		Offset: &res.Offset,
		Total:  &res.Total,
	}
	return vres
}

// newSIPWorkflows converts projected type SIPWorkflows to service type
// SIPWorkflows.
func newSIPWorkflows(vres *ingestviews.SIPWorkflowsView) *SIPWorkflows {
	res := &SIPWorkflows{}
	if vres.Workflows != nil {
		res.Workflows = newSIPWorkflowCollection(vres.Workflows)
	}
	return res
}

// newSIPWorkflowsView projects result type SIPWorkflows to projected type
// SIPWorkflowsView using the "default" view.
func newSIPWorkflowsView(res *SIPWorkflows) *ingestviews.SIPWorkflowsView {
	vres := &ingestviews.SIPWorkflowsView{}
	if res.Workflows != nil {
		vres.Workflows = newSIPWorkflowCollectionView(res.Workflows)
	}
	return vres
}

// newSIPWorkflowCollectionSimple converts projected type SIPWorkflowCollection
// to service type SIPWorkflowCollection.
func newSIPWorkflowCollectionSimple(vres ingestviews.SIPWorkflowCollectionView) SIPWorkflowCollection {
	res := make(SIPWorkflowCollection, len(vres))
	for i, n := range vres {
		res[i] = newSIPWorkflowSimple(n)
	}
	return res
}

// newSIPWorkflowCollection converts projected type SIPWorkflowCollection to
// service type SIPWorkflowCollection.
func newSIPWorkflowCollection(vres ingestviews.SIPWorkflowCollectionView) SIPWorkflowCollection {
	res := make(SIPWorkflowCollection, len(vres))
	for i, n := range vres {
		res[i] = newSIPWorkflow(n)
	}
	return res
}

// newSIPWorkflowCollectionViewSimple projects result type
// SIPWorkflowCollection to projected type SIPWorkflowCollectionView using the
// "simple" view.
func newSIPWorkflowCollectionViewSimple(res SIPWorkflowCollection) ingestviews.SIPWorkflowCollectionView {
	vres := make(ingestviews.SIPWorkflowCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSIPWorkflowViewSimple(n)
	}
	return vres
}

// newSIPWorkflowCollectionView projects result type SIPWorkflowCollection to
// projected type SIPWorkflowCollectionView using the "default" view.
func newSIPWorkflowCollectionView(res SIPWorkflowCollection) ingestviews.SIPWorkflowCollectionView {
	vres := make(ingestviews.SIPWorkflowCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSIPWorkflowView(n)
	}
	return vres
}

// newSIPWorkflowSimple converts projected type SIPWorkflow to service type
// SIPWorkflow.
func newSIPWorkflowSimple(vres *ingestviews.SIPWorkflowView) *SIPWorkflow {
	res := &SIPWorkflow{
		CompletedAt: vres.CompletedAt,
		SipID:       vres.SipID,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.WorkflowID != nil {
		res.WorkflowID = *vres.WorkflowID
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.StartedAt != nil {
		res.StartedAt = *vres.StartedAt
	}
	if vres.Tasks != nil {
		res.Tasks = newSIPTaskCollection(vres.Tasks)
	}
	return res
}

// newSIPWorkflow converts projected type SIPWorkflow to service type
// SIPWorkflow.
func newSIPWorkflow(vres *ingestviews.SIPWorkflowView) *SIPWorkflow {
	res := &SIPWorkflow{
		CompletedAt: vres.CompletedAt,
		SipID:       vres.SipID,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.WorkflowID != nil {
		res.WorkflowID = *vres.WorkflowID
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.StartedAt != nil {
		res.StartedAt = *vres.StartedAt
	}
	if vres.Tasks != nil {
		res.Tasks = newSIPTaskCollection(vres.Tasks)
	}
	return res
}

// newSIPWorkflowViewSimple projects result type SIPWorkflow to projected type
// SIPWorkflowView using the "simple" view.
func newSIPWorkflowViewSimple(res *SIPWorkflow) *ingestviews.SIPWorkflowView {
	vres := &ingestviews.SIPWorkflowView{
		ID:          &res.ID,
		WorkflowID:  &res.WorkflowID,
		Type:        &res.Type,
		Status:      &res.Status,
		StartedAt:   &res.StartedAt,
		CompletedAt: res.CompletedAt,
		SipID:       res.SipID,
	}
	return vres
}

// newSIPWorkflowView projects result type SIPWorkflow to projected type
// SIPWorkflowView using the "default" view.
func newSIPWorkflowView(res *SIPWorkflow) *ingestviews.SIPWorkflowView {
	vres := &ingestviews.SIPWorkflowView{
		ID:          &res.ID,
		WorkflowID:  &res.WorkflowID,
		Type:        &res.Type,
		Status:      &res.Status,
		StartedAt:   &res.StartedAt,
		CompletedAt: res.CompletedAt,
		SipID:       res.SipID,
	}
	if res.Tasks != nil {
		vres.Tasks = newSIPTaskCollectionView(res.Tasks)
	}
	return vres
}

// newSIPTaskCollection converts projected type SIPTaskCollection to service
// type SIPTaskCollection.
func newSIPTaskCollection(vres ingestviews.SIPTaskCollectionView) SIPTaskCollection {
	res := make(SIPTaskCollection, len(vres))
	for i, n := range vres {
		res[i] = newSIPTask(n)
	}
	return res
}

// newSIPTaskCollectionView projects result type SIPTaskCollection to projected
// type SIPTaskCollectionView using the "default" view.
func newSIPTaskCollectionView(res SIPTaskCollection) ingestviews.SIPTaskCollectionView {
	vres := make(ingestviews.SIPTaskCollectionView, len(res))
	for i, n := range res {
		vres[i] = newSIPTaskView(n)
	}
	return vres
}

// newSIPTask converts projected type SIPTask to service type SIPTask.
func newSIPTask(vres *ingestviews.SIPTaskView) *SIPTask {
	res := &SIPTask{
		CompletedAt: vres.CompletedAt,
		Note:        vres.Note,
		WorkflowID:  vres.WorkflowID,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.TaskID != nil {
		res.TaskID = *vres.TaskID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.StartedAt != nil {
		res.StartedAt = *vres.StartedAt
	}
	return res
}

// newSIPTaskView projects result type SIPTask to projected type SIPTaskView
// using the "default" view.
func newSIPTaskView(res *SIPTask) *ingestviews.SIPTaskView {
	vres := &ingestviews.SIPTaskView{
		ID:          &res.ID,
		TaskID:      &res.TaskID,
		Name:        &res.Name,
		Status:      &res.Status,
		StartedAt:   &res.StartedAt,
		CompletedAt: res.CompletedAt,
		Note:        res.Note,
		WorkflowID:  res.WorkflowID,
	}
	return vres
}
