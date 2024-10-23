// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ScheduledMaintenancesImpactEntity struct {
	ID     *string         `json:"id,omitempty"`
	Type   *string         `json:"type,omitempty"`
	Impact *SuccinctEntity `json:"impact,omitempty"`
	// SeverityMatrix_ConditionEntity model
	Condition *SeverityMatrixConditionEntity `json:"condition,omitempty"`
}

func (o *ScheduledMaintenancesImpactEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ScheduledMaintenancesImpactEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ScheduledMaintenancesImpactEntity) GetImpact() *SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Impact
}

func (o *ScheduledMaintenancesImpactEntity) GetCondition() *SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.Condition
}
