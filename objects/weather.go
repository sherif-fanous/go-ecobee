package objects

// The Weather object contains the weather and forecast information for the
// thermostat's location.
type Weather struct {
	// The time stamp in UTC of the weather forecast.
	Timestamp *string `json:"timestamp,omitempty"`

	// The weather station identifier.
	WeatherStation *string `json:"weatherStation,omitempty"`

	// The list of latest weather station forecasts.
	Forecasts []WeatherForecast `json:"forecasts,omitempty"`
}
