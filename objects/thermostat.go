package objects

// The thermostat object is the central piece of the ecobee API. All objects
// relate in one way or another to a real thermostat. The thermostat object and
// its component objects define the real thermostat device.
type Thermostat struct {
	// The unique thermostat serial number.
	Identifier *string `json:"identifier,omitempty"`

	// A user defined name for a thermostat.
	Name *string `json:"name,omitempty"`

	// The current thermostat configuration revision.
	ThermostatRev *string `json:"thermostatRev,omitempty"`

	// Whether the user registered the thermostat.
	IsRegistered *bool `json:"isRegistered,omitempty"`

	// The thermostat model number. Values: apolloSmart, apolloEms, idtSmart,
	// idtEms, siSmart, siEms, athenaSmart, athenaEms, corSmart, nikeSmart, nikeEms
	ModelNumber *string `json:"modelNumber,omitempty"`

	// The thermostat brand.
	Brand *string `json:"brand,omitempty"`

	// The comma-separated list of the thermostat's additional features, if any.
	Features *string `json:"features,omitempty"`

	// The last modified date time for the thermostat configuration.
	LastModified *string `json:"lastModified,omitempty"`

	// The current time in the thermostat's time zone
	ThermostatTime *string `json:"thermostatTime,omitempty"`

	// The current time in UTC.
	UTCTime *string `json:"utcTime,omitempty"`

	// The thermostat audio configuration
	Audio *Audio `json:"audio,omitempty"`

	// The list of Alert objects tied to the thermostat
	Alerts []Alert `json:"alerts,omitempty"`

	//
	Reminders []ThermostatReminder2 `json:"reminders,omitempty"`

	// The thermostat Setting object linked to the thermostat
	Settings *Settings `json:"settings,omitempty"`

	// The Runtime state object for the thermostat
	Runtime *Runtime `json:"runtime,omitempty"`

	// The ExtendedRuntime object for the thermostat
	ExtendedRuntime *ExtendedRuntime `json:"extendedRuntime,omitempty"`

	// The Electricity object for the thermostat
	Electricity *Electricity `json:"electricity,omitempty"`

	// The list of Device objects linked to the thermostat
	Devices []Device `json:"devices,omitempty"`

	// The Location object for the thermostat
	Location *Location `json:"location,omitempty"`

	// The thermostat energy configuration
	Energy *Energy `json:"energy,omitempty"`

	// The Technician object associated with the thermostat containing the
	// technician contact information
	Technician *Technician `json:"technician,omitempty"`

	// The Utility object associated with the thermostat containing the utility
	// company information
	Utility *Utility `json:"utility,omitempty"`

	// The Management object associated with the thermostat containing the
	// management company information
	Management *Management `json:"management,omitempty"`

	// The Weather object linked to the thermostat representing the current weather
	// on the thermostat.
	Weather *Weather `json:"weather,omitempty"`

	// The list of Event objects linked to the thermostat representing any events
	// that are active or scheduled.
	Events []Event `json:"events,omitempty"`

	// The Program object for the thermostat
	Program *Program `json:"program,omitempty"`

	// The houseDetails object contains contains the information about the house
	// the thermostat is installed in.
	HouseDetails *HouseDetails `json:"houseDetails,omitempty"`

	// The OemCfg object contains information about the OEM specific thermostat.
	OEMCfg *ThermostatOEMCfg `json:"oemCfg,omitempty"`

	// The status of all equipment controlled by this Thermostat. Only running
	// equipment is listed in the CSV String. Values: heatPump, heatPump2,
	// heatPump3, compCool1, compCool2, auxHeat1, auxHeat2, auxHeat3, fan,
	// humidifier, dehumidifier, ventilator, economizer, compHotWater, auxHotWater.
	// Note: If no equipment is currently running an empty String is returned. If
	// Settings.hasHeatPump is true, heatPump value will be returned for heating,
	// compCool for cooling, and auxHeat for aux heat. If Settings.hasForcedAir or
	// Settings.hasBoiler is true, auxHeat value will be returned for heating and
	// compCool for cooling (heatPump will not show up for heating).
	EquipmentStatus *string `json:"equipmentStatus,omitempty"`

	// The NotificationSettings object containing the configuration for Alert and
	// Reminders for the Thermostat.
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`

	// The Privacy object containing the privacy settings for the Thermostat. Note:
	// access to this object is restricted to callers with implicit authentication.
	Privacy *ThermostatPrivacy `json:"privacy,omitempty"`

	// The Version object containing the firmware version information for the
	// Thermostat. For example: "3.5.0.3957".
	Version *Version `json:"version,omitempty"`

	// The SecuritySettings object containing the security settings for the Thermostat.
	SecuritySettings *SecuritySettings `json:"securitySettings,omitempty"`

	//
	FilterSubscription *APIFilterSubscription `json:"filterSubscription,omitempty"`

	// The list of RemoteSensor objects for the Thermostat.
	RemoteSensors []RemoteSensor `json:"remoteSensors,omitempty"`
}
