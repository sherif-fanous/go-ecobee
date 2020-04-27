package objects

// The function object is defined by its functionType and one or more additional
// properties. The property list is variable depending on the type of function.
// Functions are used to perform more complex operations on a thermostat or user
// which are too complex with simple property modifications. Functions are used
// to modify read-only objects where appropriate.
//
// Each function takes different parameters, the object format is below:
type Function struct {
	// The function type name. See the type name in the function documentation.
	Type *string `json:"type,omitempty"`

	// A map of key=value pairs as the parameters to the function. See individual
	// function documentation for the properties.
	Params map[string]interface{} `json:"params,omitempty"`
}
