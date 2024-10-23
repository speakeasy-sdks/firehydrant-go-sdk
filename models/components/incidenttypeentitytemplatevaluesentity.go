// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Runbooks - A hash mapping runbook IDs to runbook names.
type Runbooks struct {
}

type IncidentTypeEntityTemplateValuesEntity struct {
	Services        []IncidentTypeEntityTemplateImpactEntity `json:"services,omitempty"`
	Functionalities []IncidentTypeEntityTemplateImpactEntity `json:"functionalities,omitempty"`
	Environments    []IncidentTypeEntityTemplateImpactEntity `json:"environments,omitempty"`
	// A hash mapping runbook IDs to runbook names.
	Runbooks *Runbooks    `json:"runbooks,omitempty"`
	Teams    []TeamEntity `json:"teams,omitempty"`
}

func (o *IncidentTypeEntityTemplateValuesEntity) GetServices() []IncidentTypeEntityTemplateImpactEntity {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *IncidentTypeEntityTemplateValuesEntity) GetFunctionalities() []IncidentTypeEntityTemplateImpactEntity {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *IncidentTypeEntityTemplateValuesEntity) GetEnvironments() []IncidentTypeEntityTemplateImpactEntity {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *IncidentTypeEntityTemplateValuesEntity) GetRunbooks() *Runbooks {
	if o == nil {
		return nil
	}
	return o.Runbooks
}

func (o *IncidentTypeEntityTemplateValuesEntity) GetTeams() []TeamEntity {
	if o == nil {
		return nil
	}
	return o.Teams
}
