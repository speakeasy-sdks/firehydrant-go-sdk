// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PostV1PostMortemsReportsReportIDReasons - Add a new contributing factor to an incident
type PostV1PostMortemsReportsReportIDReasons struct {
	Summary string `json:"summary"`
}

func (o *PostV1PostMortemsReportsReportIDReasons) GetSummary() string {
	if o == nil {
		return ""
	}
	return o.Summary
}
