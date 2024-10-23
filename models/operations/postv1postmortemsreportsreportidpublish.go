// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1PostMortemsReportsReportIDPublishRequest struct {
	ReportID                                string                                             `pathParam:"style=simple,explode=false,name=report_id"`
	PostV1PostMortemsReportsReportIDPublish components.PostV1PostMortemsReportsReportIDPublish `request:"mediaType=application/json"`
}

func (o *PostV1PostMortemsReportsReportIDPublishRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *PostV1PostMortemsReportsReportIDPublishRequest) GetPostV1PostMortemsReportsReportIDPublish() components.PostV1PostMortemsReportsReportIDPublish {
	if o == nil {
		return components.PostV1PostMortemsReportsReportIDPublish{}
	}
	return o.PostV1PostMortemsReportsReportIDPublish
}

type PostV1PostMortemsReportsReportIDPublishResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Marks an incident retrospective as published and emails all of the participants in the report the summary
	PostMortemsPostMortemReportEntity *components.PostMortemsPostMortemReportEntity
}

func (o *PostV1PostMortemsReportsReportIDPublishResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1PostMortemsReportsReportIDPublishResponse) GetPostMortemsPostMortemReportEntity() *components.PostMortemsPostMortemReportEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsPostMortemReportEntity
}
