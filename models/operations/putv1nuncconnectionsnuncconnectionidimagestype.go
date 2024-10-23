// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PutV1NuncConnectionsNuncConnectionIDImagesTypeFile struct {
	FileName string `multipartForm:"name=file"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeFile) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type PutV1NuncConnectionsNuncConnectionIDImagesTypeRequestBody struct {
	File *PutV1NuncConnectionsNuncConnectionIDImagesTypeFile `multipartForm:"file"`
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeRequestBody) GetFile() *PutV1NuncConnectionsNuncConnectionIDImagesTypeFile {
	if o == nil {
		return nil
	}
	return o.File
}

type PutV1NuncConnectionsNuncConnectionIDImagesTypeRequest struct {
	NuncConnectionID string                                                     `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	Type             string                                                     `pathParam:"style=simple,explode=false,name=type"`
	RequestBody      *PutV1NuncConnectionsNuncConnectionIDImagesTypeRequestBody `request:"mediaType=multipart/form-data"`
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeRequest) GetRequestBody() *PutV1NuncConnectionsNuncConnectionIDImagesTypeRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PutV1NuncConnectionsNuncConnectionIDImagesTypeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Add or replace an image attached to a FireHydrant status page
	NuncConnectionEntity *components.NuncConnectionEntity
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeResponse) GetNuncConnectionEntity() *components.NuncConnectionEntity {
	if o == nil {
		return nil
	}
	return o.NuncConnectionEntity
}