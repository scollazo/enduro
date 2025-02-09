package storage

import (
	"context"
	"fmt"
	"time"

	temporalsdk_api_enums "go.temporal.io/api/enums/v1"
	temporalsdk_client "go.temporal.io/sdk/client"

	"github.com/artefactual-labs/enduro/internal/temporal"
)

const (
	CopyToPermanentLocationActivityName = "copy-to-permanent-location-activity"
	DeleteFromLocationActivityName      = "delete-from-location-activity"
	StorageUploadWorkflowName           = "storage-upload-workflow"
	StorageMoveWorkflowName             = "storage-move-workflow"
	UploadDoneSignalName                = "upload-done-signal"
)

type StorageUploadWorkflowRequest struct {
	AIPID string
}

type StorageMoveWorkflowRequest struct {
	AIPID    string
	Location string
}

type CopyToPermanentLocationActivityParams struct {
	AIPID    string
	Location string
}

type UploadDoneSignal struct{}

func InitStorageUploadWorkflow(ctx context.Context, tc temporalsdk_client.Client, req *StorageUploadWorkflowRequest) (temporalsdk_client.WorkflowRun, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	opts := temporalsdk_client.StartWorkflowOptions{
		ID:                    fmt.Sprintf("%s-%s", StorageUploadWorkflowName, req.AIPID),
		TaskQueue:             temporal.GlobalTaskQueue,
		WorkflowIDReusePolicy: temporalsdk_api_enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}
	exec, err := tc.ExecuteWorkflow(ctx, opts, StorageUploadWorkflowName, req)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func InitStorageMoveWorkflow(ctx context.Context, tc temporalsdk_client.Client, req *StorageMoveWorkflowRequest) (temporalsdk_client.WorkflowRun, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	opts := temporalsdk_client.StartWorkflowOptions{
		ID:                    fmt.Sprintf("%s-%s", StorageMoveWorkflowName, req.AIPID),
		TaskQueue:             temporal.GlobalTaskQueue,
		WorkflowIDReusePolicy: temporalsdk_api_enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}
	exec, err := tc.ExecuteWorkflow(ctx, opts, StorageMoveWorkflowName, req)
	if err != nil {
		return nil, err
	}
	return exec, nil
}
