// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdatePostMortemReportRequest struct {
	ReportID                          string                                       `pathParam:"style=simple,explode=false,name=report_id"`
	PatchV1PostMortemsReportsReportID components.PatchV1PostMortemsReportsReportID `request:"mediaType=application/json"`
}

func (o *UpdatePostMortemReportRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *UpdatePostMortemReportRequest) GetPatchV1PostMortemsReportsReportID() components.PatchV1PostMortemsReportsReportID {
	if o == nil {
		return components.PatchV1PostMortemsReportsReportID{}
	}
	return o.PatchV1PostMortemsReportsReportID
}

type UpdatePostMortemReportResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a report
	PostMortemsPostMortemReportEntity *components.PostMortemsPostMortemReportEntity
}

func (o *UpdatePostMortemReportResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePostMortemReportResponse) GetPostMortemsPostMortemReportEntity() *components.PostMortemsPostMortemReportEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsPostMortemReportEntity
}