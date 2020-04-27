package objects

// The Voice Engine object contains information about the voice assistant that
// the selected thermostat supports.
type VoiceEngine struct {
	// The name of the voice engine.
	Name *string `json:"name,omitempty"`

	// True if the voice engine is currently enabled (paired) for the ecobee
	// account of selected thermostat. You can change the flag value by using
	// UnlinkVoiceEngine thermostat function.
	Enabled *bool `json:"enabled,omitempty"`
}
