// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1IntegrationsAwsConnectionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// AWS account ID containing the role to be assumed
	AwsAccountID *string `queryParam:"style=form,explode=true,name=aws_account_id"`
	// ARN of the role to be assumed
	TargetArn *string `queryParam:"style=form,explode=true,name=target_arn"`
	// The external ID supplied when assuming the role
	ExternalID *string `queryParam:"style=form,explode=true,name=external_id"`
}

func (o *GetV1IntegrationsAwsConnectionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IntegrationsAwsConnectionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1IntegrationsAwsConnectionsRequest) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *GetV1IntegrationsAwsConnectionsRequest) GetTargetArn() *string {
	if o == nil {
		return nil
	}
	return o.TargetArn
}

func (o *GetV1IntegrationsAwsConnectionsRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

type GetV1IntegrationsAwsConnectionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Lists the available and configured AWS integration connections for the authenticated organization.
	IntegrationsAwsConnectionEntityPaginated *components.IntegrationsAwsConnectionEntityPaginated
}

func (o *GetV1IntegrationsAwsConnectionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IntegrationsAwsConnectionsResponse) GetIntegrationsAwsConnectionEntityPaginated() *components.IntegrationsAwsConnectionEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsConnectionEntityPaginated
}
