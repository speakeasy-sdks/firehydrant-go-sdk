// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type OrganizationEntity struct {
	Name *string `json:"name,omitempty"`
	ID   *string `json:"id,omitempty"`
}

func (o *OrganizationEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OrganizationEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}