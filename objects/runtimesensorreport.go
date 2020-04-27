package objects

// The Runtime Sensor Report object is a container for the sensor metadata and
// runtime data for a single thermostat.
type RuntimeSensorReport struct {
	// The thermostat identifier for the report.
	ThermostatIdentifier *string `json:"thermostatIdentifier,omitempty"`

	// The list of sensor metadata configured in the thermostat.
	Sensors []RuntimeSensorMetadata `json:"sensors,omitempty"`

	// The list of column names returned in the data property. The sensor data
	// column names match the sensorId within the sensor metadata. The first two
	// columns are the date and time, the following are the defined sensorIds.
	Columns []string `json:"columns,omitempty"`

	// The list of CSV rows containing the column data as defined in the columns
	// property.
	Data []string `json:"data,omitempty"`
}
