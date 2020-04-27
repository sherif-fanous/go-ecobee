package objects

// The Utility information the Thermostat
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/Thermostat.shtml)
// belongs to. The utility may not be modified through the API.
type Utility struct {
	// The Utility company name.
	Name *string `json:"name,omitempty"`

	// The Utility company contact phone number.
	Phone *string `json:"phone,omitempty"`

	// The Utility company email address.
	Email *string `json:"email,omitempty"`

	// The Utility company web site.
	Web *string `json:"web,omitempty"`
}
