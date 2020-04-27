package objects

// The GeneralSetting object represent the General alert/reminder type. It is
// used when getting/setting the Thermostat NotificationSettings
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/NotificationSettings.shtml)
// object.
//
// The type corresponds to the Alert.notificationType returned when alerts are
// included in the selection. See Alert
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/Alert.shtml)
// for more information.
type GeneralSetting struct {
	// Boolean value representing whether or not alerts/reminders are enabled for
	// this notification type or not.
	Enabled *bool `json:"enabled,omitempty"`

	// The type of notification. Possible values are: temp
	Type *string `json:"type,omitempty"`

	// Boolean value representing whether or not alerts/reminders should be sent to
	// the technician/contractor associated with the thermostat.
	RemindTechnician *bool `json:"remindTechnician,omitempty"`
}
