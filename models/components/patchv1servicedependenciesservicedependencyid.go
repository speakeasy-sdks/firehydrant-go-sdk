// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PatchV1ServiceDependenciesServiceDependencyID - Update the notes of the service dependency
type PatchV1ServiceDependenciesServiceDependencyID struct {
	// A note to describe the service dependency relationship
	Notes *string `json:"notes,omitempty"`
}

func (o *PatchV1ServiceDependenciesServiceDependencyID) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}