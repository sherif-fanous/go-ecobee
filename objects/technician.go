package objects

// The technician object contains information pertaining to the technician
// associated with a thermostat. The technician may not be modified through the
// API.
type Technician struct {
	// The internal ecobee unique identifier for this contractor.
	ContractorRef *string `json:"contractorRef,omitempty"`

	// The company name of the technician.
	Name *string `json:"name,omitempty"`

	// The technician's contact phone number.
	Phone *string `json:"phone,omitempty"`

	// The technician's street address.
	StreetAddress *string `json:"streetAddress,omitempty"`

	// The technician's city.
	City *string `json:"city,omitempty"`

	// The technician's State or Province.
	ProvinceState *string `json:"provinceState,omitempty"`

	// The technician's country.
	Country *string `json:"country,omitempty"`

	// The technician's ZIP or Postal Code.
	PostalCode *string `json:"postalCode,omitempty"`

	// The technician's email address.
	Email *string `json:"email,omitempty"`

	// The technician's web site.
	Web *string `json:"web,omitempty"`
}
