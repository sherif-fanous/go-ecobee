package objects

// This object is currently undocumented by ecobee.
type ThermostatOEMCfg struct {
	SerialNumber *string                `json:"serialNumber,omitempty"`
	CCfg         map[string]interface{} `json:"cCfg,omitempty"`
}
