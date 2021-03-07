package keys

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
)

var _ registry.KeyRepository = (*postgres)(nil)

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) registry.KeyRepository {
	return &postgres{
		db: db,
	}
}

func (p postgres) Add(ctx context.Context, key, value string) (err error) {
	queryStr := "INSERT INTO keys (id,value) VALUES($1,$2);"
	_, err = p.db.Exec(queryStr,key,value)
	return err
}

func (p postgres) Get(ctx context.Context, key string) (value string, err error) {
	q := "SELECT value FROM keys WHERE id=$1;"
	row := p.db.QueryRow(q, key)

	switch err := row.Scan(&value); err {

	case sql.ErrNoRows:
		message := errors.New(fmt.Sprintf("key not found %s\n",sql.ErrNoRows.Error()))
		return "",message

	case nil:
		return "", nil

	default:
		return value, err
	}
}

func (p postgres) Delete(ctx context.Context, key string) (err error) {
	q := "DELETE FROM keys WHERE id=$1;"
	_,err = p.db.Exec(q,key)
	return
}
