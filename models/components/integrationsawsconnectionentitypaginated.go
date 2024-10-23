// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IntegrationsAwsConnectionEntityPaginated - Integrations_Aws_ConnectionEntityPaginated model
type IntegrationsAwsConnectionEntityPaginated struct {
	Data       []IntegrationsAwsConnectionEntity `json:"data,omitempty"`
	Pagination *PaginationEntity                 `json:"pagination,omitempty"`
}

func (o *IntegrationsAwsConnectionEntityPaginated) GetData() []IntegrationsAwsConnectionEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IntegrationsAwsConnectionEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
