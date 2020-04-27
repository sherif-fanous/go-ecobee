package objects

// The event object represents a scheduled thermostat program change. All events
// have a start and end time during which the thermostat runtime settings will
// be modified. Events may not be directly modified, various Functions provide
// the capability to modify the calendar events and to modify the program. The
// event list is sorted with events ordered by whether they are currently
// running and the internal priority of each event. It is safe to take the first
// event which is running and show it as the currently running event. When the
// resume function is used, events are removed in the order they are listed
// here.
//
// Note that the start/end date/time for the event must be in thermostat time
// and are not specified in UTC.
type Event struct {
	// The type of event. Values: hold, demandResponse, sensor, switchOccupancy,
	// vacation, quickSave, today, autoAway, autoHome
	Type *string `json:"type,omitempty"`

	// The unique event name.
	Name *string `json:"name,omitempty"`

	// Whether the event is currently active or not.
	Running *bool `json:"running,omitempty"`

	// The event start date in thermostat local time.
	StartDate *string `json:"startDate,omitempty"`

	// The event start time in thermostat local time.
	StartTime *string `json:"startTime,omitempty"`

	// The event end date in thermostat local time.
	EndDate *string `json:"endDate,omitempty"`

	// The event end time in thermostat local time.
	EndTime *string `json:"endTime,omitempty"`

	// Whether there are persons occupying the property during the event.
	IsOccupied *bool `json:"isOccupied,omitempty"`

	// Whether cooling will be turned off during the event.
	IsCoolOff *bool `json:"isCoolOff,omitempty"`

	// Whether heating will be turned off during the event.
	IsHeatOff *bool `json:"isHeatOff,omitempty"`

	// The cooling absolute temperature to set.
	CoolHoldTemp *int `json:"coolHoldTemp,omitempty"`

	// The heating absolute temperature to set.
	HeatHoldTemp *int `json:"heatHoldTemp,omitempty"`

	// The fan mode during the event. Values: auto, on Default: based on current
	// climate and hvac mode.
	Fan *string `json:"fan,omitempty"`

	// The ventilator mode during the vent. Values: auto, minontime, on, off.
	Vent *string `json:"vent,omitempty"`

	// The minimum amount of time the ventilator equipment must stay on on each
	// duty cycle.
	VentilatorMinOnTime *int `json:"ventilatorMinOnTime,omitempty"`

	// Whether this event is mandatory or the end user can cancel it.
	IsOptional *bool `json:"isOptional,omitempty"`

	// Whether the event is using a relative temperature setting to the currently
	// active program climate. See the Note at the bottom of this page for more
	// information.
	IsTemperatureRelative *bool `json:"isTemperatureRelative,omitempty"`

	// The relative cool temperature adjustment.
	CoolRelativeTemp *int `json:"coolRelativeTemp,omitempty"`

	// The relative heat temperature adjustment.
	HeatRelativeTemp *int `json:"heatRelativeTemp,omitempty"`

	// Whether the event uses absolute temperatures to set the values. Default:
	// true for DRs. See the Note at the bottom of this page for more information.
	IsTemperatureAbsolute *bool `json:"isTemperatureAbsolute,omitempty"`

	// Indicates the % scheduled runtime during a Demand Response event. Valid
	// range is 0 - 100%. Default = 100, indicates no change to schedule.
	DutyCyclePercentage *int `json:"dutyCyclePercentage,omitempty"`

	// The minimum number of minutes to run the fan each hour. Range: 0-60, Default: 0
	FanMinOnTime *int `json:"fanMinOnTime,omitempty"`

	// True if this calendar event was created because of the occupied sensor.
	OccupiedSensorActive *bool `json:"occupiedSensorActive,omitempty"`

	// True if this calendar event was created because of the unoccupied sensor
	UnoccupiedSensorActive *bool `json:"unoccupiedSensorActive,omitempty"`

	// Unsupported. Future feature.
	DRRampUpTemp *int `json:"drRampUpTemp,omitempty"`

	// Unsupported. Future feature.
	DRRampUpTime *int `json:"drRampUpTime,omitempty"`

	// Unique identifier set by the server to link one or more events and alerts
	// together.
	LinkRef *string `json:"linkRef,omitempty"`

	// Used for display purposes to indicate what climate (if any) is being used
	// for the hold.
	HoldClimateRef *string `json:"holdClimateRef,omitempty"`
}
