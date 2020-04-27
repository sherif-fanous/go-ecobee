package objects

// The extended runtime object contains the last three 5 minute interval values
// sent by the thermostat for the past 15 minutes of runtime. The interval
// values are valuable when you are interested in analyzing the runtime data in
// a more granular fashion, at 5 minute increments rather than the more general
// 15 minute value from the Runtime Object.
//
// For the runtime values (i.e. heatPump, auxHeat, cool, etc.) refer to the
// Thermostat.Settings values (hasHeatPump, heatStages, coolStages) to determine
// whether a heat pump exists and how many stages the thermostat supports.
//
// The actual temperature and humidity will also be updated when the equipment
// state changes by the thermostat, this may occur at a frequency of 3 minutes,
// however it is only transmitted when there is an equipment state change on the
// thermostat.
//
// See Thermostat Interval Report Data
// (https://www.ecobee.com/home/developer/api/introduction/core-concepts.shtml#data)
// for additional information about the interval readings.
//
// The extended runtime object is read-only.
type ExtendedRuntime struct {
	// The UTC timestamp of the last value read. This timestamp is updated at a 15
	// min interval by the thermostat. For the 1st value, it is timestamp - 10
	// mins, for the 2nd value it is timestamp - 5 mins. Consider day boundaries
	// being straddled when using these values.
	LastReadingTimestamp *string `json:"lastReadingTimestamp,omitempty"`

	// The UTC date of the last runtime reading. Format: YYYY-MM-DD
	RuntimeDate *string `json:"runtimeDate,omitempty"`

	// The last 5 minute interval which was updated by the thermostat telemetry
	// update. Subtract 2 from this interval to obtain the beginning interval for
	// the last 3 readings. Multiply by 5 mins to obtain the minutes of the day.
	// Range: 0-287
	RuntimeInterval *int `json:"runtimeInterval,omitempty"`

	// The last three 5 minute actual temperature readings
	ActualTemperature []int `json:"actualTemperature,omitempty"`

	// The last three 5 minute actual humidity readings.
	ActualHumidity []int `json:"actualHumidity,omitempty"`

	// The last three 5 minute desired heat temperature readings.
	DesiredHeat []int `json:"desiredHeat,omitempty"`

	// The last three 5 minute desired cool temperature readings.
	DesiredCool []int `json:"desiredCool,omitempty"`

	// The last three 5 minute desired humidity readings.
	DesiredHumidity []int `json:"desiredHumidity,omitempty"`

	// The last three 5 minute desired de-humidification readings.
	DesiredDehumidity []int `json:"desiredDehumidity,omitempty"`

	// The last three 5 minute desired Demand Management temperature offsets. This
	// value is Demand Management adjustment value which was applied by the
	// thermostat. If the thermostat decided not to honour the adjustment, it will
	// send 0 for the interval. Compare these values with the values sent in the DM
	// message to determine whether the thermostat applied the adjustment.
	DMOffset []int `json:"dmOffset,omitempty"`

	// The last three 5 minute HVAC Mode reading. These values indicate which stage
	// was energized in the 5 minute interval. Values: heatStage10n, heatStage20n,
	// heatStage30n, heatOff, compressorCoolStage10n, compressorCoolStage20n,
	// compressorCoolOff, compressorHeatStage10n, compressorHeatStage20n,
	// compressorHeatOff, economyCycle.
	HVACMode []string `json:"hvacMode,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the heat pump stage 1 runtime.
	HeatPump1 []int `json:"heatPump1,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the heat pump stage 2 runtime.
	HeatPump2 []int `json:"heatPump2,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the auxiliary heat stage 1. If the
	// thermostat does not have a heat pump, this is heat stage 1.
	AuxHeat1 []int `json:"auxHeat1,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the auxiliary heat stage 2. If the
	// thermostat does not have a heat pump, this is heat stage 2.
	AuxHeat2 []int `json:"auxHeat2,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the heat stage 3 if the thermostat does
	// not have a heat pump. Auxiliary stage 3 is not supported.
	AuxHeat3 []int `json:"auxHeat3,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the cooling stage 1.
	Cool1 []int `json:"cool1,omitempty"`

	// The last three 5 minute HVAC Runtime values in seconds (0-300 seconds) per
	// interval. This value corresponds to the cooling stage 2.
	Cool2 []int `json:"cool2,omitempty"`

	// The last three 5 minute fan Runtime values in seconds (0-300 seconds) per
	// interval.
	Fan []int `json:"fan,omitempty"`

	// The last three 5 minute humidifier Runtime values in seconds (0-300 seconds)
	// per interval.
	Humidifier []int `json:"humidifier,omitempty"`

	// The last three 5 minute de-humidifier Runtime values in seconds (0-300
	// seconds) per interval.
	Dehumidifier []int `json:"dehumidifier,omitempty"`

	// The last three 5 minute economizer Runtime values in seconds (0-300 seconds)
	// per interval.
	Economizer []int `json:"economizer,omitempty"`

	// The last three 5 minute ventilator Runtime values in seconds (0-300 seconds)
	// per interval.
	Ventilator []int `json:"ventilator,omitempty"`

	// The latest value of the current electricity bill as interpolated from the
	// thermostat's readings from a paired electricity meter.
	CurrentElectricityBill *int `json:"currentElectricityBill,omitempty"`

	// The latest estimate of the projected electricity bill as interpolated from
	// the thermostat's readings from a paired electricity meter.
	ProjectedElectricityBill *int `json:"projectedElectricityBill,omitempty"`
}
