package keys

import (
	"context"
	"database/sql"
	"github.com/dataleodev/registry"
)

var _ registry.KeyRepository = (*postgres)(nil)

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB)registry.KeyRepository {
	return &postgres{
		db: db,
	}
}

func (p postgres) Add(ctx context.Context, key, value string) (err error) {
	panic("implement me")
}

func (p postgres) Get(ctx context.Context, key string) (value string, err error) {
	panic("implement me")
}

func (p postgres) Delete(ctx context.Context, key string) (err error) {
	panic("implement me")
}


