package objects

// The EquipmentSetting object represents the alert/reminder type which is
// associated with and dependent upon specific equipment controlled by the
// Thermostat. It is used when getting/setting the Thermostat
// NotificationSettings
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/NotificationSettings.shtml)
// object.
//
// Note: Only the notification settings for the equipment/devices currently
// controlled by the Thermostat are returned during GET request, and only those
// same settings can be updated using the POST request.
//
// The type corresponds to the Alert.notificationType returned when alerts are
// also included in the selection. See Alert
// (https://www.ecobee.com/home/developer/api/documentation/v1/objects/Alert.shtml)
// for more information.
type EquipmentSetting struct {
	// The date the filter was last changed for this equipment. String format:
	// YYYY-MM-DD
	FilterLastChanged *string `json:"filterLastChanged,omitempty"`

	// The value representing the life of the filter. This value is expressed in
	// month or hour, which is specified in the the filterLifeUnits property.
	FilterLife *int `json:"filterLife,omitempty"`

	// The units the filterLife field is measured in. Possible values are: month,
	// hour. month has a range of 1 - 12. hour has a range of 100 - 10000.
	FilterLifeUnits *string `json:"filterLifeUnits,omitempty"`

	// The date the reminder will be triggered. This is a read-only field and
	// cannot be modified through the API. The value is calculated and set by the
	// thermostat.
	RemindMeDate *string `json:"remindMeDate,omitempty"`

	// Boolean value representing whether or not alerts/reminders are enabled for
	// this notification type or not.
	Enabled *bool `json:"enabled,omitempty"`

	// The type of notification. Possible values are: hvac, furnaceFilter,
	// humidifierFilter, dehumidifierFilter, ventilator, ac, airFilter, airCleaner,
	// uvLamp
	Type *string `json:"type,omitempty"`

	// Boolean value representing whether or not alerts/reminders should be sent to
	// the technician/contractor associated with the thermostat.
	RemindTechnician *bool `json:"remindTechnician,omitempty"`
}
