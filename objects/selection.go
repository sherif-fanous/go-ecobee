package objects

// The selection object defines the resources and information to return as part
// of a response. The selection is required in all requests however meaning of
// some selection fields is only meaningful to certain types of requests.
//
// The selectionType parameter defines the type of selection to perform. The
// selectionMatch specifies the matching criteria for the type specified.
type Selection struct {
	// The type of match data supplied: Values: thermostats, registered, managementSet.
	SelectionType *string `json:"selectionType,omitempty"`

	// The match data based on selectionType (e.g. a list of thermostat idendifiers
	// in the case of a selectionType of thermostats)
	SelectionMatch *string `json:"selectionMatch,omitempty"`

	// Include the thermostat runtime object. If not specified, defaults to false.
	IncludeRuntime *bool `json:"includeRuntime,omitempty"`

	// Include the extended thermostat runtime object. If not specified, defaults
	// to false.
	IncludeExtendedRuntime *bool `json:"includeExtendedRuntime,omitempty"`

	// Include the electricity readings object. If not specified, defaults to false.
	IncludeElectricity *bool `json:"includeElectricity,omitempty"`

	// Include the thermostat settings object. If not specified, defaults to false.
	IncludeSettings *bool `json:"includeSettings,omitempty"`

	// Include the thermostat location object. If not specified, defaults to false.
	IncludeLocation *bool `json:"includeLocation,omitempty"`

	// Include the thermostat program object. If not specified, defaults to false.
	IncludeProgram *bool `json:"includeProgram,omitempty"`

	// Include the thermostat calendar events objects. If not specified, defaults
	// to false.
	IncludeEvents *bool `json:"includeEvents,omitempty"`

	// Include the thermostat device configuration objects. If not specified,
	// defaults to false.
	IncludeDevice *bool `json:"includeDevice,omitempty"`

	// Include the thermostat technician object. If not specified, defaults to false.
	IncludeTechnician *bool `json:"includeTechnician,omitempty"`

	// Include the thermostat utility company object. If not specified, defaults to
	// false.
	IncludeUtility *bool `json:"includeUtility,omitempty"`

	// Include the thermostat management company object. If not specified, defaults
	// to false.
	IncludeManagement *bool `json:"includeManagement,omitempty"`

	// Include the thermostat's unacknowledged alert objects. If not specified,
	// defaults to false.
	IncludeAlerts *bool `json:"includeAlerts,omitempty"`

	//
	IncludeReminders *bool `json:"includeReminders,omitempty"`

	// Include the current thermostat weather forecast object. If not specified,
	// defaults to false.
	IncludeWeather *bool `json:"includeWeather,omitempty"`

	// Include the current thermostat house details object. If not specified,
	// defaults to false.
	IncludeHouseDetails *bool `json:"includeHouseDetails,omitempty"`

	// Include the current thermostat OemCfg object. If not specified, defaults to
	// false.
	IncludeOEMCfg *bool `json:"includeOemCfg,omitempty"`

	// Include the current thermostat equipment status information. If not
	// specified, defaults to false.
	IncludeEquipmentStatus *bool `json:"includeEquipmentStatus,omitempty"`

	// Include the current thermostat alert and reminders settings. If not
	// specified, defaults to false.
	IncludeNotificationSettings *bool `json:"includeNotificationSettings,omitempty"`

	// Include the current thermostat privacy settings. Note: access to this object
	// is restricted to callers with implict authentication, setting this value to
	// true without proper credentials will result in an authentication exception.
	IncludePrivacy *bool `json:"includePrivacy,omitempty"`

	// Include the current firmware version the Thermostat is running. If not
	// specified, defaults to false.
	IncludeVersion *bool `json:"includeVersion,omitempty"`

	// Include the current securitySettings object for the selected Thermostat(s).
	// If not specified, defaults to false.
	IncludeSecuritySettings *bool `json:"includeSecuritySettings,omitempty"`

	// Include the list of current thermostatRemoteSensor objects for the selected
	// Thermostat(s). If not specified, defaults to false.
	IncludeSensors *bool `json:"includeSensors,omitempty"`

	// Include the audio configuration for the selected Thermostat(s). If not
	// specified, defaults to false.
	IncludeAudio *bool `json:"includeAudio,omitempty"`

	// Include the energy configuration for the selected Thermostat(s). If not
	// specified, defaults to false.
	IncludeEnergy *bool `json:"includeEnergy,omitempty"`
}
