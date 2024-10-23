// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TicketingProjectFieldMapEntity - Ticketing_ProjectFieldMapEntity model
type TicketingProjectFieldMapEntity struct {
	ID                 *string                              `json:"id,omitempty"`
	ConnectionID       *string                              `json:"connection_id,omitempty"`
	ConnectionType     *string                              `json:"connection_type,omitempty"`
	TicketingProjectID *string                              `json:"ticketing_project_id,omitempty"`
	Body               []TicketingProjectFieldMapBodyEntity `json:"body,omitempty"`
}

func (o *TicketingProjectFieldMapEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TicketingProjectFieldMapEntity) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *TicketingProjectFieldMapEntity) GetConnectionType() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionType
}

func (o *TicketingProjectFieldMapEntity) GetTicketingProjectID() *string {
	if o == nil {
		return nil
	}
	return o.TicketingProjectID
}

func (o *TicketingProjectFieldMapEntity) GetBody() []TicketingProjectFieldMapBodyEntity {
	if o == nil {
		return nil
	}
	return o.Body
}