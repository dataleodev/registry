package postgres

import (
	"database/sql"
	"fmt"
)

//hostname = "localhost"
//	port     = "5432"
//	user     = "postgres"
//	password = "postgres"
//	dbname   = "postgres"
//	sslmode  = "disable"
type DBConfig struct {
	Hostname string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func New(cfg DBConfig) (db *sql.DB, err error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Hostname, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SslMode)
	db, err = sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return
}
