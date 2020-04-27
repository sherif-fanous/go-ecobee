package objects

// The HouseDetails object contains contains the information about the house the
// thermostat is installed in.
//
// NOTE: the windowEfficiency field can be used to set the window efficiency of
// the house. This is set as an integer value through the API, however these
// integer values are actually converted by the thermostat to specific R values
// which are not sorted. The mapping of API integer value to R value is as
// follows:
type HouseDetails struct {
	// The style of house. Values: other, apartment, condominium, detached, loft,
	// multiPlex, rowHouse, semiDetached, townhouse, and 0 for unknown.
	Style *string `json:"style,omitempty"`

	// The size of the house in square feet.
	Size *int `json:"size,omitempty"`

	// The number of floors or levels in the house.
	NumberOfFloors *int `json:"numberOfFloors,omitempty"`

	// The number of rooms in the house.
	NumberOfRooms *int `json:"numberOfRooms,omitempty"`

	// The number of occupants living in the house.
	NumberOfOccupants *int `json:"numberOfOccupants,omitempty"`

	// The age of house in years.
	Age *int `json:"age,omitempty"`

	// This field defines the window efficiency of the house. Valid values are in
	// the range 1 - 7. Changing the value of this field alters the settings the
	// thermostat uses for the humidifier when in 'frost Control' mode. See the
	// NOTE above before updating this value.
	WindowEfficiency *int `json:"windowEfficiency,omitempty"`
}
