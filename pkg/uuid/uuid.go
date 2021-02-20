package uuid

import (
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
	"github.com/gofrs/uuid"
)

// ErrGeneratingID indicates error in generating UUID
var ErrGeneratingID = errors.New("generating id failed")

var _ beanpay.IDProvider = (*uuidProvider)(nil)

type uuidProvider struct{}

// New instantiates a UUID provider.
func New() beanpay.IDProvider {
	return &uuidProvider{}
}

func (up *uuidProvider) ID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(ErrGeneratingID, err)
	}

	return id.String(), nil
}
