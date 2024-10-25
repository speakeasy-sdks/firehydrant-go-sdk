// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
	ConfigID           string `pathParam:"style=simple,explode=false,name=config_id"`
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

type DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archive configuration for a ticketing project
	TicketingProjectConfigEntity *components.TicketingProjectConfigEntity
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse) GetTicketingProjectConfigEntity() *components.TicketingProjectConfigEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectConfigEntity
}
