// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PongEntity model
type PongEntity struct {
	Response     *string             `json:"response,omitempty"`
	Actor        *ActorEntity        `json:"actor,omitempty"`
	Organization *OrganizationEntity `json:"organization,omitempty"`
}

func (o *PongEntity) GetResponse() *string {
	if o == nil {
		return nil
	}
	return o.Response
}

func (o *PongEntity) GetActor() *ActorEntity {
	if o == nil {
		return nil
	}
	return o.Actor
}

func (o *PongEntity) GetOrganization() *OrganizationEntity {
	if o == nil {
		return nil
	}
	return o.Organization
}
