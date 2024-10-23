// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type IncidentsContextObjectEntity struct {
	ObjectType         *string `json:"object_type,omitempty"`
	ObjectID           *string `json:"object_id,omitempty"`
	ContextTag         *string `json:"context_tag,omitempty"`
	ContextDescription *string `json:"context_description,omitempty"`
}

func (o *IncidentsContextObjectEntity) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *IncidentsContextObjectEntity) GetObjectID() *string {
	if o == nil {
		return nil
	}
	return o.ObjectID
}

func (o *IncidentsContextObjectEntity) GetContextTag() *string {
	if o == nil {
		return nil
	}
	return o.ContextTag
}

func (o *IncidentsContextObjectEntity) GetContextDescription() *string {
	if o == nil {
		return nil
	}
	return o.ContextDescription
}
