// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateRetrospectiveReportReasonRequest struct {
	ReportID                                string                                             `pathParam:"style=simple,explode=false,name=report_id"`
	PostV1PostMortemsReportsReportIDReasons components.PostV1PostMortemsReportsReportIDReasons `request:"mediaType=application/json"`
}

func (o *CreateRetrospectiveReportReasonRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *CreateRetrospectiveReportReasonRequest) GetPostV1PostMortemsReportsReportIDReasons() components.PostV1PostMortemsReportsReportIDReasons {
	if o == nil {
		return components.PostV1PostMortemsReportsReportIDReasons{}
	}
	return o.PostV1PostMortemsReportsReportIDReasons
}

type CreateRetrospectiveReportReasonResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Add a new contributing factor to an incident
	PostMortemsReasonEntity *components.PostMortemsReasonEntity
}

func (o *CreateRetrospectiveReportReasonResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateRetrospectiveReportReasonResponse) GetPostMortemsReasonEntity() *components.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}
