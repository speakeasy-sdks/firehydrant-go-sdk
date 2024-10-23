// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// SeverityMatrixSeverityMatrixEntity - SeverityMatrix_SeverityMatrixEntity model
type SeverityMatrixSeverityMatrixEntity struct {
	Matrix     []SeverityMatrixItemEntity      `json:"matrix,omitempty"`
	Impacts    []SeverityMatrixImpactEntity    `json:"impacts,omitempty"`
	Conditions []SeverityMatrixConditionEntity `json:"conditions,omitempty"`
}

func (o *SeverityMatrixSeverityMatrixEntity) GetMatrix() []SeverityMatrixItemEntity {
	if o == nil {
		return nil
	}
	return o.Matrix
}

func (o *SeverityMatrixSeverityMatrixEntity) GetImpacts() []SeverityMatrixImpactEntity {
	if o == nil {
		return nil
	}
	return o.Impacts
}

func (o *SeverityMatrixSeverityMatrixEntity) GetConditions() []SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.Conditions
}
