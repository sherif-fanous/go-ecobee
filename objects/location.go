package objects

// The Location describes the physical location and coordinates of the
// thermostat as entered by the thermostat owner. The address information is
// used in a geocode look up to obtain the thermostat coordinates. The
// coordinates are used to obtain accurate weather information.
type Location struct {
	// The timezone offset in minutes from UTC.
	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"`

	// The Olson timezone the thermostat resides in (e.g America/Toronto).
	TimeZone *string `json:"timeZone,omitempty"`

	// Whether the thermostat should factor in daylight savings when displaying the
	// date and time.
	IsDaylightSaving *bool `json:"isDaylightSaving,omitempty"`

	// The thermostat location street address.
	StreetAddress *string `json:"streetAddress,omitempty"`

	// The thermostat location city.
	City *string `json:"city,omitempty"`

	// The thermostat location State or Province.
	ProvinceState *string `json:"provinceState,omitempty"`

	// The thermostat location country.
	Country *string `json:"country,omitempty"`

	// The thermostat location ZIP or Postal code.
	PostalCode *string `json:"postalCode,omitempty"`

	// The thermostat owner's phone number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// The lat/long geographic coordinates of the thermostat location.
	MapCoordinates *string `json:"mapCoordinates,omitempty"`
}
