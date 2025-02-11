// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteNuncSubscriptionRequest struct {
	UnsubscribeToken string `pathParam:"style=simple,explode=false,name=unsubscribe_token"`
}

func (o *DeleteNuncSubscriptionRequest) GetUnsubscribeToken() string {
	if o == nil {
		return ""
	}
	return o.UnsubscribeToken
}

type DeleteNuncSubscriptionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Unsubscribe from status page updates
	NuncNuncSubscription *components.NuncNuncSubscription
}

func (o *DeleteNuncSubscriptionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteNuncSubscriptionResponse) GetNuncNuncSubscription() *components.NuncNuncSubscription {
	if o == nil {
		return nil
	}
	return o.NuncNuncSubscription
}
