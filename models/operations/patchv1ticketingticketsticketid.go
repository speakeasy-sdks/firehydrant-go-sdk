// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1TicketingTicketsTicketIDRequest struct {
	TicketID                        string                                     `pathParam:"style=simple,explode=false,name=ticket_id"`
	PatchV1TicketingTicketsTicketID components.PatchV1TicketingTicketsTicketID `request:"mediaType=application/json"`
}

func (o *PatchV1TicketingTicketsTicketIDRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

func (o *PatchV1TicketingTicketsTicketIDRequest) GetPatchV1TicketingTicketsTicketID() components.PatchV1TicketingTicketsTicketID {
	if o == nil {
		return components.PatchV1TicketingTicketsTicketID{}
	}
	return o.PatchV1TicketingTicketsTicketID
}

type PatchV1TicketingTicketsTicketIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a ticket's attributes
	TicketingTicketEntity *components.TicketingTicketEntity
}

func (o *PatchV1TicketingTicketsTicketIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1TicketingTicketsTicketIDResponse) GetTicketingTicketEntity() *components.TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.TicketingTicketEntity
}
