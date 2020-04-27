package objects

// A SecuritySettings object defines the security settings which a thermostat
// may have. Currently this object stores data specific to access control.
//
// If any of the XXXAccess fields are not supplied they will default to false.
// So to set all to false where previously some were set to true the caller can
// either pass all the XXXAccess fields explicitly, or pass none and the default
// will be set for each.
//
// The userAccessCode may be modified. However, it is important to note that for
// security reason calls to GET this value will return null if no value is
// present, or a masked value of '****' if the value has been set.
//
// Note that this object is only available to applications which have implicit
// authorization such as utilities. See here for more information: Authorization
// Intro
// (https://www.ecobee.com/home/developer/api/documentation/v1/auth/auth-intro.shtml).
type SecuritySettings struct {
	// The 4-digit user access code for the thermostat. The code must be set when
	// enabling access control. See the callout above for more information.
	UserAccessCode *string `json:"userAccessCode,omitempty"`

	// The flag for determing whether there are any restrictions on the thermostat
	// regarding access control. Default value is false. If all other values are
	// true this value will default to true.
	AllUserAccess *bool `json:"allUserAccess,omitempty"`

	// The flag for determing whether there are any restrictions on the thermostat
	// regarding access control to the Thermostat.Program. Default value is false,
	// unless allUserAccess is true.
	ProgramAccess *bool `json:"programAccess,omitempty"`

	// The flag for determing whether there are any restrictions on the thermostat
	// regarding access control to the Thermostat system and settings. Default
	// value is false, unless allUserAccess is true.
	DetailsAccess *bool `json:"detailsAccess,omitempty"`

	// The flag for determing whether there are any restrictions on the thermostat
	// regarding access control to the Thermostat quick save functionality. Default
	// value is false, unless allUserAccess is true.
	QuickSaveAccess *bool `json:"quickSaveAccess,omitempty"`

	// The flag for determing whether there are any restrictions on the thermostat
	// regarding access control to the Thermostat vacation functionality. Default
	// value is false, unless allUserAccess is true.
	VacationAccess *bool `json:"vacationAccess,omitempty"`
}
