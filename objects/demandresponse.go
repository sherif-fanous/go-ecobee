package objects

// The demand response object contains information pertaining to a program
// change event for a set of thermostats. Demand response is only available to
// EMS and Utility accounts.
type DemandResponse struct {
	// The user defined name of the Demand Response event.
	Name *string `json:"name,omitempty"`

	// A system generated ID to uniquely identify the event.
	DemandResponseRef *string `json:"demandResponseRef,omitempty"`

	// User (event creator) comments regarding the event
	Comments *string `json:"comments,omitempty"`

	// The text message to send to the user (recipient of the event)
	Message *string `json:"message,omitempty"`

	// The date the demand response is deferred to.
	DeferredDate *string `json:"deferredDate,omitempty"`

	// The time the demand response has been deferred to.
	DeferredTime *string `json:"deferredTime,omitempty"`

	// Whether to show the event message on the thermostat. Default: true
	ShowIDT *bool `json:"showIdt,omitempty"`

	// Whether to show the event message on the web portal. Default: true
	ShowWeb *bool `json:"showWeb,omitempty"`

	// Whether to send the event message as an email to the user. Default: true
	SendEmail *bool `json:"sendEmail,omitempty"`

	// Whether to randomize the start time of the event for each thermostat in
	// order to mitigate severe electric grid load drops when the event starts.
	// Default: false
	RandomizeStartTime *bool `json:"randomizeStartTime,omitempty"`

	// The number of seconds between 0 and the value to randomize the start time
	// with. Default: 1800
	RandomStartTimeSeconds *int `json:"randomStartTimeSeconds,omitempty"`

	// Whether to randomize the end time of the event for each thermostat in order
	// to mitigate electric grid load spikes when the event ends and thermostats
	// resume normal program operation. Default: false
	RandomizeEndTime *bool `json:"randomizeEndTime,omitempty"`

	// The number of seconds between 0 and the value to randomize the end time
	// with. Default: 1800
	RandomEndTimeSeconds *int `json:"randomEndTimeSeconds,omitempty"`

	// The Event Object whose properties to use when creating the DR event.
	Event *Event `json:"event,omitempty"`

	// The list of thermostats associated with this DR. Populated by the server
	// when listing DRs.
	Thermostats []string `json:"thermostats,omitempty"`

	//
	ExternalRef *string `json:"externalRef,omitempty"`

	//
	ExternalRefType *string `json:"externalRefType,omitempty"`

	//
	Priority *int `json:"priority,omitempty"`
}
