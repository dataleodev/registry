package users

import (
	"context"
	"database/sql"
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/logger"
	"github.com/dataleodev/registry/pkg/errors"
	regsql "github.com/dataleodev/registry/sql"
	"os"
	"time"
)

var _ registry.UserRepository = (*postgres)(nil)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUserNotUpdated = errors.New("user not updated")
)

type dbUser struct {
	ID       string `json:"id,omitempty"`                  //id or user token | uuid
	Name     string `json:"name"`                          //fullname
	Email    string `json:"email"`                         //email
	Password string `json:"password,omitempty"`            //password of user
	Region   string `json:"region_of_operation,omitempty"` //operating region in case of multi cloud
	Created  string `json:"created,omitempty"`
}

func (u dbUser) toUser() registry.User {
	return registry.User{
		UUID:     u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Region:   u.Region,
		Created:  u.Created,
	}
}

func fromUser(user registry.User) (dbUser, error) {

	//	now, err := time.Parse(time.RFC3339, user.Created)
	now := time.Now().Format(time.RFC3339)

	//	if err != nil {
	//		return dbUser{}, err
	//	}
	return dbUser{
		ID:       user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Region:   user.Region,
		Created:  now,
	}, nil
}

type postgres struct {
	db       *sql.DB
	dbLogger logger.Logger
}

func NewRepository(db *sql.DB) registry.UserRepository {
	dlog, _ := logger.New(os.Stdout, "debug")
	return &postgres{
		db:       db,
		dbLogger: dlog,
	}
}

func (p postgres) Get(ctx context.Context, id string) (registry.User, error) {

	row := p.db.QueryRow(regsql.UserSelectById, id)
	dUser := dbUser{}

	switch err := row.Scan(
		&dUser.ID, &dUser.Name, &dUser.Email,
		&dUser.Region, &dUser.Password, &dUser.Created); err {

	case sql.ErrNoRows:
		return registry.User{}, ErrUserNotFound

	case nil:
		return dUser.toUser(), nil

	default:
		return registry.User{}, err
	}
}

func (p postgres) Add(ctx context.Context, user registry.User) error {
	// id VARCHAR(100) UNIQUE NOT NULL PRIMARY KEY,
	//    name VARCHAR(100) UNIQUE NOT NULL,
	//    email VARCHAR(100) UNIQUE NOT NULL,
	//    region VARCHAR(100) UNIQUE NOT NULL,
	//    password VARCHAR(100) UNIQUE NOT NULL,
	//    created VARCHAR(100) UNIQUE NOT NULL

	dUser, err := fromUser(user)

	if err != nil {
		return err
	}

	_, err = p.db.Exec(regsql.UserInsertNew,
		dUser.ID, dUser.Name, dUser.Email, dUser.Region, dUser.Password, dUser.Created)

	if err != nil {
		return err
	}

	return nil
}

func (p postgres) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (p postgres) List(ctx context.Context) (users []registry.User, err error) {
	rows, err := p.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user registry.User

		err = rows.Scan(
			&user.UUID,
			&user.Name,
			&user.Email,
			&user.Region,
			&user.Password,
			&user.Created,
		)
		if err != nil {
			return users, err
		}
		users = append(users,user)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return users, err
	}

	return
}

func (p postgres) Update(ctx context.Context, id string, user registry.User) (registry.User, error) {
	panic("implement me")
}
