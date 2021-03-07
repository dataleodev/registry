package nodes

import (
	"context"
	"database/sql"
	"github.com/dataleodev/registry"
)

var _ registry.NodeRepository = (*postgres)(nil)

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) registry.NodeRepository {
	return &postgres{
		db: db,
	}
}

func (p postgres) Get(ctx context.Context, id string) (registry.Node, error) {
	panic("implement me")
}

func (p postgres) Add(ctx context.Context, user registry.Node) error {
	panic("implement me")
}

func (p postgres) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (p postgres) List(ctx context.Context) ([]registry.Node, error) {
	panic("implement me")
}

func (p postgres) Update(ctx context.Context, id string, user registry.Node) (registry.Node, error) {
	panic("implement me")
}
