package ulid

import (
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
	"github.com/oklog/ulid/v2"
	mathrand "math/rand"
	"time"
)

// ErrGeneratingID indicates error in generating ULID
var ErrGeneratingID = errors.New("generating id failed")

var _ registry.IDProvider = (*ulidProvider)(nil)

type ulidProvider struct {
	entropy *mathrand.Rand
}

// New instantiates a ULID provider.
func New() registry.IDProvider {
	seed := time.Now().UnixNano()
	source := mathrand.NewSource(seed)
	return &ulidProvider{
		entropy: mathrand.New(source),
	}
}

func (up *ulidProvider) ID() (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), up.entropy)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
