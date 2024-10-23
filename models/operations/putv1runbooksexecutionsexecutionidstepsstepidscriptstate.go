// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
	StepID      string `pathParam:"style=simple,explode=false,name=step_id"`
	State       string `pathParam:"style=simple,explode=false,name=state"`
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateRequest) GetStepID() string {
	if o == nil {
		return ""
	}
	return o.StepID
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateRequest) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

type PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updates the execution's step.
	RunbooksExecutionEntity *components.RunbooksExecutionEntity
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateResponse) GetRunbooksExecutionEntity() *components.RunbooksExecutionEntity {
	if o == nil {
		return nil
	}
	return o.RunbooksExecutionEntity
}
