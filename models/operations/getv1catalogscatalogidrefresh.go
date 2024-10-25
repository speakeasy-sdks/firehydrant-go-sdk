// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1CatalogsCatalogIDRefreshRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *GetV1CatalogsCatalogIDRefreshRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type GetV1CatalogsCatalogIDRefreshResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1CatalogsCatalogIDRefreshResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
