// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1IntegrationsAwsCloudtrailBatchesIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetV1IntegrationsAwsCloudtrailBatchesIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single CloudTrail batch.
	IntegrationsAwsCloudtrailBatchEntity *components.IntegrationsAwsCloudtrailBatchEntity
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesIDResponse) GetIntegrationsAwsCloudtrailBatchEntity() *components.IntegrationsAwsCloudtrailBatchEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsCloudtrailBatchEntity
}