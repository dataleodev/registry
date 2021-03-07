package regions

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
)

var _ registry.RegionRepository = (*postgres)(nil)

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) registry.RegionRepository {
	return &postgres{
		db: db,
	}
}

func (p postgres) Get(ctx context.Context, id string) (region registry.Region, err error) {
	q := "SELECT * FROM regions WHERE id=$1"
	row := p.db.QueryRow(q, id)

	switch err := row.Scan(&region.ID, &region.Name, &region.Desc); err {

	case sql.ErrNoRows:
		message := errors.New(fmt.Sprintf("key not found %s\n",sql.ErrNoRows.Error()))
		return region,message

	case nil:
		return region, nil

	default:
		return region, err
	}
}

func (p postgres) Add(ctx context.Context, region registry.Region) (err error) {
	queryStr := "INSERT INTO regions (id,name,desc) VALUES($1,$2,$3);"
	_, err = p.db.Exec(queryStr,region.ID,region.Name,region.Desc)
	return err
}

func (p postgres) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (p postgres) List(ctx context.Context) (regions []registry.Region, err error) {
	rows, err := p.db.Query("SELECT * FROM regions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var region registry.Region

		err = rows.Scan(
			&region.ID,
			&region.Name,
			&region.Desc,

		)
		if err != nil {
			return regions, err
		}
		regions = append(regions,region)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return regions, err
	}

	return
}

func (p postgres) Update(ctx context.Context, id string, user registry.Region) (registry.Region, error) {
	panic("implement me")
}
