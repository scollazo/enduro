package workflows

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	temporalsdk_temporal "go.temporal.io/sdk/temporal"
	temporalsdk_workflow "go.temporal.io/sdk/workflow"

	"github.com/artefactual-sdps/enduro/internal/storage"
	"github.com/artefactual-sdps/enduro/internal/storage/enums"
)

type StorageMoveWorkflow struct {
	storagesvc storage.Service
}

func NewStorageMoveWorkflow(storagesvc storage.Service) *StorageMoveWorkflow {
	return &StorageMoveWorkflow{
		storagesvc: storagesvc,
	}
}

func (w *StorageMoveWorkflow) Execute(
	ctx temporalsdk_workflow.Context,
	req storage.StorageMoveWorkflowRequest,
) (e error) {
	// Set AIP status to moving.
	{
		if err := w.updateAIPStatus(ctx, enums.AIPStatusMoving, req.AIPID); err != nil {
			return err
		}
	}

	// Create persistence workflow.
	workflowDBID, err := createWorkflow(ctx, w.storagesvc, req.AIPID, enums.WorkflowTypeMoveAip)
	if err != nil {
		return err
	}

	// Complete persistence workflow.
	defer func() {
		status := enums.WorkflowStatusDone
		if e != nil {
			status = enums.WorkflowStatusError
		}

		err := completeWorkflow(ctx, w.storagesvc, workflowDBID, status)
		if err != nil {
			e = errors.Join(e, err)
		}
	}()

	// Create copy AIP task.
	copyTaskID, err := createTask(ctx, w.storagesvc, workflowDBID, "Copy AIP", "Copying AIP to target location")
	if err != nil {
		return err
	}

	// Copy AIP from its current bucket to a new permanent location bucket
	{
		activityOpts := temporalsdk_workflow.WithActivityOptions(ctx, temporalsdk_workflow.ActivityOptions{
			StartToCloseTimeout: time.Hour * 2,
			RetryPolicy: &temporalsdk_temporal.RetryPolicy{
				InitialInterval:    time.Second,
				BackoffCoefficient: 2,
				MaximumInterval:    time.Minute * 10,
				MaximumAttempts:    5,
				NonRetryableErrorTypes: []string{
					"TemporalTimeout:StartToClose",
				},
			},
		})
		err := temporalsdk_workflow.ExecuteActivity(
			activityOpts,
			storage.CopyToPermanentLocationActivityName,
			&storage.CopyToPermanentLocationActivityParams{
				AIPID:      req.AIPID,
				LocationID: req.LocationID,
			},
		).Get(activityOpts, nil)

		// Complete copy AIP task.
		taskStatus := enums.TaskStatusDone
		taskNote := "AIP copied to target location"
		if err != nil {
			taskStatus = enums.TaskStatusError
			taskNote = fmt.Sprintf("Failed to copy AIP to target location\n%s", err.Error())
		}
		taskErr := completeTask(ctx, w.storagesvc, copyTaskID, taskStatus, taskNote)
		if err = errors.Join(err, taskErr); err != nil {
			return err
		}
	}

	// Create delete AIP task.
	deleteTaskID, err := createTask(ctx, w.storagesvc, workflowDBID, "Delete AIP", "Deleting AIP from source location")
	if err != nil {
		return err
	}

	// Delete AIP from its current bucket
	{
		activityOpts := localActivityOptions(ctx)
		err := temporalsdk_workflow.ExecuteLocalActivity(
			activityOpts,
			storage.DeleteFromLocationLocalActivity,
			w.storagesvc,
			&storage.DeleteFromLocationLocalActivityParams{
				AIPID: req.AIPID,
			},
		).Get(activityOpts, nil)

		// Complete delete AIP task.
		taskStatus := enums.TaskStatusDone
		taskNote := "AIP deleted from source location"
		if err != nil {
			taskStatus = enums.TaskStatusError
			taskNote = fmt.Sprintf("Failed to delete AIP from source location\n%s", err.Error())
		}
		taskErr := completeTask(ctx, w.storagesvc, deleteTaskID, taskStatus, taskNote)
		if err = errors.Join(err, taskErr); err != nil {
			return err
		}
	}

	// Update AIP location
	{
		activityOpts := localActivityOptions(ctx)
		err := temporalsdk_workflow.ExecuteLocalActivity(
			activityOpts,
			storage.UpdateAIPLocationLocalActivity,
			w.storagesvc,
			&storage.UpdateAIPLocationLocalActivityParams{
				AIPID:      req.AIPID,
				LocationID: req.LocationID,
			},
		).Get(activityOpts, nil)
		if err != nil {
			return err
		}
	}

	// Set AIP status to stored.
	{
		if err := w.updateAIPStatus(ctx, enums.AIPStatusStored, req.AIPID); err != nil {
			return err
		}
	}

	return nil
}

func (w *StorageMoveWorkflow) updateAIPStatus(
	ctx temporalsdk_workflow.Context,
	st enums.AIPStatus,
	aipID uuid.UUID,
) error {
	activityOpts := temporalsdk_workflow.WithLocalActivityOptions(ctx, temporalsdk_workflow.LocalActivityOptions{
		ScheduleToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporalsdk_temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    3,
		},
	})

	params := &storage.UpdateAIPStatusLocalActivityParams{
		AIPID:  aipID,
		Status: st,
	}

	return temporalsdk_workflow.ExecuteLocalActivity(activityOpts, storage.UpdateAIPStatusLocalActivity, w.storagesvc, params).
		Get(activityOpts, nil)
}
