// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TagEntityPaginated model
type TagEntityPaginated struct {
	Data       []TagEntity       `json:"data,omitempty"`
	Pagination *PaginationEntity `json:"pagination,omitempty"`
}

func (o *TagEntityPaginated) GetData() []TagEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TagEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
