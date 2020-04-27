package objects

// The Version object contains version information about the thermostat.
type Version struct {
	// The thermostat firmware version number.
	ThermostatFirmwareVersion *string `json:"thermostatFirmwareVersion,omitempty"`
}
