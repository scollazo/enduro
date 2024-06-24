// Code generated by goa v3.15.2, DO NOT EDIT.
//
// package HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	package_ "github.com/artefactual-sdps/enduro/internal/api/gen/package_"
	goa "goa.design/goa/v3/pkg"
)

// BuildMonitorRequestPayload builds the payload for the package
// monitor_request endpoint from CLI flags.
func BuildMonitorRequestPayload(package_MonitorRequestToken string) (*package_.MonitorRequestPayload, error) {
	var token *string
	{
		if package_MonitorRequestToken != "" {
			token = &package_MonitorRequestToken
		}
	}
	v := &package_.MonitorRequestPayload{}
	v.Token = token

	return v, nil
}

// BuildMonitorPayload builds the payload for the package monitor endpoint from
// CLI flags.
func BuildMonitorPayload(package_MonitorTicket string) (*package_.MonitorPayload, error) {
	var ticket *string
	{
		if package_MonitorTicket != "" {
			ticket = &package_MonitorTicket
		}
	}
	v := &package_.MonitorPayload{}
	v.Ticket = ticket

	return v, nil
}

// BuildListPayload builds the payload for the package list endpoint from CLI
// flags.
func BuildListPayload(package_ListName string, package_ListAipID string, package_ListEarliestCreatedTime string, package_ListLatestCreatedTime string, package_ListLocationID string, package_ListStatus string, package_ListCursor string, package_ListToken string) (*package_.ListPayload, error) {
	var err error
	var name *string
	{
		if package_ListName != "" {
			name = &package_ListName
		}
	}
	var aipID *string
	{
		if package_ListAipID != "" {
			aipID = &package_ListAipID
			err = goa.MergeErrors(err, goa.ValidateFormat("aip_id", *aipID, goa.FormatUUID))
			if err != nil {
				return nil, err
			}
		}
	}
	var earliestCreatedTime *string
	{
		if package_ListEarliestCreatedTime != "" {
			earliestCreatedTime = &package_ListEarliestCreatedTime
			err = goa.MergeErrors(err, goa.ValidateFormat("earliest_created_time", *earliestCreatedTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var latestCreatedTime *string
	{
		if package_ListLatestCreatedTime != "" {
			latestCreatedTime = &package_ListLatestCreatedTime
			err = goa.MergeErrors(err, goa.ValidateFormat("latest_created_time", *latestCreatedTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var locationID *string
	{
		if package_ListLocationID != "" {
			locationID = &package_ListLocationID
			err = goa.MergeErrors(err, goa.ValidateFormat("location_id", *locationID, goa.FormatUUID))
			if err != nil {
				return nil, err
			}
		}
	}
	var status *string
	{
		if package_ListStatus != "" {
			status = &package_ListStatus
			if !(*status == "new" || *status == "in progress" || *status == "done" || *status == "error" || *status == "unknown" || *status == "queued" || *status == "abandoned" || *status == "pending") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("status", *status, []any{"new", "in progress", "done", "error", "unknown", "queued", "abandoned", "pending"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var cursor *string
	{
		if package_ListCursor != "" {
			cursor = &package_ListCursor
		}
	}
	var token *string
	{
		if package_ListToken != "" {
			token = &package_ListToken
		}
	}
	v := &package_.ListPayload{}
	v.Name = name
	v.AipID = aipID
	v.EarliestCreatedTime = earliestCreatedTime
	v.LatestCreatedTime = latestCreatedTime
	v.LocationID = locationID
	v.Status = status
	v.Cursor = cursor
	v.Token = token

	return v, nil
}

// BuildShowPayload builds the payload for the package show endpoint from CLI
// flags.
func BuildShowPayload(package_ShowID string, package_ShowToken string) (*package_.ShowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_ShowID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token *string
	{
		if package_ShowToken != "" {
			token = &package_ShowToken
		}
	}
	v := &package_.ShowPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildPreservationActionsPayload builds the payload for the package
// preservation_actions endpoint from CLI flags.
func BuildPreservationActionsPayload(package_PreservationActionsID string, package_PreservationActionsToken string) (*package_.PreservationActionsPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_PreservationActionsID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token *string
	{
		if package_PreservationActionsToken != "" {
			token = &package_PreservationActionsToken
		}
	}
	v := &package_.PreservationActionsPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildConfirmPayload builds the payload for the package confirm endpoint from
// CLI flags.
func BuildConfirmPayload(package_ConfirmBody string, package_ConfirmID string, package_ConfirmToken string) (*package_.ConfirmPayload, error) {
	var err error
	var body ConfirmRequestBody
	{
		err = json.Unmarshal([]byte(package_ConfirmBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"location_id\": \"d1845cb6-a5ea-474a-9ab8-26f9bcd919f5\"\n   }'")
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_ConfirmID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token *string
	{
		if package_ConfirmToken != "" {
			token = &package_ConfirmToken
		}
	}
	v := &package_.ConfirmPayload{
		LocationID: body.LocationID,
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildRejectPayload builds the payload for the package reject endpoint from
// CLI flags.
func BuildRejectPayload(package_RejectID string, package_RejectToken string) (*package_.RejectPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_RejectID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token *string
	{
		if package_RejectToken != "" {
			token = &package_RejectToken
		}
	}
	v := &package_.RejectPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildMovePayload builds the payload for the package move endpoint from CLI
// flags.
func BuildMovePayload(package_MoveBody string, package_MoveID string, package_MoveToken string) (*package_.MovePayload, error) {
	var err error
	var body MoveRequestBody
	{
		err = json.Unmarshal([]byte(package_MoveBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"location_id\": \"d1845cb6-a5ea-474a-9ab8-26f9bcd919f5\"\n   }'")
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_MoveID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token *string
	{
		if package_MoveToken != "" {
			token = &package_MoveToken
		}
	}
	v := &package_.MovePayload{
		LocationID: body.LocationID,
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildMoveStatusPayload builds the payload for the package move_status
// endpoint from CLI flags.
func BuildMoveStatusPayload(package_MoveStatusID string, package_MoveStatusToken string) (*package_.MoveStatusPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_MoveStatusID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	var token *string
	{
		if package_MoveStatusToken != "" {
			token = &package_MoveStatusToken
		}
	}
	v := &package_.MoveStatusPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}
