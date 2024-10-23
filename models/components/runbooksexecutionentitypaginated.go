// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RunbooksExecutionEntityPaginated - Runbooks_ExecutionEntityPaginated model
type RunbooksExecutionEntityPaginated struct {
	Data       []RunbooksExecutionEntity `json:"data,omitempty"`
	Pagination *PaginationEntity         `json:"pagination,omitempty"`
}

func (o *RunbooksExecutionEntityPaginated) GetData() []RunbooksExecutionEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *RunbooksExecutionEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}