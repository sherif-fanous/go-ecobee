package objects

// The thermostat Program is a container for the Schedule and its Climates.
//
// See Core Concepts
// (https://www.ecobee.com/home/developer/api/introduction/core-concepts.shtml#program)
// for details on how the program is structured. The schedule property is a two
// dimensional array containing the climate names.
type Program struct {
	// The Schedule object defining the program schedule.
	Schedule [][]string `json:"schedule,omitempty"`

	// The list of Climate objects defining all the climates in the program schedule.
	Climates []Climate `json:"climates,omitempty"`

	// The currently active climate, identified by its ClimateRef.
	CurrentClimateRef *string `json:"currentClimateRef,omitempty"`
}
