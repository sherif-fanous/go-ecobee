package objects

// An output is a relay connected to the thermostat.
type Output struct {
	// The name of the outpute
	Name *string `json:"name,omitempty"`

	// The thermostat zone the output is associated with
	Zone *int `json:"zone,omitempty"`

	// The unique output identifier number.
	OutputID *int `json:"outputId,omitempty"`

	// The type of output. Values: compressor1, compressor2, dehumidifier,
	// economizer, fan, heat1, heat2, heat3, heatPumpReversal, humidifer, none,
	// occupancy, userDefined, ventilator, zoneCool, zoneFan, zoneHeat.
	Type *string `json:"type,omitempty"`

	// Whether to send an update message.
	SendUpdate *bool `json:"sendUpdate,omitempty"`

	// If true, when this output is activated it will close the relay. Otherwise,
	// activating the relay will open the relay.
	ActiveClosed *bool `json:"activeClosed,omitempty"`

	// Time to activate relay - in seconds.
	ActivationTime *int `json:"activationTime,omitempty"`

	// Time to deactivate relay - in seconds.
	DeactivationTime *int `json:"deactivationTime,omitempty"`
}
