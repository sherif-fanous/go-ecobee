package objects

// The Settings contains all the configuration properties of the thermostat.
type Settings struct {
	// The current HVAC mode the thermostat is in. Values: auto, auxHeatOnly, cool,
	// heat, off.
	HVACMode *string `json:"hvacMode,omitempty"`

	// The last service date of the HVAC equipment.
	LastServiceDate *string `json:"lastServiceDate,omitempty"`

	// Whether to send an alert when service is required again.
	ServiceRemindMe *bool `json:"serviceRemindMe,omitempty"`

	// The user configured monthly interval between HVAC service reminders
	MonthsBetweenService *int `json:"monthsBetweenService,omitempty"`

	// Date to be reminded about the next HVAC service date.
	RemindMeDate *string `json:"remindMeDate,omitempty"`

	// The ventilator mode. Values: auto, minontime, on, off.
	Vent *string `json:"vent,omitempty"`

	// The minimum time in minutes the ventilator is configured to run. The
	// thermostat will always guarantee that the ventilator runs for this minimum
	// duration whenever engaged.
	VentilatorMinOnTime *int `json:"ventilatorMinOnTime,omitempty"`

	// Whether the technician associated with this thermostat should receive the
	// HVAC service reminders as well.
	ServiceRemindTechnician *bool `json:"serviceRemindTechnician,omitempty"`

	// A note about the physical location where the SMART or EMS Equipment
	// Interface module is located.
	EILocation *string `json:"eiLocation,omitempty"`

	// The temperature at which a cold temp alert is triggered.
	ColdTempAlert *int `json:"coldTempAlert,omitempty"`

	// Whether cold temperature alerts are enabled.
	ColdTempAlertEnabled *bool `json:"coldTempAlertEnabled,omitempty"`

	// The temperature at which a hot temp alert is triggered.
	HotTempAlert *int `json:"hotTempAlert,omitempty"`

	// Whether hot temperature alerts are enabled.
	HotTempAlertEnabled *bool `json:"hotTempAlertEnabled,omitempty"`

	// The number of cool stages the connected HVAC equipment supports.
	CoolStages *int `json:"coolStages,omitempty"`

	// The number of heat stages the connected HVAC equipment supports.
	HeatStages *int `json:"heatStages,omitempty"`

	// The maximum automated set point set back offset allowed in degrees.
	MaxSetBack *int `json:"maxSetBack,omitempty"`

	// The maximum automated set point set forward offset allowed in degrees.
	MaxSetForward *int `json:"maxSetForward,omitempty"`

	// The set point set back offset, in degrees, configured for a quick save event.
	QuickSaveSetBack *int `json:"quickSaveSetBack,omitempty"`

	// The set point set forward offset, in degrees, configured for a quick save event.
	QuickSaveSetForward *int `json:"quickSaveSetForward,omitempty"`

	// Whether the thermostat is controlling a heat pump.
	HasHeatPump *bool `json:"hasHeatPump,omitempty"`

	// Whether the thermostat is controlling a forced air furnace.
	HasForcedAir *bool `json:"hasForcedAir,omitempty"`

	// Whether the thermostat is controlling a boiler.
	HasBoiler *bool `json:"hasBoiler,omitempty"`

	// Whether the thermostat is controlling a humidifier.
	HasHumidifier *bool `json:"hasHumidifier,omitempty"`

	// Whether the thermostat is controlling an energy recovery ventilator.
	HasERV *bool `json:"hasErv,omitempty"`

	// Whether the thermostat is controlling a heat recovery ventilator.
	HasHRV *bool `json:"hasHrv,omitempty"`

	// Whether the thermostat is in frost control mode.
	CondensationAvoid *bool `json:"condensationAvoid,omitempty"`

	// Whether the thermostat is configured to report in degrees Celsius.
	UseCelsius *bool `json:"useCelsius,omitempty"`

	// Whether the thermostat is using 12hr time format.
	UseTimeFormat12 *bool `json:"useTimeFormat12,omitempty"`

	// Multilanguage support, currently only "en" - english is supported. In future
	// others locales can be supported.
	Locale *string `json:"locale,omitempty"`

	// The minimum humidity level (in percent) set point for the humidifier
	Humidity *string `json:"humidity,omitempty"`

	// The humidifier mode. Values: auto, manual, off.
	HumidifierMode *string `json:"humidifierMode,omitempty"`

	// The thermostat backlight intensity when on. A value between 0 and 10, with 0
	// meaning 'off' - the zero value may not be honored by all ecobee versions.
	BacklightOnIntensity *int `json:"backlightOnIntensity,omitempty"`

	// The thermostat backlight intensity when asleep. A value between 0 and 10,
	// with 0 meaning 'off' - the zero value may not be honored by all ecobee
	// versions.
	BacklightSleepIntensity *int `json:"backlightSleepIntensity,omitempty"`

	// The time in seconds before the thermostat screen goes into sleep mode.
	BacklightOffTime *int `json:"backlightOffTime,omitempty"`

	// The field is deprecated. Please use Audio.soundTickVolume.
	SoundTickVolume *int `json:"soundTickVolume,omitempty"`

	// The field is deprecated. Please use Audio.soundAlertVolume.
	SoundAlertVolume *int `json:"soundAlertVolume,omitempty"`

	// The minimum time the compressor must be off for in order to prevent
	// short-cycling.
	CompressorProtectionMinTime *int `json:"compressorProtectionMinTime,omitempty"`

	// The minimum outdoor temperature that the compressor can operate at - applies
	// more to air source heat pumps than geothermal.
	CompressorProtectionMinTemp *int `json:"compressorProtectionMinTemp,omitempty"`

	// The difference between current temperature and set-point that will trigger
	// stage 2 heating.
	Stage1HeatingDifferentialTemp *int `json:"stage1HeatingDifferentialTemp,omitempty"`

	// The difference between current temperature and set-point that will trigger
	// stage 2 cooling.
	Stage1CoolingDifferentialTemp *int `json:"stage1CoolingDifferentialTemp,omitempty"`

	// The time after a heating cycle that the fan will run for to extract any
	// heating left in the system - 30 second default.
	Stage1HeatingDissipationTime *int `json:"stage1HeatingDissipationTime,omitempty"`

	// The time after a cooling cycle that the fan will run for to extract any
	// cooling left in the system - 30 second default (for not)
	Stage1CoolingDissipationTime *int `json:"stage1CoolingDissipationTime,omitempty"`

	// The flag to tell if the heat pump is in heating mode or in cooling when the
	// relay is engaged. If set to zero it's heating when the reversing valve is
	// open, cooling when closed and if it's one - it's the opposite.
	HeatPumpReversalOnCool *bool `json:"heatPumpReversalOnCool,omitempty"`

	// Whether fan control by the Thermostat is required in auxiliary heating
	// (gas/electric/boiler), otherwise controlled by furnace.
	FanControlRequired *bool `json:"fanControlRequired,omitempty"`

	// The minimum time, in minutes, to run the fan each hour. Value from 1 to 60.
	FanMinOnTime *int `json:"fanMinOnTime,omitempty"`

	// The minimum temperature difference between the heat and cool values. Used to
	// ensure that when thermostat is in auto mode, the heat and cool values are
	// separated by at least this value.
	HeatCoolMinDelta *int `json:"heatCoolMinDelta,omitempty"`

	// The amount to adjust the temperature reading in degrees F - this value is
	// subtracted from the temperature read from the sensor.
	TempCorrection *int `json:"tempCorrection,omitempty"`

	// The default end time setting the thermostat applies to user temperature
	// holds. Values useEndTime4hour, useEndTime2hour (EMS Only), nextPeriod,
	// indefinite, askMe
	HoldAction *string `json:"holdAction,omitempty"`

	// Whether the Thermostat uses a geothermal / ground source heat pump.
	HeatPumpGroundWater *bool `json:"heatPumpGroundWater,omitempty"`

	// Whether the thermostat is connected to an electric HVAC system.
	HasElectric *bool `json:"hasElectric,omitempty"`

	// Whether the thermostat is connected to a dehumidifier. If true or
	// dehumidifyOvercoolOffset > 0 then allow setting dehumidifierMode and
	// dehumidifierLevel.
	HasDehumidifier *bool `json:"hasDehumidifier,omitempty"`

	// The dehumidifier mode. Values: on, off. If set to off then the dehumidifier
	// will not run, nor will the AC overcool run.
	DehumidifierMode *string `json:"dehumidifierMode,omitempty"`

	// The dehumidification set point in percentage.
	DehumidifierLevel *int `json:"dehumidifierLevel,omitempty"`

	// Whether the thermostat should use AC overcool to dehumidify. When set to
	// true a positive integer value must be supplied for dehumidifyOvercoolOffset
	// otherwise an API validation exception will be thrown.
	DehumidifyWithAC *bool `json:"dehumidifyWithAC,omitempty"`

	// Whether the thermostat should use AC overcool to dehumidify and what that
	// temperature offset should be. A value of 0 means this feature is disabled
	// and dehumidifyWithAC will be set to false. Value represents the value in F
	// to subract from the current set point. Values should be in the range 0 - 50
	// and be divisible by 5.
	DehumidifyOvercoolOffset *int `json:"dehumidifyOvercoolOffset,omitempty"`

	// If enabled, allows the Thermostat to be put in HVACAuto mode.
	AutoHeatCoolFeatureEnabled *bool `json:"autoHeatCoolFeatureEnabled,omitempty"`

	// Whether the alert for when wifi is offline is enabled.
	WiFiOfflineAlert *bool `json:"wifiOfflineAlert,omitempty"`

	// The minimum heat set point allowed by the thermostat firmware.
	HeatMinTemp *int `json:"heatMinTemp,omitempty"`

	// The maximum heat set point allowed by the thermostat firmware.
	HeatMaxTemp *int `json:"heatMaxTemp,omitempty"`

	// The minimum cool set point allowed by the thermostat firmware.
	CoolMinTemp *int `json:"coolMinTemp,omitempty"`

	// The maximum cool set point allowed by the thermostat firmware.
	CoolMaxTemp *int `json:"coolMaxTemp,omitempty"`

	// The maximum heat set point configured by the user's preferences.
	HeatRangeHigh *int `json:"heatRangeHigh,omitempty"`

	// The minimum heat set point configured by the user's preferences.
	HeatRangeLow *int `json:"heatRangeLow,omitempty"`

	// The maximum cool set point configured by the user's preferences.
	CoolRangeHigh *int `json:"coolRangeHigh,omitempty"`

	// The minimum heat set point configured by the user's preferences.
	CoolRangeLow *int `json:"coolRangeLow,omitempty"`

	// The user access code value for this thermostat. See the SecuritySettings
	// object for more information.
	UserAccessCode *string `json:"userAccessCode,omitempty"`

	// The integer representation of the user access settings. See the
	// SecuritySettings object for more information.
	UserAccessSetting *int `json:"userAccessSetting,omitempty"`

	// The temperature at which an auxHeat temperature alert is triggered.
	AuxRuntimeAlert *int `json:"auxRuntimeAlert,omitempty"`

	// The temperature at which an auxOutdoor temperature alert is triggered.
	AuxOutdoorTempAlert *int `json:"auxOutdoorTempAlert,omitempty"`

	// The maximum outdoor temperature above which aux heat will not run.
	AuxMaxOutdoorTemp *int `json:"auxMaxOutdoorTemp,omitempty"`

	// Whether the auxHeat temperature alerts are enabled.
	AuxRuntimeAlertNotify *bool `json:"auxRuntimeAlertNotify,omitempty"`

	// Whether the auxOutdoor temperature alerts are enabled.
	AuxOutdoorTempAlertNotify *bool `json:"auxOutdoorTempAlertNotify,omitempty"`

	// Whether the auxHeat temperature alerts for the technician are enabled.
	AuxRuntimeAlertNotifyTechnician *bool `json:"auxRuntimeAlertNotifyTechnician,omitempty"`

	// Whether the auxOutdoor temperature alerts for the technician are enabled.
	AuxOutdoorTempAlertNotifyTechnician *bool `json:"auxOutdoorTempAlertNotifyTechnician,omitempty"`

	// Whether the thermostat should use pre heating to reach the set point on time.
	DisablePreHeating *bool `json:"disablePreHeating,omitempty"`

	// Whether the thermostat should use pre cooling to reach the set point on time.
	DisablePreCooling *bool `json:"disablePreCooling,omitempty"`

	// Whether an installer code is required.
	InstallerCodeRequired *bool `json:"installerCodeRequired,omitempty"`

	// Whether Demand Response requests are accepted by this thermostat. Possible
	// values are: always, askMe, customerSelect, defaultAccept, defaultDecline,
	// never.
	DRAccept *string `json:"drAccept,omitempty"`

	// Whether the property is a rental, or not.
	IsRentalProperty *bool `json:"isRentalProperty,omitempty"`

	// Whether to use a zone controller or not.
	UseZoneController *bool `json:"useZoneController,omitempty"`

	// Whether random start delay is enabled for cooling.
	RandomStartDelayCool *int `json:"randomStartDelayCool,omitempty"`

	// Whether random start delay is enabled for heating.
	RandomStartDelayHeat *int `json:"randomStartDelayHeat,omitempty"`

	// The humidity level to trigger a high humidity alert.
	HumidityHighAlert *int `json:"humidityHighAlert,omitempty"`

	// The humidity level to trigger a low humidity alert.
	HumidityLowAlert *int `json:"humidityLowAlert,omitempty"`

	// Whether heat pump alerts are disabled.
	DisableHeatPumpAlerts *bool `json:"disableHeatPumpAlerts,omitempty"`

	// Whether alerts are disabled from showing on the thermostat.
	DisableAlertsOnIDT *bool `json:"disableAlertsOnIdt,omitempty"`

	// Whether humidification alerts are enabled to the thermsotat owner.
	HumidityAlertNotify *bool `json:"humidityAlertNotify,omitempty"`

	// Whether humidification alerts are enabled to the technician associated with
	// the thermsotat.
	HumidityAlertNotifyTechnician *bool `json:"humidityAlertNotifyTechnician,omitempty"`

	// Whether temperature alerts are enabled to the thermsotat owner.
	TempAlertNotify *bool `json:"tempAlertNotify,omitempty"`

	// Whether temperature alerts are enabled to the technician associated with the
	// thermostat.
	TempAlertNotifyTechnician *bool `json:"tempAlertNotifyTechnician,omitempty"`

	// The dollar amount the owner specifies for their desired maximum electricity bill.
	MonthlyElectricityBillLimit *int `json:"monthlyElectricityBillLimit,omitempty"`

	// Whether electricity bill alerts are enabled.
	EnableElectricityBillAlert *bool `json:"enableElectricityBillAlert,omitempty"`

	// Whether electricity bill projection alerts are enabled
	EnableProjectedElectricityBillAlert *bool `json:"enableProjectedElectricityBillAlert,omitempty"`

	// The day of the month the owner's electricity usage is billed.
	ElectricityBillingDayOfMonth *int `json:"electricityBillingDayOfMonth,omitempty"`

	// The owners billing cycle duration in months.
	ElectricityBillCycleMonths *int `json:"electricityBillCycleMonths,omitempty"`

	// The annual start month of the owners billing cycle.
	ElectricityBillStartMonth *int `json:"electricityBillStartMonth,omitempty"`

	// The number of minutes to run ventilator per hour when home.
	VentilatorMinOnTimeHome *int `json:"ventilatorMinOnTimeHome,omitempty"`

	// The number of minutes to run ventilator per hour when away.
	VentilatorMinOnTimeAway *int `json:"ventilatorMinOnTimeAway,omitempty"`

	// Determines whether or not to turn the backlight off during sleep.
	BacklightOffDuringSleep *bool `json:"backlightOffDuringSleep,omitempty"`

	// When set to true if no occupancy motion detected thermostat will go into
	// indefinite away hold, until either the user presses resume schedule or
	// motion is detected.
	AutoAway *bool `json:"autoAway,omitempty"`

	// When set to true if a larger than normal delta is found between sensors the
	// fan will be engaged for 15min/hour.
	SmartCirculation *bool `json:"smartCirculation,omitempty"`

	// When set to true if a sensor has detected presence for more than 10 minutes
	// then include that sensor in temp average. If no activity has been seen on a
	// sensor for more than 1 hour then remove this sensor from temperature
	// average.
	FollowMeComfort *bool `json:"followMeComfort,omitempty"`

	// This read-only field represents the type of ventilator present for the
	// Thermostat. The possible values are none, ventilator, hrv, and erv.
	VentilatorType *string `json:"ventilatorType,omitempty"`

	// This Boolean field represents whether the ventilator timer is on or off. The
	// default value is false. If set to true the ventilatorOffDateTime is set to
	// now() + 20 minutes. If set to false the ventilatorOffDateTime is set to it's
	// default value.
	IsVentilatorTimerOn *bool `json:"isVentilatorTimerOn,omitempty"`

	// This read-only field represents the Date and Time the ventilator will run
	// until. The default value is 2014-01-01 00:00:00.
	VentilatorOffDateTime *string `json:"ventilatorOffDateTime,omitempty"`

	// This Boolean field represents whether the HVAC system has a UV filter. The
	// default value is true.
	HasUVFilter *bool `json:"hasUVFilter,omitempty"`

	// This field represents whether to permit the cooling to operate when the
	// Outdoor temeperature is under a specific threshold, currently 55F. The
	// default value is false.
	CoolingLockout *bool `json:"coolingLockout,omitempty"`

	// Whether to use the ventilator to dehumidify when climate or calendar event
	// indicates the owner is home. The default value is false.
	VentilatorFreeCooling *bool `json:"ventilatorFreeCooling,omitempty"`

	// This field represents whether to permit dehumidifer to operate when the
	// heating is running. The default value is false.
	DehumidifyWhenHeating *bool `json:"dehumidifyWhenHeating,omitempty"`

	// This field represents whether or not to allow dehumification when cooling.
	// The default value is true.
	VentilatorDehumidify *bool `json:"ventilatorDehumidify,omitempty"`

	// The unique reference to the group this thermostat belongs to, if any. See
	// GET Group request and POST Group request for more information.
	GroupRef *string `json:"groupRef,omitempty"`

	// The name of the the group this thermostat belongs to, if any. See GET Group
	// request and POST Group request for more information.
	GroupName *string `json:"groupName,omitempty"`

	// The setting value for the group this thermostat belongs to, if any. See GET
	// Group request and POST Group request for more information.
	GroupSetting *int `json:"groupSetting,omitempty"`
}
