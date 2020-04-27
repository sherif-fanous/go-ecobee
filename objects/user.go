package objects

// The User object. The User object contains information pertaining to the User
// associated with a thermostat.
type User struct {
	// The User login userName. Usually a valid email address.
	UserName *string `json:"userName,omitempty"`

	// The User display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The User first name.
	FirstName *string `json:"firstName,omitempty"`

	// The User last name.
	LastName *string `json:"lastName,omitempty"`

	// The User title such as Mr. or Mrs.
	Honorific *string `json:"honorific,omitempty"`

	// The User date of registration.
	RegisterDate *string `json:"registerDate,omitempty"`

	// The User time of registration.
	RegisterTime *string `json:"registerTime,omitempty"`

	// The Thermostat identifier this User is associated with.
	DefaultThermostatIdentifier *string `json:"defaultThermostatIdentifier,omitempty"`

	// The User management reference.
	ManagementRef *string `json:"managementRef,omitempty"`

	// The User utility reference.
	UtilityRef *string `json:"utilityRef,omitempty"`

	// The User support reference.
	SupportRef *string `json:"supportRef,omitempty"`

	// The User phone Number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	//
	UtilityTimeZone *string `json:"utilityTimeZone,omitempty"`

	//
	ManagementTimeZone *string `json:"managementTimeZone,omitempty"`

	//
	IsResidential *bool `json:"isResidential,omitempty"`

	// Whether this user has a developer role.
	IsDeveloper *bool `json:"isDeveloper,omitempty"`

	// Whether this user has a management role.
	IsManagement *bool `json:"isManagement,omitempty"`

	// Whether this user has a utility role.
	IsUtility *bool `json:"isUtility,omitempty"`

	// Whether this user has a contractor role.
	IsContractor *bool `json:"isContractor,omitempty"`
}
