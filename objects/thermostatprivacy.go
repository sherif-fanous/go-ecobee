package objects

// This object is currently undocumented by ecobee.
type ThermostatPrivacy struct {
	AllowContractorAlerts                *bool `json:"allowContractorAlerts,omitempty"`
	AllowContractorHVACReports           *bool `json:"allowContractorHvacReports,omitempty"`
	AllowContractorHeatPumpConfiguration *bool `json:"allowContractorHeatPumpConfiguration,omitempty"`
}
