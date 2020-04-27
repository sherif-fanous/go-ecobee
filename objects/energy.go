package objects

// This object is currently undocumented by ecobee.
type Energy struct {
	TimeOfUse          *TimeOfUse `json:"tou,omitempty"`
	EnergyFeatureState *string    `json:"energyFeatureState,omitempty"`
	FeelsLikeMode      *string    `json:"feelsLikeMode,omitempty"`
	ComfortPreferences *string    `json:"comfortPreferences,omitempty"`
}
