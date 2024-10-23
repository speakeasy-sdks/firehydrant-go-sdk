// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID - Update a Signals on-call shift by ID
type PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID struct {
	// The start time of the shift in ISO8601 format.
	StartTime *string `json:"start_time,omitempty"`
	// The end time of the shift in ISO8601 format.
	EndTime *string `json:"end_time,omitempty"`
	// The ID of the user who is on-call for the shift. If not provided, the shift will be unassigned.
	UserID *string `json:"user_id,omitempty"`
	// A description of why coverage is needed for this shift. If the shift is re-assigned, this will automatically be cleared unless provided again.
	CoverageRequest *string `json:"coverage_request,omitempty"`
}

func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID) GetEndTime() *string {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID) GetCoverageRequest() *string {
	if o == nil {
		return nil
	}
	return o.CoverageRequest
}
