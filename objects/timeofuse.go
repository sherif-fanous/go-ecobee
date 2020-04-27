package objects

// This object is currently undocumented by ecobee.
type TimeOfUse struct {
	FeatureState *string `json:"featureState,omitempty"`
	Savings      *string `json:"savings,omitempty"`
}
