package postgres

import (
	"database/sql"
)


type DBConfig struct {

}

func NewDatabase(config DBConfig) (db *sql.DB,err error) {
	db,err = sql.Open("","")
	return
}




