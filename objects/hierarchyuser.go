package objects

// The information about a user in the EMS or Utility company.
type HierarchyUser struct {
	// The user name and login ID. It must be a valid email address.
	UserName *string `json:"userName,omitempty"`

	// The first name.
	FirstName *string `json:"firstName,omitempty"`

	// The last name.
	LastName *string `json:"lastName,omitempty"`

	// The phone number.
	Phone *string `json:"phone,omitempty"`

	// The timestamp of the last user login into the web portal. Format: YYYY-MM-DD
	// HH:MM:SS
	LastLogin *string `json:"lastLogin,omitempty"`

	// Whether the user is active and permitted to access to the system.
	Active *bool `json:"active,omitempty"`

	// Whether the user will receive alerts in email.
	EmailAlerts *bool `json:"emailAlerts,omitempty"`
}
