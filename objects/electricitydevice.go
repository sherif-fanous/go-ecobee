package objects

// An electricity device represents an energy recording device. At this time,
// only meters are supported by the API.
type ElectricityDevice struct {
	// The name of the device
	Name *string `json:"name,omitempty"`

	// The list of Electricity Tiers containing the break down of daily electricity
	// consumption of the device for the day, broken down per pricing tier.
	Tiers []ElectricityTier `json:"tiers,omitempty"`

	// The last date/time the reading was updated in UTC time.
	LastUpdate *string `json:"lastUpdate,omitempty"`

	// The last three daily electricity cost reads from the device in cents with a
	// three decimal place precision.
	Cost []string `json:"cost,omitempty"`

	// The last three daily electricity consumption reads from the device in KWh
	// with a three decimal place precision.
	Consumption []string `json:"consumption,omitempty"`
}
