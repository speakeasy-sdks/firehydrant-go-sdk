// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateRetrospectiveReportReasonOrderRequest struct {
	ReportID                                    string                                                 `pathParam:"style=simple,explode=false,name=report_id"`
	PutV1PostMortemsReportsReportIDReasonsOrder components.PutV1PostMortemsReportsReportIDReasonsOrder `request:"mediaType=application/json"`
}

func (o *UpdateRetrospectiveReportReasonOrderRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *UpdateRetrospectiveReportReasonOrderRequest) GetPutV1PostMortemsReportsReportIDReasonsOrder() components.PutV1PostMortemsReportsReportIDReasonsOrder {
	if o == nil {
		return components.PutV1PostMortemsReportsReportIDReasonsOrder{}
	}
	return o.PutV1PostMortemsReportsReportIDReasonsOrder
}

type UpdateRetrospectiveReportReasonOrderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Reorder a contributing factor
	PostMortemsReasonEntity *components.PostMortemsReasonEntity
}

func (o *UpdateRetrospectiveReportReasonOrderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateRetrospectiveReportReasonOrderResponse) GetPostMortemsReasonEntity() *components.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}