package objects

// The hierarchy privileges assigned to a user for a management set.
//
// If the user has any privilege, other than none on a set, this is equivalent
// to the view privilege.The privilege all grants all privileges regardless of
// the other privileges defined.
//
// The none privilege supersedes all other privileges and denies access to a
// set. The following privileges are only considered for the Utility portal:
// all, none, view. The others have no effect.
type HierarchyPrivilege struct {
	// The path of the management set.
	SetPath *string `json:"setPath,omitempty"`

	// The name of the management set.
	SetName *string `json:"setName,omitempty"`

	// The user name of the user associated with this privilege.
	UserName *string `json:"userName,omitempty"`

	// The user is permitted all privileges on the set.
	AllowAll *bool `json:"allowAll,omitempty"`

	// The user is denied any privilege and may not view the set.
	AllowNone *bool `json:"allowNone,omitempty"`

	// The user is only permitted to view the set and its contents.
	AllowView *bool `json:"allowView,omitempty"`

	// The user is permitted to make program changes.
	AllowProgram *bool `json:"allowProgram,omitempty"`

	// The user is permitted to create and edit vacation events.
	AllowVacation *bool `json:"allowVacation,omitempty"`

	// The user is permitted to edit thermostat settings.
	AllowSettings *bool `json:"allowSettings,omitempty"`

	// The user is permitted to access thermostat details such as desired
	// temperature, HVAC mode and humidity settings.
	AllowDetails *bool `json:"allowDetails,omitempty"`

	// The user is permitted to view thermostat reports.
	AllowReport *bool `json:"allowReport,omitempty"`

	// The user is permitted to manage user security.
	AllowSecurity *bool `json:"allowSecurity,omitempty"`

	// The user is permitted to manage management sets.
	AllowHierarchy *bool `json:"allowHierarchy,omitempty"`

	// The user is permitted to manage alerts.
	AllowAlerts *bool `json:"allowAlerts,omitempty"`

	// The user is permitted to manage account information and register/unregister
	// new users.
	AllowManageAccount *bool `json:"allowManageAccount,omitempty"`
}
