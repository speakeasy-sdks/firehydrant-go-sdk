// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
	"time"
)

type Config struct {
}

type StepElements struct {
}

type RunbooksExecutionStepEntity struct {
	ID              *string        `json:"id,omitempty"`
	Name            *string        `json:"name,omitempty"`
	ActionSlug      *string        `json:"action_slug,omitempty"`
	ActionType      *string        `json:"action_type,omitempty"`
	IntegrationName *string        `json:"integration_name,omitempty"`
	IntegrationSlug *string        `json:"integration_slug,omitempty"`
	Automatic       *bool          `json:"automatic,omitempty"`
	Config          *Config        `json:"config,omitempty"`
	StepElements    []StepElements `json:"step_elements,omitempty"`
	Executable      *bool          `json:"executable,omitempty"`
	Repeats         *bool          `json:"repeats,omitempty"`
	// ISO8601 formatted duration string
	RepeatsDuration *string                               `json:"repeats_duration,omitempty"`
	RepeatsAt       *time.Time                            `json:"repeats_at,omitempty"`
	HasBeenRerun    *bool                                 `json:"has_been_rerun,omitempty"`
	HasBeenRetried  *bool                                 `json:"has_been_retried,omitempty"`
	Execution       *RunbooksExecutionStepExecutionEntity `json:"execution,omitempty"`
	Repeatable      *bool                                 `json:"repeatable,omitempty"`
	Rule            *RulesRuleEntity                      `json:"rule,omitempty"`
}

func (r RunbooksExecutionStepEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RunbooksExecutionStepEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RunbooksExecutionStepEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RunbooksExecutionStepEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RunbooksExecutionStepEntity) GetActionSlug() *string {
	if o == nil {
		return nil
	}
	return o.ActionSlug
}

func (o *RunbooksExecutionStepEntity) GetActionType() *string {
	if o == nil {
		return nil
	}
	return o.ActionType
}

func (o *RunbooksExecutionStepEntity) GetIntegrationName() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationName
}

func (o *RunbooksExecutionStepEntity) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *RunbooksExecutionStepEntity) GetAutomatic() *bool {
	if o == nil {
		return nil
	}
	return o.Automatic
}

func (o *RunbooksExecutionStepEntity) GetConfig() *Config {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *RunbooksExecutionStepEntity) GetStepElements() []StepElements {
	if o == nil {
		return nil
	}
	return o.StepElements
}

func (o *RunbooksExecutionStepEntity) GetExecutable() *bool {
	if o == nil {
		return nil
	}
	return o.Executable
}

func (o *RunbooksExecutionStepEntity) GetRepeats() *bool {
	if o == nil {
		return nil
	}
	return o.Repeats
}

func (o *RunbooksExecutionStepEntity) GetRepeatsDuration() *string {
	if o == nil {
		return nil
	}
	return o.RepeatsDuration
}

func (o *RunbooksExecutionStepEntity) GetRepeatsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RepeatsAt
}

func (o *RunbooksExecutionStepEntity) GetHasBeenRerun() *bool {
	if o == nil {
		return nil
	}
	return o.HasBeenRerun
}

func (o *RunbooksExecutionStepEntity) GetHasBeenRetried() *bool {
	if o == nil {
		return nil
	}
	return o.HasBeenRetried
}

func (o *RunbooksExecutionStepEntity) GetExecution() *RunbooksExecutionStepExecutionEntity {
	if o == nil {
		return nil
	}
	return o.Execution
}

func (o *RunbooksExecutionStepEntity) GetRepeatable() *bool {
	if o == nil {
		return nil
	}
	return o.Repeatable
}

func (o *RunbooksExecutionStepEntity) GetRule() *RulesRuleEntity {
	if o == nil {
		return nil
	}
	return o.Rule
}
