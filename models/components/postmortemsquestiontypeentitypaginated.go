// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PostMortemsQuestionTypeEntityPaginated - PostMortems_QuestionTypeEntityPaginated model
type PostMortemsQuestionTypeEntityPaginated struct {
	Data       []PostMortemsQuestionTypeEntity `json:"data,omitempty"`
	Pagination *PaginationEntity               `json:"pagination,omitempty"`
}

func (o *PostMortemsQuestionTypeEntityPaginated) GetData() []PostMortemsQuestionTypeEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PostMortemsQuestionTypeEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
