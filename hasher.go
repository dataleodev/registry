package registry

// Hasher specifies an API for generating hashes of an arbitrary textual
// content.
type Hasher interface {
	// Hash generates the hashed string from plain-text.
	Hash(plainText string) (hashedPassword string, err error)

	// Compare compares plain-text version to the hashed one. An error should
	// indicate failed comparison.
	Compare(hashedPassword string, plainText string) error
}
