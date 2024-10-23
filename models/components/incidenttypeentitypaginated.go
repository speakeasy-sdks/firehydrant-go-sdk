// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IncidentTypeEntityPaginated model
type IncidentTypeEntityPaginated struct {
	Data       []IncidentTypeEntity `json:"data,omitempty"`
	Pagination *PaginationEntity    `json:"pagination,omitempty"`
}

func (o *IncidentTypeEntityPaginated) GetData() []IncidentTypeEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IncidentTypeEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
