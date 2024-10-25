// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1TicketingPrioritiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all ticketing priorities available to the organization
	TicketingPriorityEntity *components.TicketingPriorityEntity
}

func (o *GetV1TicketingPrioritiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1TicketingPrioritiesResponse) GetTicketingPriorityEntity() *components.TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.TicketingPriorityEntity
}
