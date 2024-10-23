// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
	"time"
)

// IncidentEntityLabels - A key/value of labels
type IncidentEntityLabels struct {
}

type RetroExports struct {
}

// IncidentEntity model
type IncidentEntity struct {
	// UUID of the Incident
	ID *string `json:"id,omitempty"`
	// Name of the incident
	Name *string `json:"name,omitempty"`
	// The time the incident was opened
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time the incident started
	StartedAt *time.Time `json:"started_at,omitempty"`
	// The time the incident was archived
	DiscardedAt           *time.Time `json:"discarded_at,omitempty"`
	Summary               *string    `json:"summary,omitempty"`
	CustomerImpactSummary *string    `json:"customer_impact_summary,omitempty"`
	Description           *string    `json:"description,omitempty"`
	// The type/slug of the current milestone. Will be one of the currently configured milestones for the given incident.
	CurrentMilestone *string `json:"current_milestone,omitempty"`
	// Incident number
	Number            *int     `json:"number,omitempty"`
	Priority          *string  `json:"priority,omitempty"`
	Severity          *string  `json:"severity,omitempty"`
	SeverityColor     *string  `json:"severity_color,omitempty"`
	SeverityImpact    *string  `json:"severity_impact,omitempty"`
	SeverityCondition *string  `json:"severity_condition,omitempty"`
	TagList           []string `json:"tag_list,omitempty"`
	// SeverityMatrix_ImpactEntity model
	SeverityImpactObject *SeverityMatrixImpactEntity `json:"severity_impact_object,omitempty"`
	// SeverityMatrix_ConditionEntity model
	SeverityConditionObject *SeverityMatrixConditionEntity `json:"severity_condition_object,omitempty"`
	PrivateID               *string                        `json:"private_id,omitempty"`
	OrganizationID          *string                        `json:"organization_id,omitempty"`
	// DEPRECATED: Please use lifecycle phases instead
	Milestones            []IncidentsMilestoneEntity            `json:"milestones,omitempty"`
	LifecyclePhases       []IncidentsLifecyclePhaseEntity       `json:"lifecycle_phases,omitempty"`
	LifecycleMeasurements []IncidentsLifecycleMeasurementEntity `json:"lifecycle_measurements,omitempty"`
	Active                *bool                                 `json:"active,omitempty"`
	// A key/value of labels
	Labels               *IncidentEntityLabels           `json:"labels,omitempty"`
	RoleAssignments      []IncidentsRoleAssignmentEntity `json:"role_assignments,omitempty"`
	StatusPages          []IncidentsStatusPageEntity     `json:"status_pages,omitempty"`
	IncidentURL          *string                         `json:"incident_url,omitempty"`
	PrivateStatusPageURL *string                         `json:"private_status_page_url,omitempty"`
	Organization         *OrganizationEntity             `json:"organization,omitempty"`
	CustomersImpacted    *int                            `json:"customers_impacted,omitempty"`
	MonetaryImpact       *int                            `json:"monetary_impact,omitempty"`
	MonetaryImpactCents  *int                            `json:"monetary_impact_cents,omitempty"`
	LastUpdate           *string                         `json:"last_update,omitempty"`
	// Event_NoteEntity model
	LastNote          *EventNoteEntity `json:"last_note,omitempty"`
	ReportID          *string          `json:"report_id,omitempty"`
	AiIncidentSummary *string          `json:"ai_incident_summary,omitempty"`
	Services          []SuccinctEntity `json:"services,omitempty"`
	Environments      []SuccinctEntity `json:"environments,omitempty"`
	Functionalities   []SuccinctEntity `json:"functionalities,omitempty"`
	ChannelName       *string          `json:"channel_name,omitempty"`
	ChannelReference  *string          `json:"channel_reference,omitempty"`
	ChannelID         *string          `json:"channel_id,omitempty"`
	// inoperative: 0, operational: 1, archived: 2
	ChannelStatus   *string                 `json:"channel_status,omitempty"`
	IncidentTickets []TicketingTicketEntity `json:"incident_tickets,omitempty"`
	// Ticketing_TicketEntity model
	Ticket            *TicketingTicketEntity            `json:"ticket,omitempty"`
	Impacts           []IncidentsImpactEntity           `json:"impacts,omitempty"`
	ConferenceBridges []IncidentsConferenceBridgeEntity `json:"conference_bridges,omitempty"`
	IncidentChannels  []IncidentsChannelEntity          `json:"incident_channels,omitempty"`
	// A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity
	RetroExports      []RetroExports                         `json:"retro_exports,omitempty"`
	CreatedBy         *AuthorEntity                          `json:"created_by,omitempty"`
	ContextObject     *IncidentsContextObjectEntity          `json:"context_object,omitempty"`
	TeamAssignments   []IncidentsTeamAssignmentEntity        `json:"team_assignments,omitempty"`
	Conversations     []ConversationsAPIEntitiesReference    `json:"conversations,omitempty"`
	CustomFields      []CustomFieldsFieldValue               `json:"custom_fields,omitempty"`
	FieldRequirements []IncidentEntityFieldRequirementEntity `json:"field_requirements,omitempty"`
}

func (i IncidentEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IncidentEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IncidentEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IncidentEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IncidentEntity) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *IncidentEntity) GetDiscardedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DiscardedAt
}

func (o *IncidentEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *IncidentEntity) GetCustomerImpactSummary() *string {
	if o == nil {
		return nil
	}
	return o.CustomerImpactSummary
}

