package objects

// The thermostat meter report is a container for the meter data of a single
// thermostat. It contains meter data for each type of meter requested and its
// associated data.
type MeterReport struct {
	// The thermostat identifier this report is for.
	ThermostatIdentifier *string `json:"thermostatIdentifier,omitempty"`

	// The list of meter data for the meters requested. If the thermostat has no
	// meter, the object for that meter will not be included in the list. A
	// thermostat with no meters will have this list empty.
	MeterList []MeterReportData `json:"meterList,omitempty"`
}
