package objects

// Contains the reading data from a single meter.
//
// See Meter Report
// (https://www.ecobee.com/home/developer/api/documentation/v1/operations/get-meter-report.shtml)
// for information on the columns and data returned.
type MeterReportData struct {
	// The type of meter the data is for.
	MeterType *string `json:"meterType,omitempty"`

	// The columns provided in the data.
	Columns *string `json:"columns,omitempty"`

	// A list of rows of CSV data matching the columns property.
	Data []string `json:"data,omitempty"`
}
