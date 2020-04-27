package objects

// A sensor action defines an action to take when a SensorState is triggered.
type Action struct {
	// Values: activateRelay, adjustTemp, doNothing, shutdownAC, shutdownAuxHeat,
	// shutdownSystem, shutdownCompression, switchToOccupied, switchToUnoccupied,
	// turnOffDehumidifer, turnOffHumidifier, turnOnCool, turnOnDehumidifier,
	// turnOnFan, turnOnHeat, turnOnHumidifier.
	Type *string `json:"type,omitempty"`

	// Flag to enable an alert to be generated when the state is triggered
	SendAlert *bool `json:"sendAlert,omitempty"`

	// Whether to send an update message.
	SendUpdate *bool `json:"sendUpdate,omitempty"`

	// Delay in seconds before the action is triggered by the state change.
	ActivationDelay *int `json:"activationDelay,omitempty"`

	// The amount of time to wait before deactivating this state after the state
	// has been cleared.
	DeactivationDelay *int `json:"deactivationDelay,omitempty"`

	// The minimum length of time to maintain action after sensor has been deactivated.
	MinActionDuration *int `json:"minActionDuration,omitempty"`

	// The amount to increase/decrease current setpoint if the type = adjustTemp.
	HeatAdjustTemp *int `json:"heatAdjustTemp,omitempty"`

	// The amount to increase/decrease current setpoint if the type = adjustTemp.
	CoolAdjustTemp *int `json:"coolAdjustTemp,omitempty"`

	// The user defined relay to be activated, only used for type == activateRelay.
	ActivateRelay *string `json:"activateRelay,omitempty"`

	// Select if relay is to be open or closed when activated, only used for type
	// == activateRelay.
	ActivateRelayOpen *bool `json:"activateRelayOpen,omitempty"`
}
