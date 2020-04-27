package objects

// The Audio object contains all the audio properties of the thermostat (only
// applicable to ecobee4).
type Audio struct {
	// The volume level for audio playback. This includes volume of the voice
	// assistant. A value between 0 and 100.
	PlaybackVolume *int `json:"playbackVolume,omitempty"`

	// Turn microphone (privacy mode) on and off.
	MicrophoneEnabled *bool `json:"microphoneEnabled,omitempty"`

	// The volume level for alerts on the thermostat. A value between 0 and 10,
	// with 0 meaning 'off' - the zero value may not be honored by all ecobee
	// versions.
	SoundAlertVolume *int `json:"soundAlertVolume,omitempty"`

	// The volume level for key presses on the thermostat. A value between 0 and
	// 10, with 0 meaning 'off' - the zero value may not be honored by all ecobee
	// versions.
	SoundTickVolume *int `json:"soundTickVolume,omitempty"`

	// The list of voice engines compatible with the selected thermostat.
	VoiceEngines []VoiceEngine `json:"voiceEngines,omitempty"`
}
