package users

import (
	"context"
	"database/sql"
	"github.com/dataleodev/registry"
)

var _ registry.UserRepository = (*postgres)(nil)

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB)registry.UserRepository {
	return &postgres{
		db: db,
	}
}

func (p postgres) Get(ctx context.Context, id string) (registry.User, error) {
	panic("implement me")
}

func (p postgres) Add(ctx context.Context, user registry.User) error {
	panic("implement me")
}

func (p postgres) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (p postgres) List(ctx context.Context) ([]registry.User, error) {
	panic("implement me")
}

func (p postgres) Update(ctx context.Context, id string, user registry.User) (registry.User, error) {
	panic("implement me")
}


