package ecobee

// Bool is a helper function that returns a pointer to a bool value
func Bool(b bool) *bool {
	return &b
}

// Int is a helper function that returns a pointer to an int value
func Int(i int) *int {
	return &i
}

// String is a helper function that returns a pointer to a string value
func String(s string) *string {
	return &s
}
