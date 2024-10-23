// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// SeverityEntityPaginated model
type SeverityEntityPaginated struct {
	Data       []SeverityEntity  `json:"data,omitempty"`
	Pagination *PaginationEntity `json:"pagination,omitempty"`
}

func (o *SeverityEntityPaginated) GetData() []SeverityEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *SeverityEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
