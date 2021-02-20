package registry

import (
	"context"
	"github.com/dataleodev/registry/logger"
)

var (
	_ Service = (*service)(nil)
)

// Service describes the service.
type Service interface {
	AuthThing(ctx context.Context, uuid string, authToken string) (node Node, err error)

	//Register to be used by tools like web dashboards and cli tools
	//to register admins
	//name,email, password are needed, on successful registration
	//uuid v4, api key will be returned
	Register(ctx context.Context, name, email, password string) (uuid string, err error)

	//Login returns access token with a life of 20 minutes after a user has supplied
	//uuid, password correctly
	Login(ctx context.Context, uuid, password string) (token string, err error)

	// ViewUser retrieves user info for a given user ID and an authorized token.
	ViewUser(ctx context.Context, token, id string) (user User, err error)

	// ListUsers retrieves users list for a valid admin token.
	ListUsers(ctx context.Context, token string, args map[string]string) (users []User, err error)

	// UpdateUser updates the user metadata.
	UpdateUser(ctx context.Context, token string, user User) (err error)

	// ChangePassword change users password for authenticated user.
	ChangePassword(ctx context.Context, authToken, password, oldPassword string) (err error)

	AddNode(ctx context.Context, token string, node Node) (err error)

	//GetUser fetches all users details by specifying the id
	//id is the user uuid/email
	//token is a generated token/password if a user is admin
	GetNode(ctx context.Context, token string, id string) (node Node, err error)

	//ListUser returns all the list of all available users
	ListNodes(ctx context.Context, token, region string) (nodes []Node, err error)

	DeleteNode(ctx context.Context, token, id string) (err error)

	UpdateNode(ctx context.Context, token, id string, node Node) (err error)

	AddRegion(ctx context.Context, token string, region Region) (err error)

	ListRegions(ctx context.Context, token string) (regions []Region, err error)
}

type service struct {
	users   UserRepository
	nodes   NodeRepository
	regions RegionRepository
	ids     IDProvider
	hasher  Hasher
	log     logger.Logger
	auth    AuthNZ
}

func (s *service) AuthThing(ctx context.Context, uuid string, authToken string) (node Node, err error) {
	// TODO implement the business logic of AuthThing
	return node, err
}

func (s *service) Register(ctx context.Context, name string, email string, password string) (uuid string, err error) {
	// TODO implement the business logic of Register
	return uuid, err
}
func (s *service) Login(ctx context.Context, uuid string, password string) (token string, err error) {
	// TODO implement the business logic of Login
	return token, err
}
func (s *service) ViewUser(ctx context.Context, token string, id string) (user User, err error) {
	// TODO implement the business logic of ViewUser
	return user, err
}
func (s *service) ListUsers(ctx context.Context, token string, args map[string]string) (users []User, err error) {
	// TODO implement the business logic of ListUsers
	return users, err
}
func (s *service) UpdateUser(ctx context.Context, token string, user User) (err error) {
	// TODO implement the business logic of UpdateUser
	return err
}
func (s *service) ChangePassword(ctx context.Context, authToken string, password string, oldPassword string) (err error) {
	// TODO implement the business logic of ChangePassword
	return err
}
func (s *service) AddNode(ctx context.Context, token string, node Node) (err error) {
	// TODO implement the business logic of AddNode
	return err
}
func (s *service) GetNode(ctx context.Context, token string, id string) (node Node, err error) {
	// TODO implement the business logic of GetNode
	return node, err
}
func (s *service) ListNodes(ctx context.Context, token string, region string) (nodes []Node, err error) {
	// TODO implement the business logic of ListNodes
	return nodes, err
}
func (s *service) DeleteNode(ctx context.Context, token string, id string) (err error) {
	// TODO implement the business logic of DeleteNode
	return err
}
func (s *service) UpdateNode(ctx context.Context, token string, id string, node Node) (err error) {
	// TODO implement the business logic of UpdateNode
	return err
}
func (s *service) AddRegion(ctx context.Context, token string, region Region) (err error) {
	// TODO implement the business logic of AddRegion
	return err
}
func (s *service) ListRegions(ctx context.Context, token string) (regions []Region, err error) {
	// TODO implement the business logic of ListRegions
	return regions, err
}

// NewBasicRegistryService returns a naive, stateless implementation of Service.
func NewBasicRegistryService() Service {
	return &service{}
}

// New returns a Service with all of the expected middleware wired in.
func New(middleware []Middleware) Service {
	var svc Service = NewBasicRegistryService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