func (o *IncidentEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentEntity) GetCurrentMilestone() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestone
}

func (o *IncidentEntity) GetNumber() *int {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *IncidentEntity) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *IncidentEntity) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *IncidentEntity) GetSeverityColor() *string {
	if o == nil {
		return nil
	}
	return o.SeverityColor
}

func (o *IncidentEntity) GetSeverityImpact() *string {
	if o == nil {
		return nil
	}
	return o.SeverityImpact
}

func (o *IncidentEntity) GetSeverityCondition() *string {
	if o == nil {
		return nil
	}
	return o.SeverityCondition
}

func (o *IncidentEntity) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *IncidentEntity) GetSeverityImpactObject() *SeverityMatrixImpactEntity {
	if o == nil {
		return nil
	}
	return o.SeverityImpactObject
}

func (o *IncidentEntity) GetSeverityConditionObject() *SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.SeverityConditionObject
}

func (o *IncidentEntity) GetPrivateID() *string {
	if o == nil {
		return nil
	}
	return o.PrivateID
}

func (o *IncidentEntity) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *IncidentEntity) GetMilestones() []IncidentsMilestoneEntity {
	if o == nil {
		return nil
	}
	return o.Milestones
}

func (o *IncidentEntity) GetLifecyclePhases() []IncidentsLifecyclePhaseEntity {
	if o == nil {
		return nil
	}
	return o.LifecyclePhases
}

func (o *IncidentEntity) GetLifecycleMeasurements() []IncidentsLifecycleMeasurementEntity {
	if o == nil {
		return nil
	}
	return o.LifecycleMeasurements
}

func (o *IncidentEntity) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *IncidentEntity) GetLabels() *IncidentEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *IncidentEntity) GetRoleAssignments() []IncidentsRoleAssignmentEntity {
	if o == nil {
		return nil
	}
	return o.RoleAssignments
}

func (o *IncidentEntity) GetStatusPages() []IncidentsStatusPageEntity {
	if o == nil {
		return nil
	}
	return o.StatusPages
}

func (o *IncidentEntity) GetIncidentURL() *string {
	if o == nil {
		return nil
	}
	return o.IncidentURL
}

func (o *IncidentEntity) GetPrivateStatusPageURL() *string {
	if o == nil {
		return nil
	}
	return o.PrivateStatusPageURL
}

func (o *IncidentEntity) GetOrganization() *OrganizationEntity {
	if o == nil {
		return nil
	}
	return o.Organization
}

func (o *IncidentEntity) GetCustomersImpacted() *int {
	if o == nil {
		return nil
	}
	return o.CustomersImpacted
}

func (o *IncidentEntity) GetMonetaryImpact() *int {
	if o == nil {
		return nil
	}
	return o.MonetaryImpact
}

func (o *IncidentEntity) GetMonetaryImpactCents() *int {
	if o == nil {
		return nil
	}
	return o.MonetaryImpactCents
}

func (o *IncidentEntity) GetLastUpdate() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdate
}

func (o *IncidentEntity) GetLastNote() *EventNoteEntity {
	if o == nil {
		return nil
	}
	return o.LastNote
}

func (o *IncidentEntity) GetReportID() *string {
	if o == nil {
		return nil
	}
	return o.ReportID
}

func (o *IncidentEntity) GetAiIncidentSummary() *string {
	if o == nil {
		return nil
	}
	return o.AiIncidentSummary
}

func (o *IncidentEntity) GetServices() []SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *IncidentEntity) GetEnvironments() []SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *IncidentEntity) GetFunctionalities() []SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *IncidentEntity) GetChannelName() *string {
	if o == nil {
		return nil
	}
	return o.ChannelName
}

func (o *IncidentEntity) GetChannelReference() *string {
	if o == nil {
		return nil
	}
	return o.ChannelReference
}

func (o *IncidentEntity) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *IncidentEntity) GetChannelStatus() *string {
	if o == nil {
		return nil
	}
	return o.ChannelStatus
}

func (o *IncidentEntity) GetIncidentTickets() []TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.IncidentTickets
}

func (o *IncidentEntity) GetTicket() *TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.Ticket
}

func (o *IncidentEntity) GetImpacts() []IncidentsImpactEntity {
	if o == nil {
		return nil
	}
	return o.Impacts
}

func (o *IncidentEntity) GetConferenceBridges() []IncidentsConferenceBridgeEntity {
	if o == nil {
		return nil
	}
	return o.ConferenceBridges
}

func (o *IncidentEntity) GetIncidentChannels() []IncidentsChannelEntity {
	if o == nil {
		return nil
	}
	return o.IncidentChannels
}

func (o *IncidentEntity) GetRetroExports() []RetroExports {
	if o == nil {
		return nil
	}
	return o.RetroExports
}

func (o *IncidentEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *IncidentEntity) GetContextObject() *IncidentsContextObjectEntity {
	if o == nil {
		return nil
	}
	return o.ContextObject
}

func (o *IncidentEntity) GetTeamAssignments() []IncidentsTeamAssignmentEntity {
	if o == nil {
		return nil
	}
	return o.TeamAssignments
}

func (o *IncidentEntity) GetConversations() []ConversationsAPIEntitiesReference {
	if o == nil {
		return nil
	}
	return o.Conversations
}

func (o *IncidentEntity) GetCustomFields() []CustomFieldsFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *IncidentEntity) GetFieldRequirements() []IncidentEntityFieldRequirementEntity {
	if o == nil {
		return nil
	}
	return o.FieldRequirements
}
