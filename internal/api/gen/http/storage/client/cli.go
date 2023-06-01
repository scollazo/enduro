// Code generated by goa v3.11.3, DO NOT EDIT.
//
// storage HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	goa "goa.design/goa/v3/pkg"
)

// BuildSubmitPayload builds the payload for the storage submit endpoint from
// CLI flags.
func BuildSubmitPayload(storageSubmitBody string, storageSubmitAipID string, storageSubmitOauthToken string) (*storage.SubmitPayload, error) {
	var err error
	var body SubmitRequestBody
	{
		err = json.Unmarshal([]byte(storageSubmitBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Aperiam velit.\"\n   }'")
		}
	}
	var aipID string
	{
		aipID = storageSubmitAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageSubmitOauthToken != "" {
			oauthToken = &storageSubmitOauthToken
		}
	}
	v := &storage.SubmitPayload{
		Name: body.Name,
	}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildUpdatePayload builds the payload for the storage update endpoint from
// CLI flags.
func BuildUpdatePayload(storageUpdateAipID string, storageUpdateOauthToken string) (*storage.UpdatePayload, error) {
	var err error
	var aipID string
	{
		aipID = storageUpdateAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageUpdateOauthToken != "" {
			oauthToken = &storageUpdateOauthToken
		}
	}
	v := &storage.UpdatePayload{}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildDownloadPayload builds the payload for the storage download endpoint
// from CLI flags.
func BuildDownloadPayload(storageDownloadAipID string, storageDownloadOauthToken string) (*storage.DownloadPayload, error) {
	var err error
	var aipID string
	{
		aipID = storageDownloadAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageDownloadOauthToken != "" {
			oauthToken = &storageDownloadOauthToken
		}
	}
	v := &storage.DownloadPayload{}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildLocationsPayload builds the payload for the storage locations endpoint
// from CLI flags.
func BuildLocationsPayload(storageLocationsOauthToken string) (*storage.LocationsPayload, error) {
	var oauthToken *string
	{
		if storageLocationsOauthToken != "" {
			oauthToken = &storageLocationsOauthToken
		}
	}
	v := &storage.LocationsPayload{}
	v.OauthToken = oauthToken

	return v, nil
}

// BuildAddLocationPayload builds the payload for the storage add_location
// endpoint from CLI flags.
func BuildAddLocationPayload(storageAddLocationBody string, storageAddLocationOauthToken string) (*storage.AddLocationPayload, error) {
	var err error
	var body AddLocationRequestBody
	{
		err = json.Unmarshal([]byte(storageAddLocationBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"config\": {\n         \"Type\": \"s3\",\n         \"Value\": \"\\\"JSON\\\"\"\n      },\n      \"description\": \"Architecto sed voluptas quasi vel.\",\n      \"name\": \"Repellat commodi.\",\n      \"purpose\": \"aip_store\",\n      \"source\": \"unspecified\"\n   }'")
		}
		if !(body.Source == "unspecified" || body.Source == "minio" || body.Source == "sftp") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.source", body.Source, []any{"unspecified", "minio", "sftp"}))
		}
		if !(body.Purpose == "unspecified" || body.Purpose == "aip_store") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.purpose", body.Purpose, []any{"unspecified", "aip_store"}))
		}
		if body.Config != nil {
			if !(body.Config.Type == "s3" || body.Config.Type == "sftp") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.config.Type", body.Config.Type, []any{"s3", "sftp"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageAddLocationOauthToken != "" {
			oauthToken = &storageAddLocationOauthToken
		}
	}
	v := &storage.AddLocationPayload{
		Name:        body.Name,
		Description: body.Description,
		Source:      body.Source,
		Purpose:     body.Purpose,
	}
	if body.Config != nil {
		switch body.Config.Type {
		case "s3":
			var val *storage.S3Config
			json.Unmarshal([]byte(body.Config.Value), &val)
			v.Config = val
		case "sftp":
			var val *storage.SFTPConfig
			json.Unmarshal([]byte(body.Config.Value), &val)
			v.Config = val
		}
	}
	v.OauthToken = oauthToken

	return v, nil
}

// BuildMovePayload builds the payload for the storage move endpoint from CLI
// flags.
func BuildMovePayload(storageMoveBody string, storageMoveAipID string, storageMoveOauthToken string) (*storage.MovePayload, error) {
	var err error
	var body MoveRequestBody
	{
		err = json.Unmarshal([]byte(storageMoveBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"location_id\": \"Enim quis vel ipsa laudantium harum sunt.\"\n   }'")
		}
	}
	var aipID string
	{
		aipID = storageMoveAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageMoveOauthToken != "" {
			oauthToken = &storageMoveOauthToken
		}
	}
	v := &storage.MovePayload{
		LocationID: body.LocationID,
	}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildMoveStatusPayload builds the payload for the storage move_status
// endpoint from CLI flags.
func BuildMoveStatusPayload(storageMoveStatusAipID string, storageMoveStatusOauthToken string) (*storage.MoveStatusPayload, error) {
	var err error
	var aipID string
	{
		aipID = storageMoveStatusAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageMoveStatusOauthToken != "" {
			oauthToken = &storageMoveStatusOauthToken
		}
	}
	v := &storage.MoveStatusPayload{}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildRejectPayload builds the payload for the storage reject endpoint from
// CLI flags.
func BuildRejectPayload(storageRejectAipID string, storageRejectOauthToken string) (*storage.RejectPayload, error) {
	var err error
	var aipID string
	{
		aipID = storageRejectAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageRejectOauthToken != "" {
			oauthToken = &storageRejectOauthToken
		}
	}
	v := &storage.RejectPayload{}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildShowPayload builds the payload for the storage show endpoint from CLI
// flags.
func BuildShowPayload(storageShowAipID string, storageShowOauthToken string) (*storage.ShowPayload, error) {
	var err error
	var aipID string
	{
		aipID = storageShowAipID
		err = goa.MergeErrors(err, goa.ValidateFormat("aipID", aipID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageShowOauthToken != "" {
			oauthToken = &storageShowOauthToken
		}
	}
	v := &storage.ShowPayload{}
	v.AipID = aipID
	v.OauthToken = oauthToken

	return v, nil
}

// BuildShowLocationPayload builds the payload for the storage show_location
// endpoint from CLI flags.
func BuildShowLocationPayload(storageShowLocationUUID string, storageShowLocationOauthToken string) (*storage.ShowLocationPayload, error) {
	var err error
	var uuid string
	{
		uuid = storageShowLocationUUID
		err = goa.MergeErrors(err, goa.ValidateFormat("uuid", uuid, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageShowLocationOauthToken != "" {
			oauthToken = &storageShowLocationOauthToken
		}
	}
	v := &storage.ShowLocationPayload{}
	v.UUID = uuid
	v.OauthToken = oauthToken

	return v, nil
}

// BuildLocationPackagesPayload builds the payload for the storage
// location_packages endpoint from CLI flags.
func BuildLocationPackagesPayload(storageLocationPackagesUUID string, storageLocationPackagesOauthToken string) (*storage.LocationPackagesPayload, error) {
	var err error
	var uuid string
	{
		uuid = storageLocationPackagesUUID
		err = goa.MergeErrors(err, goa.ValidateFormat("uuid", uuid, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var oauthToken *string
	{
		if storageLocationPackagesOauthToken != "" {
			oauthToken = &storageLocationPackagesOauthToken
		}
	}
	v := &storage.LocationPackagesPayload{}
	v.UUID = uuid
	v.OauthToken = oauthToken

	return v, nil
}
