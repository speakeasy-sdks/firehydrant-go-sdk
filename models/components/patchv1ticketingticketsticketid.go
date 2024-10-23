// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PatchV1TicketingTicketsTicketID - Update a ticket's attributes
type PatchV1TicketingTicketsTicketID struct {
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	State       *string `json:"state,omitempty"`
	Type        *string `json:"type,omitempty"`
	PriorityID  *string `json:"priority_id,omitempty"`
	// List of tags for the ticket
	TagList []string `json:"tag_list,omitempty"`
}

func (o *PatchV1TicketingTicketsTicketID) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PatchV1TicketingTicketsTicketID) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1TicketingTicketsTicketID) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *PatchV1TicketingTicketsTicketID) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PatchV1TicketingTicketsTicketID) GetPriorityID() *string {
	if o == nil {
		return nil
	}
	return o.PriorityID
}

func (o *PatchV1TicketingTicketsTicketID) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}
