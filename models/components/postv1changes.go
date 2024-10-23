// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PostV1Changes - Create a new change entry
type PostV1Changes struct {
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	// A labels hash of keys and values
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *PostV1Changes) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PostV1Changes) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1Changes) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
