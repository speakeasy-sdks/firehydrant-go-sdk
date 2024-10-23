// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RunbooksElementEntity struct {
	ID            *string                             `json:"id,omitempty"`
	Type          *string                             `json:"type,omitempty"`
	Markdown      *RunbooksElementMarkdownEntity      `json:"markdown,omitempty"`
	Textarea      *RunbooksElementTextareaEntity      `json:"textarea,omitempty"`
	Input         *RunbooksElementInputEntity         `json:"input,omitempty"`
	PlainText     *RunbooksElementMarkdownEntity      `json:"plain_text,omitempty"`
	DynamicSelect *RunbooksElementDynamicSelectEntity `json:"dynamic_select,omitempty"`
}

func (o *RunbooksElementEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RunbooksElementEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *RunbooksElementEntity) GetMarkdown() *RunbooksElementMarkdownEntity {
	if o == nil {
		return nil
	}
	return o.Markdown
}

func (o *RunbooksElementEntity) GetTextarea() *RunbooksElementTextareaEntity {
	if o == nil {
		return nil
	}
	return o.Textarea
}

func (o *RunbooksElementEntity) GetInput() *RunbooksElementInputEntity {
	if o == nil {
		return nil
	}
	return o.Input
}

func (o *RunbooksElementEntity) GetPlainText() *RunbooksElementMarkdownEntity {
	if o == nil {
		return nil
	}
	return o.PlainText
}

func (o *RunbooksElementEntity) GetDynamicSelect() *RunbooksElementDynamicSelectEntity {
	if o == nil {
		return nil
	}
	return o.DynamicSelect
}