package objects

// The Management object contains information about the management company the
// thermostat belongs to. The Management object is read-only, it may be modified
// in the web portal.
type Management struct {
	// The administrative contact name.
	AdministrativeContact *string `json:"administrativeContact,omitempty"`

	// The billing contact name.
	BillingContact *string `json:"billingContact,omitempty"`

	// The company name.
	Name *string `json:"name,omitempty"`

	// The phone number.
	Phone *string `json:"phone,omitempty"`

	// The contact email address.
	Email *string `json:"email,omitempty"`

	// The company web site.
	Web *string `json:"web,omitempty"`

	// Whether to show management alerts on the thermostat.
	ShowAlertIDT *bool `json:"showAlertIdt,omitempty"`

	// Whether to show management alerts in the web portal.
	ShowAlertWeb *bool `json:"showAlertWeb,omitempty"`
}
