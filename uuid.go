package registry

// IDProvider specifies an API for generating unique identifiers.
type IDProvider interface {
	// ID generates the unique identifier.
	ID() (string, error)
}
