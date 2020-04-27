package objects

// The Weather Forecast contains the weather forecast information for the
// thermostat. The first forecast is the most accurate, later forecasts become
// less accurate in distance and time.
type WeatherForecast struct {
	// The Integer value used to map to a weatherSymbol. See list of mappings above.
	WeatherSymbol *int `json:"weatherSymbol,omitempty"`

	// The time stamp of the weather forecast.
	DateTime *string `json:"dateTime,omitempty"`

	// A text value representing the current weather condition.
	Condition *string `json:"condition,omitempty"`

	// The current temperature.
	Temperature *int `json:"temperature,omitempty"`

	// The current barometric pressure.
	Pressure *int `json:"pressure,omitempty"`

	// The current humidity.
	RelativeHumidity *int `json:"relativeHumidity,omitempty"`

	// The dewpoint.
	Dewpoint *int `json:"dewpoint,omitempty"`

	// The visibility in meters; 0 - 70,000.
	Visibility *int `json:"visibility,omitempty"`

	// The wind speed as an integer in mph * 1000.
	WindSpeed *int `json:"windSpeed,omitempty"`

	// The wind gust speed.
	WindGust *int `json:"windGust,omitempty"`

	// The wind direction.
	WindDirection *string `json:"windDirection,omitempty"`

	// The wind bearing.
	WindBearing *int `json:"windBearing,omitempty"`

	// Probability of precipitation.
	Pop *int `json:"pop,omitempty"`

	// The predicted high temperature for the day.
	TempHigh *int `json:"tempHigh,omitempty"`

	// The predicted low temperature for the day.
	TempLow *int `json:"tempLow,omitempty"`

	// The cloud cover condition.
	Sky *int `json:"sky,omitempty"`
}
