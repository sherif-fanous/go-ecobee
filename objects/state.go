package objects

// A sensor state is a configurable trigger for a number of SensorActions.
type State struct {
	// The maximum value the sensor can generate.
	MaxValue *int `json:"maxValue,omitempty"`

	// The minimum value the sensor can generate.
	MinValue *int `json:"minValue,omitempty"`

	// Values: coolHigh, coolLow, heatHigh, heatLow, high, low, transitionCount,
	// normal.
	Type *string `json:"type,omitempty"`

	// The list of StateAction objects associated with the sensor.
	Actions []Action `json:"actions,omitempty"`
}
