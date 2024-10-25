// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1NuncConnectionsNuncConnectionIDSubscribersRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	// A list of subscriber IDs to unsubscribe.
	SubscriberIds string `queryParam:"style=form,explode=true,name=subscriber_ids"`
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDSubscribersRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDSubscribersRequest) GetSubscriberIds() string {
	if o == nil {
		return ""
	}
	return o.SubscriberIds
}

type DeleteV1NuncConnectionsNuncConnectionIDSubscribersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Unsubscribes one or more status page subscribers.
	NuncEmailSubscribersEntity *components.NuncEmailSubscribersEntity
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDSubscribersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDSubscribersResponse) GetNuncEmailSubscribersEntity() *components.NuncEmailSubscribersEntity {
	if o == nil {
		return nil
	}
	return o.NuncEmailSubscribersEntity
}
