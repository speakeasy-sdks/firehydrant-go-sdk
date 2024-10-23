// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1PostMortemsReportsReportIDReasonsRequest struct {
	ReportID                                string                                             `pathParam:"style=simple,explode=false,name=report_id"`
	PostV1PostMortemsReportsReportIDReasons components.PostV1PostMortemsReportsReportIDReasons `request:"mediaType=application/json"`
}

func (o *PostV1PostMortemsReportsReportIDReasonsRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *PostV1PostMortemsReportsReportIDReasonsRequest) GetPostV1PostMortemsReportsReportIDReasons() components.PostV1PostMortemsReportsReportIDReasons {
	if o == nil {
		return components.PostV1PostMortemsReportsReportIDReasons{}
	}
	return o.PostV1PostMortemsReportsReportIDReasons
}

type PostV1PostMortemsReportsReportIDReasonsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Add a new contributing factor to an incident
	PostMortemsReasonEntity *components.PostMortemsReasonEntity
}

func (o *PostV1PostMortemsReportsReportIDReasonsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1PostMortemsReportsReportIDReasonsResponse) GetPostMortemsReasonEntity() *components.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}