// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ChangeTypeEntity struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *ChangeTypeEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeTypeEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}