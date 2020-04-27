package objects

// The electricity object contains the last collected electricity usage
// measurements for the thermostat. An electricity object is composed of
// Electricity Devices, each of which contains readings from an Electricity
// Tier.
type Electricity struct {
	// The list of ElectricityDevice objects associated with the thermostat, each
	// representing a device such as an electric meter or remote load control.
	Devices []ElectricityDevice `json:"devices,omitempty"`
}
