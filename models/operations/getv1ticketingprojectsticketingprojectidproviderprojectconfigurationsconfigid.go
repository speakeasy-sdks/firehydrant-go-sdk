// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
	ConfigID           string `pathParam:"style=simple,explode=false,name=config_id"`
}

func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

type GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve configuration for a ticketing project
	TicketingProjectConfigEntity *components.TicketingProjectConfigEntity
}

func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse) GetTicketingProjectConfigEntity() *components.TicketingProjectConfigEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectConfigEntity
}
