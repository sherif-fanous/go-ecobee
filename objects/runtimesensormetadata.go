package objects

// The metadata description of a sensor as related to the sensor runtime data.
// This is a reflection of the thermostat.device[].sensor[] information.
//
// Note that the sensor metadata reflects the sensor configuration, type, usage,
// etc. at the time of the first reading taken as specified by the runtime
// report start date and interval. This information may not match the sensor
// information currently in the thermostat if it had changed after the start of
// the report period. Likewise, a change to the sensor configuration during the
// report period may result in values which do not match the expected values of
// the sensor configuration at the beginning of the report.
type RuntimeSensorMetadata struct {
	// The unique sensor identifier. It is composed of deviceName + deviceId +
	// sensorId (from thermostat.device[].sensor[]) separated by colons. This value
	// corresponds to the column name for the sensor reading values.
	SensorID *string `json:"sensorId,omitempty"`

	// The user assigned sensor name.
	SensorName *string `json:"sensorName,omitempty"`

	// The type of sensor. See Sensor Types. Values: co2, ctclamp, dryContact,
	// humidity, plug, temperature
	SensorType *string `json:"sensorType,omitempty"`

	// The usage configured for the sensor. Values: dischargeAir, indoor, monitor,
	// outdoor
	SensorUsage *string `json:"sensorUsage,omitempty"`
}
