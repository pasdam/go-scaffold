package core

// Filter is an interface used to filter strings
type Filter interface {

	// Accept returns true if the input string should be accepted, false otherwise
	Accept(value string) bool
}
