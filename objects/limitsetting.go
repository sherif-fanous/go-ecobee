package objects

// The LimitSetting object represents the alert/reminder type which is
// associated specific values, such as highHeat or lowHumidity. It is used when
// getting/setting the Thermostat NotificationSettings
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/NotificationSettings.shtml)
// object.
//
// The type corresponds to the Alert.notificationType returned when alerts are
// also included in the selection. See Alert
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/Alert.shtml)
// for more information.
//
// When POSTing updates to the LimitSettings there are certain requirements and
// behaviour the API user needs to be aware of:
type LimitSetting struct {
	// The value of the limit to set. For temperatures the value is expressed as
	// degrees Fahrenheit, multipled by 10. For humidity values are expressed as a
	// percentage from 5 to 95. See here for more information.
	Limit *int `json:"limit,omitempty"`

	// Boolean value representing whether or not alerts/reminders are enabled for
	// this notification type or not.
	Enabled *bool `json:"enabled,omitempty"`

	// The type of notification. Possible values are: lowTemp, highTemp,
	// lowHumidity, highHumidity, auxHeat, auxOutdoor
	Type *string `json:"type,omitempty"`

	// Boolean value representing whether or not alerts/reminders should be sent to
	// the technician/contractor associated with the thermostat.
	RemindTechnician *bool `json:"remindTechnician,omitempty"`
}
