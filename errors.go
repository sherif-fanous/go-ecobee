package ecobee

// APIError describes errors returned by the ecobee server while making
// requests
type APIError struct {
	errorString string
}

// Error returns the string representation of an APIError.
func (e *APIError) Error() string {
	return e.errorString
}

// AuthorizationError describes errors returned by the ecobee server while
// authorizing
type AuthorizationError struct {
	errorString string
}

// Error returns the string representation of an AuthorizationError.
func (e *AuthorizationError) Error() string {
	return e.errorString
}
