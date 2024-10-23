// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"openapi/models/components"
)

type GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType string

const (
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeChangeEvents          GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "change_events"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeIncidents             GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "incidents"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeServices              GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "services"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeScheduledMaintenances GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "scheduled_maintenances"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeTicketTasks           GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "ticket_tasks"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeTicketFollowUps       GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "ticket_follow_ups"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeAnalytics             GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "analytics"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeImpactAnalytics       GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "impact_analytics"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeAlerts                GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "alerts"
	GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeIncidentEvents        GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType = "incident_events"
)

func (e GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType) ToPointer() *GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType {
	return &e
}
func (e *GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "change_events":
		fallthrough
	case "incidents":
		fallthrough
	case "services":
		fallthrough
	case "scheduled_maintenances":
		fallthrough
	case "ticket_tasks":
		fallthrough
	case "ticket_follow_ups":
		fallthrough
	case "analytics":
		fallthrough
	case "impact_analytics":
		fallthrough
	case "alerts":
		fallthrough
	case "incident_events":
		*e = GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType: %v", v)
	}
}

type GetV1SavedSearchesResourceTypeSavedSearchIDRequest struct {
	ResourceType  GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType `pathParam:"style=simple,explode=false,name=resource_type"`
	SavedSearchID string                                                           `pathParam:"style=simple,explode=false,name=saved_search_id"`
}

func (o *GetV1SavedSearchesResourceTypeSavedSearchIDRequest) GetResourceType() GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType {
	if o == nil {
		return GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType("")
	}
	return o.ResourceType
}

func (o *GetV1SavedSearchesResourceTypeSavedSearchIDRequest) GetSavedSearchID() string {
	if o == nil {
		return ""
	}
	return o.SavedSearchID
}

type GetV1SavedSearchesResourceTypeSavedSearchIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a specific save search
	SavedSearchEntity *components.SavedSearchEntity
}

func (o *GetV1SavedSearchesResourceTypeSavedSearchIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1SavedSearchesResourceTypeSavedSearchIDResponse) GetSavedSearchEntity() *components.SavedSearchEntity {
	if o == nil {
		return nil
	}
	return o.SavedSearchEntity
}