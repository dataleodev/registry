package nodes

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
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

func (p postgres) Get(ctx context.Context, id string) (node registry.Node,err error) {
	q := "SELECT * FROM nodes WHERE uuid=$1"
	row := p.db.QueryRow(q, id)

	switch err := row.Scan(
		&node.UUID,
		&node.Addr,
		&node.Key,
		&node.Name,
		&node.Type,
		&node.Region,
		&node.Latd,
		&node.Long,
		&node.Created,
		&node.Master,
		); err {

	case sql.ErrNoRows:
		message := errors.New(fmt.Sprintf("key not found %s\n",sql.ErrNoRows.Error()))
		return node,message

	case nil:
		return node, nil

	default:
		return node, err
	}

}

func (p postgres) Add(ctx context.Context, node registry.Node) (err error){
	queryStr :=
		"INSERT INTO nodes (uuid,addr,key,name,type,region,latd,long,created, master) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);"
	_, err = p.db.Exec(queryStr,
		node.UUID,
		node.Addr,
		node.Key,
		node.Name,
		node.Type,
		node.Region,
		node.Latd,
		node.Long,
		node.Created,
		node.Master,
		)
	return err
}

func (p postgres) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (p postgres) List(ctx context.Context) (nodes []registry.Node, err error) {
	rows, err := p.db.Query("SELECT * FROM nodes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var node registry.Node

		err = rows.Scan(
			&node.UUID,
			&node.Addr,
			&node.Key,
			&node.Name,
			&node.Type,
			&node.Region,
			&node.Latd,
			&node.Long,
			&node.Created,
			&node.Master,
		)
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes,node)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nodes, err
	}

	return
}

func (p postgres) Update(ctx context.Context, id string, user registry.Node) (registry.Node, error) {
	panic("implement me")
}
