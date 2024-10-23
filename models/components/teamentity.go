// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
	"time"
)

// TeamEntity model
type TeamEntity struct {
	ID                      *string                                    `json:"id,omitempty"`
	Name                    *string                                    `json:"name,omitempty"`
	Description             *string                                    `json:"description,omitempty"`
	Slug                    *string                                    `json:"slug,omitempty"`
	CreatedAt               *time.Time                                 `json:"created_at,omitempty"`
	UpdatedAt               *time.Time                                 `json:"updated_at,omitempty"`
	SignalsIcalURL          *string                                    `json:"signals_ical_url,omitempty"`
	CreatedBy               *AuthorEntity                              `json:"created_by,omitempty"`
	SlackChannel            *IntegrationsSlackSlackChannelEntity       `json:"slack_channel,omitempty"`
	MsTeamsChannel          *IntegrationsMicrosoftTeamsV2ChannelEntity `json:"ms_teams_channel,omitempty"`
	Memberships             []MembershipEntity                         `json:"memberships,omitempty"`
	OwnedChecklistTemplates []ChecklistTemplateEntity                  `json:"owned_checklist_templates,omitempty"`
	OwnedFunctionalities    []FunctionalityEntity                      `json:"owned_functionalities,omitempty"`
	OwnedServices           []ServiceEntity                            `json:"owned_services,omitempty"`
	OwnedRunbooks           []SlimRunbookEntity                        `json:"owned_runbooks,omitempty"`
	RespondingServices      []ServiceEntity                            `json:"responding_services,omitempty"`
	Services                []ServiceEntity                            `json:"services,omitempty"`
	Functionalities         []FunctionalityEntity                      `json:"functionalities,omitempty"`
}

func (t TeamEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TeamEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TeamEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TeamEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TeamEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TeamEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *TeamEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TeamEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TeamEntity) GetSignalsIcalURL() *string {
	if o == nil {
		return nil
	}
	return o.SignalsIcalURL
}

func (o *TeamEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TeamEntity) GetSlackChannel() *IntegrationsSlackSlackChannelEntity {
	if o == nil {
		return nil
	}
	return o.SlackChannel
}

func (o *TeamEntity) GetMsTeamsChannel() *IntegrationsMicrosoftTeamsV2ChannelEntity {
	if o == nil {
		return nil
	}
	return o.MsTeamsChannel
}

func (o *TeamEntity) GetMemberships() []MembershipEntity {
	if o == nil {
		return nil
	}
	return o.Memberships
}

func (o *TeamEntity) GetOwnedChecklistTemplates() []ChecklistTemplateEntity {
	if o == nil {
		return nil
	}
	return o.OwnedChecklistTemplates
}

func (o *TeamEntity) GetOwnedFunctionalities() []FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.OwnedFunctionalities
}

func (o *TeamEntity) GetOwnedServices() []ServiceEntity {
	if o == nil {
		return nil
	}
	return o.OwnedServices
}

func (o *TeamEntity) GetOwnedRunbooks() []SlimRunbookEntity {
	if o == nil {
		return nil
	}
	return o.OwnedRunbooks
}

func (o *TeamEntity) GetRespondingServices() []ServiceEntity {
	if o == nil {
		return nil
	}
	return o.RespondingServices
}

func (o *TeamEntity) GetServices() []ServiceEntity {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *TeamEntity) GetFunctionalities() []FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.Functionalities
}