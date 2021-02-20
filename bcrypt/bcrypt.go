package bcrypt

import (
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const cost int = 10

var (
	errHashPassword    = errors.New("Generate hash from password failed")
	errComparePassword = errors.New("Compare hash and password failed")
)

var _ registry.Hasher = (*bcryptHasher)(nil)

type bcryptHasher struct{}

// New instantiates a bcrypt-based hasher implementation.
func New() registry.Hasher {
	return &bcryptHasher{}
}

func (bh *bcryptHasher) Hash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", errors.Wrap(errHashPassword, err)
	}

	return string(hash), nil
}

func (bh *bcryptHasher) Compare(plain, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return errors.Wrap(errComparePassword, err)
	}
	return nil
}
