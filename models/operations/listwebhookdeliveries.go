// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListWebhookDeliveriesRequest struct {
	// ID of a webhook
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook_id"`
}

func (o *ListWebhookDeliveriesRequest) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

type ListWebhookDeliveriesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListWebhookDeliveriesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
