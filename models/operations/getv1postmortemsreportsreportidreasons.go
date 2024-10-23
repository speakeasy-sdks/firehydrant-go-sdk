// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1PostMortemsReportsReportIDReasonsRequest struct {
	ReportID string `pathParam:"style=simple,explode=false,name=report_id"`
	Page     *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage  *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1PostMortemsReportsReportIDReasonsRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *GetV1PostMortemsReportsReportIDReasonsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1PostMortemsReportsReportIDReasonsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1PostMortemsReportsReportIDReasonsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all contributing factors to an incident
	PostMortemsReasonEntityPaginated *components.PostMortemsReasonEntityPaginated
}

func (o *GetV1PostMortemsReportsReportIDReasonsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1PostMortemsReportsReportIDReasonsResponse) GetPostMortemsReasonEntityPaginated() *components.PostMortemsReasonEntityPaginated {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntityPaginated
}
