package objects

// Represents a device attached to the thermostat. Devices may not be modified
// remotely, all changes must occur on the thermostat.
type Device struct {
	// A unique ID for the device
	DeviceID *int `json:"deviceId,omitempty"`

	// The user supplied device name
	Name *string `json:"name,omitempty"`

	// The list of Sensor Objects associated with the device.
	Sensors []Sensor `json:"sensors,omitempty"`

	// Ths list of Output Objects associated with the device
	Outputs []Output `json:"outputs,omitempty"`
}
