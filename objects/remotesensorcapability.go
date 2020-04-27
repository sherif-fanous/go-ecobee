package objects

// The RemoteSensorCapability object represents the specific capability of a
// sensor connected to the thermostat.
//
// For the occupancy type capability the data will only show computed occupancy,
// as does the thermostat.
//
// Definition - For a given sensor, computed occupancy means a sensor is
// occupied if any motion was detected in the past 30 minutes.
type RemoteSensorCapability struct {
	// The unique sensor capability identifier. For example: 1
	ID *string `json:"id,omitempty"`

	// The type of sensor capability. Values: adc, co2, dryContact, humidity,
	// temperature, occupancy, unknown.
	Type *string `json:"type,omitempty"`

	// The data value for this capability, always a String. Temperature values are
	// expressed as degrees Fahrenheit, multiplied by 10. For example, a
	// temperature of 72F would be returned as the value "720". Occupancy values
	// are "true" or "false". Humidity is expressed as a % value such as "45".
	// Unknown values are returned as "unknown".
	Value *string `json:"value,omitempty"`
}
