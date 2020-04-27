package objects

// The Sensor object represents a sensor connected to the thermostat. Sensors
// may not be modified using the API, however some configuration may occur
// through the web portal.
type Sensor struct {
	// The sensor name
	Name *string `json:"name,omitempty"`

	// The sensor manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// The sensor model
	Model *string `json:"model,omitempty"`

	// The thermostat zone the sensor is associated with
	Zone *int `json:"zone,omitempty"`

	// The unique sensor identifier
	SensorID *int `json:"sensorId,omitempty"`

	// The type of sensor. Values: adc, co2, dryCOntact, humidity, temperature,
	// unknown.
	Type *string `json:"type,omitempty"`

	// The sensor usage type. Values: dischargeAir, indoor, monitor, outdoor.
	Usage *string `json:"usage,omitempty"`

	// The number of bits the adc has been configured to use.
	NumberOfBits *int `json:"numberOfBits,omitempty"`

	// Value of the bconstant value used in temperature sensors
	Bconstant *int `json:"bconstant,omitempty"`

	// The sensor thermistor value, ie. 10K thermistor=10000, 2.5K thermistor=2500
	ThermistorSize *int `json:"thermistorSize,omitempty"`

	// The user adjustable temperature compensation applied to the temperature reading.
	TempCorrection *int `json:"tempCorrection,omitempty"`

	// The sensor thermistor gain value.
	Gain *int `json:"gain,omitempty"`

	// The sensor thermistor max voltage in Volts, 5=5V, 10=10V.
	MaxVoltage *int `json:"maxVoltage,omitempty"`

	// The multiplier value used in sensors (1000x value).
	Multiplier *int `json:"multiplier,omitempty"`

	// A list of SensorState objects
	States []State `json:"states,omitempty"`
}
