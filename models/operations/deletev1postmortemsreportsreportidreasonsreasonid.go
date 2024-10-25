// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1PostMortemsReportsReportIDReasonsReasonIDRequest struct {
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
	ReasonID string `pathParam:"style=simple,explode=false,name=reason_id"`
}

func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDRequest) GetReasonID() string {
	if o == nil {
		return ""
	}
	return o.ReasonID
}

type DeleteV1PostMortemsReportsReportIDReasonsReasonIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a contributing factor
	PostMortemsReasonEntity *components.PostMortemsReasonEntity
}

func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1PostMortemsReportsReportIDReasonsReasonIDResponse) GetPostMortemsReasonEntity() *components.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}
