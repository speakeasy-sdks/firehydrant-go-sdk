// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PutV1RunbooksRunbookIDOwner - An object representing a Team that owns the runbook
type PutV1RunbooksRunbookIDOwner struct {
	ID *string `json:"id,omitempty"`
}

func (o *PutV1RunbooksRunbookIDOwner) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PutV1RunbooksRunbookIDSeverities struct {
	ID *string `json:"id,omitempty"`
}

func (o *PutV1RunbooksRunbookIDSeverities) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PutV1RunbooksRunbookIDServices struct {
	ID *string `json:"id,omitempty"`
}

func (o *PutV1RunbooksRunbookIDServices) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Environments struct {
	ID *string `json:"id,omitempty"`
}

func (o *Environments) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type PutV1RunbooksRunbookIDAttachmentRule struct {
	// The JSON logic for the attaching the runbook
	Logic string `json:"logic"`
	// The user data for the rule
	UserData *string `json:"user_data,omitempty"`
}

func (o *PutV1RunbooksRunbookIDAttachmentRule) GetLogic() string {
	if o == nil {
		return ""
	}
	return o.Logic
}

func (o *PutV1RunbooksRunbookIDAttachmentRule) GetUserData() *string {
	if o == nil {
		return nil
	}
	return o.UserData
}

type PutV1RunbooksRunbookIDRule struct {
	// The JSON logic for the rule
	Logic string `json:"logic"`
	// The user data for the rule
	UserData *string `json:"user_data,omitempty"`
}

func (o *PutV1RunbooksRunbookIDRule) GetLogic() string {
	if o == nil {
		return ""
	}
	return o.Logic
}

func (o *PutV1RunbooksRunbookIDRule) GetUserData() *string {
	if o == nil {
		return nil
	}
	return o.UserData
}

type PutV1RunbooksRunbookIDSteps struct {
	// ID of step to be updated
	StepID *string `json:"step_id,omitempty"`
	// Name for step
	Name string `json:"name"`
	// ID of action to use for this step.
	ActionID string                      `json:"action_id"`
	Rule     *PutV1RunbooksRunbookIDRule `json:"rule,omitempty"`
}

func (o *PutV1RunbooksRunbookIDSteps) GetStepID() *string {
	if o == nil {
		return nil
	}
	return o.StepID
}

func (o *PutV1RunbooksRunbookIDSteps) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutV1RunbooksRunbookIDSteps) GetActionID() string {
	if o == nil {
		return ""
	}
	return o.ActionID
}

func (o *PutV1RunbooksRunbookIDSteps) GetRule() *PutV1RunbooksRunbookIDRule {
	if o == nil {
		return nil
	}
	return o.Rule
}

// PutV1RunbooksRunbookID - Update a runbook and any attachment rules associated with it. This endpoint is used to configure nearly everything
// about a runbook, including but not limited to the steps, environments, attachment rules, and severities.
type PutV1RunbooksRunbookID struct {
	Name        *string `json:"name,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	// Whether or not this runbook is a tutorial runbook
	Tutorial *bool `json:"tutorial,omitempty"`
	// An object representing a Team that owns the runbook
	Owner          *PutV1RunbooksRunbookIDOwner          `json:"owner,omitempty"`
	Severities     []PutV1RunbooksRunbookIDSeverities    `json:"severities,omitempty"`
	Services       []PutV1RunbooksRunbookIDServices      `json:"services,omitempty"`
	Environments   []Environments                        `json:"environments,omitempty"`
	AttachmentRule *PutV1RunbooksRunbookIDAttachmentRule `json:"attachment_rule,omitempty"`
	Steps          []PutV1RunbooksRunbookIDSteps         `json:"steps,omitempty"`
	// Whether or not this runbook should be automatically attached to restricted incidents. Note that setting this to `true` will prevent it from being attached to public incidents, even manually. Defaults to `false`.
	AutoAttachToRestrictedIncidents *bool `json:"auto_attach_to_restricted_incidents,omitempty"`
}

func (o *PutV1RunbooksRunbookID) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PutV1RunbooksRunbookID) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PutV1RunbooksRunbookID) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PutV1RunbooksRunbookID) GetTutorial() *bool {
	if o == nil {
		return nil
	}
	return o.Tutorial
}

func (o *PutV1RunbooksRunbookID) GetOwner() *PutV1RunbooksRunbookIDOwner {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *PutV1RunbooksRunbookID) GetSeverities() []PutV1RunbooksRunbookIDSeverities {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *PutV1RunbooksRunbookID) GetServices() []PutV1RunbooksRunbookIDServices {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *PutV1RunbooksRunbookID) GetEnvironments() []Environments {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *PutV1RunbooksRunbookID) GetAttachmentRule() *PutV1RunbooksRunbookIDAttachmentRule {
	if o == nil {
		return nil
	}
	return o.AttachmentRule
}

func (o *PutV1RunbooksRunbookID) GetSteps() []PutV1RunbooksRunbookIDSteps {
	if o == nil {
		return nil
	}
	return o.Steps
}

func (o *PutV1RunbooksRunbookID) GetAutoAttachToRestrictedIncidents() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAttachToRestrictedIncidents
}
