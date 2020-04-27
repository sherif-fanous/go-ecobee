package objects

// The RemoteSensor object represents a sensor connected to the thermostat.
//
// The remote sensor data will only show computed occupancy, as does the
// thermostat. Definition - For a given sensor, computed occupancy means a
// sensor is occupied if any motion was detected in the past 30 minutes.
//
// RemoteSensor data changes trigger the runtimeRevision to be updated. The data
// updates are sent at an interval of 3 mins maximum. This means that you should
// not poll quicker than once every 3 mins for revision changes.
//
// Note: remote sensors may only be modified using the UpdateSensor
// (https://www.ecobee.com/home/developer/api/documentation/v1/functions/UpdateSensor.shtml)
// function.
type RemoteSensor struct {
	// The unique sensor identifier. It is composed of deviceName + deviceId
	// separated by colons, for example: rs:100
	ID *string `json:"id,omitempty"`

	// The user assigned sensor name.
	Name *string `json:"name,omitempty"`

	// The type of sensor. Values: thermostat, ecobee3_remote_sensor,
	// monitor_sensor, control_sensor.
	Type *string `json:"type,omitempty"`

	// The unique 4-digit alphanumeric sensor code. For ecobee3 remote sensors this
	// corresponds to the code found on the back of the physical sensor.
	Code *string `json:"code,omitempty"`

	// This flag indicates whether the remote sensor is currently in use by a
	// comfort setting. See Climate for more information.
	InUse *bool `json:"inUse,omitempty"`

	// The list of remoteSensorCapability objects for the remote sensor.
	Capability []RemoteSensorCapability `json:"capability,omitempty"`
}
