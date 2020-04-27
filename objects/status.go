package objects

// The response status object contains the processing status of the request. It
// will contain any relevant error information should an error occur. The status
// object is returned with every response regardless of success or failure
// status. It is suitable for logging request failures.
type Status struct {
	// The status code for this status.
	Code *int `json:"code,omitempty"`

	// The detailed message for this status.
	Message *string `json:"message,omitempty"`
}
