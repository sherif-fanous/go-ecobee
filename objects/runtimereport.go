package objects

// The runtime report contains the historical runtime information for a thermostat.
//
// The date/times returned in the runtime report are always in thermostat time
type RuntimeReport struct {
	// The thermostat identifier for the report.
	ThermostatIdentifier *string `json:"thermostatIdentifier,omitempty"`

	// The number of report rows in this report
	RowCount *int `json:"rowCount,omitempty"`

	// A list of CSV report strings based on the columns requested.
	RowList []string `json:"rowList,omitempty"`
}
