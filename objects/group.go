package objects

// A Group object defines the group (and the related Group settings) which a
// thermostat may belong to. There could be a number of Groups and also a number
// of thermostats in each Group. The Group object allows the caller to define
// the Group name, which user preferences are shared across all thermostats in
// that Group, and indeed which Thermostats are part of that Group.
//
// The result is that if you modify the Group settings, for example set the
// synchronizeAlerts flag to true, any Alert changes made to any thermostat in
// that group will be shared with the remaining thermostats in the same group.
//
// The Grouping algorithm uses a "first group wins" strategy when a Thermostat
// is referenced in multiple groups. What this means in practice is that when
// the API request is processed and a Thermostat is referenced in more than one
// group, that Thermostat will only be added to the first Group (at head of
// array) and not to the others.
//
// If any of the synchronizeXXX fields are not supplied they will default to
// false. So to set all to false where previously some were set to true the
// caller can either pass all the synchronizeXXX fields explicitly, or pass none
// and the default will be set for each.
//
// The Group object may be modified. However, it is important to note that if
// the groupRef is not sent by the caller it is assumed that this is a new
// group, even if the groupName has not changed, and a new groupRef will be
// generated and returned. Therefore when updating a Group the groupRef must
// always be sent.
//
// Also note that if the thermostats list is not sent, or an empty list is sent,
// the Group will effectively be deleted as it will no longer contain any
// thermostats and any group information will be lost.
type Group struct {
	// The unique reference Id for the Group. If not supplied in the POST call, and
	// new groupRef will be generated.
	GroupRef *string `json:"groupRef,omitempty"`

	// The name for the Group.
	GroupName *string `json:"groupName,omitempty"`

	// Flag for whether to synchronize Thermostat Alerts with all other Thermostats
	// in the Group. Default is false.
	SynchronizeAlerts *bool `json:"synchronizeAlerts,omitempty"`

	// Flag for whether to synchronize the Thermostat mode with all other
	// Thermostats in the Group. Default is false.
	SynchronizeSystemMode *bool `json:"synchronizeSystemMode,omitempty"`

	// Flag for whether to synchronize the Thermostat schedule/Program details with
	// all other Thermostats in the Group. Default is false.
	SynchronizeSchedule *bool `json:"synchronizeSchedule,omitempty"`

	// Flag for whether to synchronize the Thermostat quick save settings with all
	// other Thermostats in the Group. Default is false.
	SynchronizeQuickSave *bool `json:"synchronizeQuickSave,omitempty"`

	// Flag for whether to synchronize the Thermostat reminders with all other
	// Thermostats in the Group. Default is false.
	SynchronizeReminders *bool `json:"synchronizeReminders,omitempty"`

	// Flag for whether to synchronize the Thermostat Technician/Contractor
	// Information with all other Thermostats in the Group. Default is false.
	SynchronizeContractorInfo *bool `json:"synchronizeContractorInfo,omitempty"`

	// Flag for whether to synchronize the Thermostat user preferences with all
	// other Thermostats in the Group. Default is false.
	SynchronizeUserPreferences *bool `json:"synchronizeUserPreferences,omitempty"`

	// Flag for whether to synchronize the Thermostat utility information with all
	// other Thermostats in the Group. Default is false.
	SynchronizeUtilityInfo *bool `json:"synchronizeUtilityInfo,omitempty"`

	// Flag for whether to synchronize the Thermostat Location with all other
	// Thermostats in the Group. Default is false.
	SynchronizeLocation *bool `json:"synchronizeLocation,omitempty"`

	// Flag for whether to synchronize the Thermostat reset with all other
	// Thermostats in the Group. Default is false.
	SynchronizeReset *bool `json:"synchronizeReset,omitempty"`

	// Flag for whether to synchronize the Thermostat vacation Program with all
	// other Thermostats in the Group. Default is false.
	SynchronizeVacation *bool `json:"synchronizeVacation,omitempty"`

	// The list of Thermostat identifiers which belong to the group. If an empty
	// list is sent the Group will be deleted.
	Thermostats []string `json:"thermostats,omitempty"`
}
