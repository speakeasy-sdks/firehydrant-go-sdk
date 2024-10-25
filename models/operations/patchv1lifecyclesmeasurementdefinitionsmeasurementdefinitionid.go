// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody struct {
	Name                *string `form:"name=name"`
	Slug                *string `form:"name=slug"`
	Description         *string `form:"name=description"`
	StartsAtMilestoneID *string `form:"name=starts_at_milestone_id"`
	EndsAtMilestoneID   *string `form:"name=ends_at_milestone_id"`
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody) GetStartsAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.StartsAtMilestoneID
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody) GetEndsAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.EndsAtMilestoneID
}

type PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequest struct {
	MeasurementDefinitionID string                                                                     `pathParam:"style=simple,explode=false,name=measurement_definition_id"`
	RequestBody             *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequest) GetMeasurementDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.MeasurementDefinitionID
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequest) GetRequestBody() *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
