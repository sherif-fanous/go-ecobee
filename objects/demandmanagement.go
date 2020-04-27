package objects

// The Demand Management object defines a series of 12 5-minute intervals in a
// given hour of the day. Each value in the series is an adjustment to the
// thermostat desired temperature which is applied for the matching 5 minute
// interval in the hour. The temperature adjustment may be made in the range of
// -2F to +2F. The temperature value must be specified as per standard
// Temperature notation already described. Specifying 0 for the adjustment in
// the series will result in no adjustment being made. Values outside the +/- 2F
// range will result in error.
type DemandManagement struct {
	// The date (UTC) for the beginning of this day's demand management series.
	Date *string `json:"date,omitempty"`

	// The hour in the day this series begins at. Range: 0-23
	Hour *int `json:"hour,omitempty"`

	// The series of 12, 5 minute interval temperature adjustments in the hour.
	// Valid integer values are required for each interval. See Temperature values.
	TempOffsets []int `json:"tempOffsets,omitempty"`
}
