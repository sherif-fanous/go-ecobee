package objects

// An electricity tier object represents the last reading from a given pricing
// tier if the utility provides such information. If there are no pricing tiers
// defined, than an unnamed tier will represent the total reading. The values
// represented here are a daily cumulative total in kWh. The cost is likewise a
// cumulative total in cents.
type ElectricityTier struct {
	// The tier name as defined by the Utility. May be an empty string if the tier
	// is undefined or the usage falls outside the defined tiers.
	Name *string `json:"name,omitempty"`

	// The last daily consumption reading collected. The reading format and
	// precision is to three decimal places in kWh.
	Consumption *string `json:"consumption,omitempty"`

	// The daily cumulative tier cost in dollars if defined by the Utility. May be
	// an empty string if undefined.
	Cost *string `json:"cost,omitempty"`
}
