// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RunbooksActionsEntityPaginated - Runbooks_ActionsEntityPaginated model
type RunbooksActionsEntityPaginated struct {
	Data       []RunbooksActionsEntity `json:"data,omitempty"`
	Pagination *PaginationEntity       `json:"pagination,omitempty"`
}

func (o *RunbooksActionsEntityPaginated) GetData() []RunbooksActionsEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *RunbooksActionsEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}