// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RunbooksExecutionEntityExecutedForEntity struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *RunbooksExecutionEntityExecutedForEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RunbooksExecutionEntityExecutedForEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RunbooksExecutionEntityExecutedForEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
