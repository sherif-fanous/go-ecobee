package objects

// A climate defines a thermostat settings template which is then applied to
// individual period cells of the Schedule. The result is that if you modify the
// Climate, all schedule cells which reference that Climate will automatically
// be changed.
//
// When adding a Climate it is optional whether you reference the new Climate in
// the schedule cells in the same request or not. However, when deleting a
// Climate (by omitting that entire Climate object from the POST request) it can
// not be be deleted if it is still referenced in the schedule cells.
type Climate struct {
	// The unique climate name. The name may be changed without affecting the
	// program integrity so long as uniqueness is maintained.
	Name *string `json:"name,omitempty"`

	// The unique climate identifier. Changing the identifier is not possible and
	// it is generated on the server for each climate. If this value is not
	// supplied a new climate will be created. For the default climates and
	// existing user created climates the climateRef should be supplied - see note
	// above.
	ClimateRef *string `json:"climateRef,omitempty"`

	// A flag indicating whether the property is occupied by persons during this
	// climate
	IsOccupied *bool `json:"isOccupied,omitempty"`

	// A flag indicating whether ecobee optimized climate settings are used by this
	// climate.
	IsOptimized *bool `json:"isOptimized,omitempty"`

	// The cooling fan mode. Default: on. Values: auto, on.
	CoolFan *string `json:"coolFan,omitempty"`

	// The heating fan mode. Default: on. Values: auto, on.
	HeatFan *string `json:"heatFan,omitempty"`

	// The ventilator mode. Default: off. Values: auto, minontime, on, off.
	Vent *string `json:"vent,omitempty"`

	// The minimum time, in minutes, to run the ventilator each hour.
	VentilatorMinOnTime *int `json:"ventilatorMinOnTime,omitempty"`

	// The climate owner. Default: system. Values: adHoc, demandResponse,
	// quickSave, sensorAction, switchOccupancy, system, template, user.
	Owner *string `json:"owner,omitempty"`

	// The type of climate. Default: program. Values: calendarEvent, program.
	Type *string `json:"type,omitempty"`

	// The integer conversion of the HEX color value used to display this climate
	// on the thermostat and on the web portal.
	Colour *int `json:"colour,omitempty"`

	// The cool temperature for this climate.
	CoolTemp *int `json:"coolTemp,omitempty"`

	// The heat temperature for this climate.
	HeatTemp *int `json:"heatTemp,omitempty"`

	// The list of sensors in use for the specific climate. The sensors listed here
	// are used for temperature averaging within that climate. Only the sensorId
	// and name are listed in the climate.
	Sensors []RemoteSensor `json:"sensors,omitempty"`
}
