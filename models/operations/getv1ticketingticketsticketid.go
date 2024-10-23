// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1TicketingTicketsTicketIDRequest struct {
	TicketID string `pathParam:"style=simple,explode=false,name=ticket_id"`
}

func (o *GetV1TicketingTicketsTicketIDRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

type GetV1TicketingTicketsTicketIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieves a single ticket by ID
	TicketingTicketEntity *components.TicketingTicketEntity
}

func (o *GetV1TicketingTicketsTicketIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1TicketingTicketsTicketIDResponse) GetTicketingTicketEntity() *components.TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.TicketingTicketEntity
}