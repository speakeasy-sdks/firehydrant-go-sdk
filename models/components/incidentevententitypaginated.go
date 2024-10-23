// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IncidentEventEntityPaginated model
type IncidentEventEntityPaginated struct {
	Data       []IncidentEventEntity `json:"data,omitempty"`
	Pagination *PaginationEntity     `json:"pagination,omitempty"`
}

func (o *IncidentEventEntityPaginated) GetData() []IncidentEventEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IncidentEventEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
