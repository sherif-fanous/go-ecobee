package objects

// The runtime object represents the last known thermostat running state. This
// state is composed from the last interval status message received from a
// thermostat. It is also updated each time the thermostat posts configuration
// changes to the server.
//
// The runtime object contains the last 5 minute interval value sent by the
// thermostat for the past 15 minutes of runtime. The thermostat updates the
// server every 15 minutes with the last three 5 minute readings.
//
// The actual temperature and humidity will also be updated when the equipment
// state changes by the thermostat, this may occur at a frequency of 3 minutes,
// however it is only transmitted when there is an equipment state change on the
// thermostat.
//
// The runtime object contains two fields, desiredHeatRange and
// desiredCoolRange, which can be queried and used to determine that any holds
// being set through the API will not be adjusted. The API caller should check
// these ranges before calling the setHold function to mitigate against the new
// set points being adjusted by the server if the values are outside the
// acceptable ranges.
//
// See Thermostat Interval Report Data
// (https://www.ecobee.com/home/developer/api/introduction/core-concepts.shtml#data)
// for additional information about the interval readings.
//
// The runtime object is read-only.
type Runtime struct {
	// The current runtime revision. Equivalent in meaning to the runtime revision
	// number in the thermostat summary call.
	RuntimeRev *string `json:"runtimeRev,omitempty"`

	// Whether the thermostat is currently connected to the server.
	Connected *bool `json:"connected,omitempty"`

	// The UTC date/time stamp of when the thermostat first connected to the ecobee
	// server.
	FirstConnected *string `json:"firstConnected,omitempty"`

	// The last recorded connection date and time.
	ConnectDateTime *string `json:"connectDateTime,omitempty"`

	// The last recorded disconnection date and time.
	DisconnectDateTime *string `json:"disconnectDateTime,omitempty"`

	// The UTC date/time stamp of when the thermostat was updated. Format:
	// YYYY-MM-DD HH:MM:SS
	LastModified *string `json:"lastModified,omitempty"`

	// The UTC date/time stamp of when the thermostat last posted its runtime
	// information. Format: YYYY-MM-DD HH:MM:SS
	LastStatusModified *string `json:"lastStatusModified,omitempty"`

	// The UTC date of the last runtime reading. Format: YYYY-MM-DD
	RuntimeDate *string `json:"runtimeDate,omitempty"`

	// The last 5 minute interval which was updated by the thermostat telemetry
	// update. Subtract 2 from this interval to obtain the beginning interval for
	// the last 3 readings. Multiply by 5 mins to obtain the minutes of the day.
	// Range: 0-287
	RuntimeInterval *int `json:"runtimeInterval,omitempty"`

	// The current temperature displayed on the thermostat.
	ActualTemperature *int `json:"actualTemperature,omitempty"`

	// The current humidity % shown on the thermostat.
	ActualHumidity *int `json:"actualHumidity,omitempty"`

	// The dry-bulb temperature recorded by the thermostat. When
	// Energy.FeelsLikeMode is set to humidex, Runtime.actualTemperature will
	// report a "feels like" temperature.
	RawTemperature *int `json:"rawTemperature,omitempty"`

	// The currently displayed icon on the thermostat.
	ShowIconMode *int `json:"showIconMode,omitempty"`

	// The desired heat temperature as per the current running program or active event.
	DesiredHeat *int `json:"desiredHeat,omitempty"`

	// The desired cool temperature as per the current running program or active event.
	DesiredCool *int `json:"desiredCool,omitempty"`

	// The desired humidity set point.
	DesiredHumidity *int `json:"desiredHumidity,omitempty"`

	// The desired dehumidification set point.
	DesiredDehumidity *int `json:"desiredDehumidity,omitempty"`

	// The desired fan mode. Values: auto, on or null if the HVAC system is off and
	// the thermostat is not controlling a fan independently.
	DesiredFanMode *string `json:"desiredFanMode,omitempty"`

	// This field provides the possible valid range for which a desiredHeat
	// setpoint can be set to. This value takes into account the thermostat heat
	// temperature limits as well the running program or active events. Values are
	// returned as an Integer array representing the canonical minimum and maximim,
	// e.g. [450,790].
	DesiredHeatRange []int `json:"desiredHeatRange,omitempty"`

	// This field provides the possible valid range for which a desiredCool
	// setpoint can be set to. This value takes into account the thermostat cool
	// temperature limits as well the running program or active events. Values are
	// returned as an Integer array representing the canonical minimum and maximim,
	// e.g. [650,920].
	DesiredCoolRange []int `json:"desiredCoolRange,omitempty"`
}
