// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IncidentsAlertEntity - Incidents_AlertEntity model
type IncidentsAlertEntity struct {
	ID *string `json:"id,omitempty"`
	// Alerts_AlertEntity model
	Alert *AlertsAlertEntity `json:"alert,omitempty"`
	// whether or not this is the primary alert for this incident
	Primary *bool `json:"primary,omitempty"`
}

func (o *IncidentsAlertEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentsAlertEntity) GetAlert() *AlertsAlertEntity {
	if o == nil {
		return nil
	}
	return o.Alert
}

func (o *IncidentsAlertEntity) GetPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.Primary
}
