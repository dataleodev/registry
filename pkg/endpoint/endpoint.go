package endpoint

import (
	"context"
	service "github.com/dataleodev/registry/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	Uuid string `json:"uuid"`
	Err  error  `json:"err"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		uuid, err := s.Register(ctx, req.Name, req.Email, req.Password)
		return RegisterResponse{
			Err:  err,
			Uuid: uuid,
		}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.Err
}

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Uuid     string `json:"uuid"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := s.Login(ctx, req.Uuid, req.Password)
		return LoginResponse{
			Err:   err,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// ViewUserRequest collects the request parameters for the ViewUser method.
type ViewUserRequest struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

// ViewUserResponse collects the response parameters for the ViewUser method.
type ViewUserResponse struct {
	User service.User `json:"user"`
	Err  error        `json:"err"`
}

// MakeViewUserEndpoint returns an endpoint that invokes ViewUser on the service.
func MakeViewUserEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewUserRequest)
		user, err := s.ViewUser(ctx, req.Token, req.Id)
		return ViewUserResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r ViewUserResponse) Failed() error {
	return r.Err
}

// ListUsersRequest collects the request parameters for the ListUsers method.
type ListUsersRequest struct {
	Token string            `json:"token"`
	Args  map[string]string `json:"args"`
}

// ListUsersResponse collects the response parameters for the ListUsers method.
type ListUsersResponse struct {
	Users []User `json:"users"`
	Err   error  `json:"err"`
}

// MakeListUsersEndpoint returns an endpoint that invokes ListUsers on the service.
func MakeListUsersEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListUsersRequest)
		users, err := s.ListUsers(ctx, req.Token, req.Args)
		return ListUsersResponse{
			Err:   err,
			Users: users,
		}, nil
	}
}

// Failed implements Failer.
func (r ListUsersResponse) Failed() error {
	return r.Err
}

// UpdateUserRequest collects the request parameters for the UpdateUser method.
type UpdateUserRequest struct {
	Token string       `json:"token"`
	User  service.User `json:"user"`
}

// UpdateUserResponse collects the response parameters for the UpdateUser method.
type UpdateUserResponse struct {
	Err error `json:"err"`
}

// MakeUpdateUserEndpoint returns an endpoint that invokes UpdateUser on the service.
func MakeUpdateUserEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := s.UpdateUser(ctx, req.Token, req.User)
		return UpdateUserResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserResponse) Failed() error {
	return r.Err
}

// ChangePasswordRequest collects the request parameters for the ChangePassword method.
type ChangePasswordRequest struct {
	AuthToken   string `json:"auth_token"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
}

// ChangePasswordResponse collects the response parameters for the ChangePassword method.
type ChangePasswordResponse struct {
	Err error `json:"err"`
}

// MakeChangePasswordEndpoint returns an endpoint that invokes ChangePassword on the service.
func MakeChangePasswordEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangePasswordRequest)
		err := s.ChangePassword(ctx, req.AuthToken, req.Password, req.OldPassword)
		return ChangePasswordResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r ChangePasswordResponse) Failed() error {
	return r.Err
}

// AddNodeRequest collects the request parameters for the AddNode method.
type AddNodeRequest struct {
	Token string       `json:"token"`
	Node  service.Node `json:"node"`
}

// AddNodeResponse collects the response parameters for the AddNode method.
type AddNodeResponse struct {
	Err error `json:"err"`
}

// MakeAddNodeEndpoint returns an endpoint that invokes AddNode on the service.
func MakeAddNodeEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNodeRequest)
		err := s.AddNode(ctx, req.Token, req.Node)
		return AddNodeResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddNodeResponse) Failed() error {
	return r.Err
}

// GetNodeRequest collects the request parameters for the GetNode method.
type GetNodeRequest struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

// GetNodeResponse collects the response parameters for the GetNode method.
type GetNodeResponse struct {
	Node service.Node `json:"node"`
	Err  error        `json:"err"`
}

// MakeGetNodeEndpoint returns an endpoint that invokes GetNode on the service.
func MakeGetNodeEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNodeRequest)
		node, err := s.GetNode(ctx, req.Token, req.Id)
		return GetNodeResponse{
			Err:  err,
			Node: node,
		}, nil
	}
}

// Failed implements Failer.
func (r GetNodeResponse) Failed() error {
	return r.Err
}

// ListNodesRequest collects the request parameters for the ListNodes method.
type ListNodesRequest struct {
	Token  string `json:"token"`
	Region string `json:"region"`
}

// ListNodesResponse collects the response parameters for the ListNodes method.
type ListNodesResponse struct {
	Nodes []Node `json:"nodes"`
	Err   error  `json:"err"`
}

// MakeListNodesEndpoint returns an endpoint that invokes ListNodes on the service.
func MakeListNodesEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListNodesRequest)
		nodes, err := s.ListNodes(ctx, req.Token, req.Region)
		return ListNodesResponse{
			Err:   err,
			Nodes: nodes,
		}, nil
	}
}

// Failed implements Failer.
func (r ListNodesResponse) Failed() error {
	return r.Err
}

// DeleteNodeRequest collects the request parameters for the DeleteNode method.
type DeleteNodeRequest struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

// DeleteNodeResponse collects the response parameters for the DeleteNode method.
type DeleteNodeResponse struct {
	Err error `json:"err"`
}

// MakeDeleteNodeEndpoint returns an endpoint that invokes DeleteNode on the service.
func MakeDeleteNodeEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteNodeRequest)
		err := s.DeleteNode(ctx, req.Token, req.Id)
		return DeleteNodeResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteNodeResponse) Failed() error {
	return r.Err
}

// UpdateNodeRequest collects the request parameters for the UpdateNode method.
type UpdateNodeRequest struct {
	Token string       `json:"token"`
	Id    string       `json:"id"`
	Node  service.Node `json:"node"`
}

// UpdateNodeResponse collects the response parameters for the UpdateNode method.
type UpdateNodeResponse struct {
	Err error `json:"err"`
}

// MakeUpdateNodeEndpoint returns an endpoint that invokes UpdateNode on the service.
func MakeUpdateNodeEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateNodeRequest)
		err := s.UpdateNode(ctx, req.Token, req.Id, req.Node)
		return UpdateNodeResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r UpdateNodeResponse) Failed() error {
	return r.Err
}

// AddRegionRequest collects the request parameters for the AddRegion method.
type AddRegionRequest struct {
	Token  string         `json:"token"`
	Region service.Region `json:"region"`
}

// AddRegionResponse collects the response parameters for the AddRegion method.
type AddRegionResponse struct {
	Err error `json:"err"`
}

// MakeAddRegionEndpoint returns an endpoint that invokes AddRegion on the service.
func MakeAddRegionEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRegionRequest)
		err := s.AddRegion(ctx, req.Token, req.Region)
		return AddRegionResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddRegionResponse) Failed() error {
	return r.Err
}

// ListRegionsRequest collects the request parameters for the ListRegions method.
type ListRegionsRequest struct {
	Token string `json:"token"`
}

// ListRegionsResponse collects the response parameters for the ListRegions method.
type ListRegionsResponse struct {
	Regions []Region `json:"regions"`
	Err     error    `json:"err"`
}

// MakeListRegionsEndpoint returns an endpoint that invokes ListRegions on the service.
func MakeListRegionsEndpoint(s service.RegistryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRegionsRequest)
		regions, err := s.ListRegions(ctx, req.Token)
		return ListRegionsResponse{
			Err:     err,
			Regions: regions,
		}, nil
	}
}

// Failed implements Failer.
func (r ListRegionsResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, name string, email string, password string) (uuid string, err error) {
	request := RegisterRequest{
		Email:    email,
		Name:     name,
		Password: password,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).Uuid, response.(RegisterResponse).Err
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, uuid string, password string) (token string, err error) {
	request := LoginRequest{
		Password: password,
		Uuid:     uuid,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Token, response.(LoginResponse).Err
}

// ViewUser implements Service. Primarily useful in a client.
func (e Endpoints) ViewUser(ctx context.Context, token string, id string) (user service.User, err error) {
	request := ViewUserRequest{
		Id:    id,
		Token: token,
	}
	response, err := e.ViewUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ViewUserResponse).User, response.(ViewUserResponse).Err
}

// ListUsers implements Service. Primarily useful in a client.
func (e Endpoints) ListUsers(ctx context.Context, token string, args map[string]string) (users []User, err error) {
	request := ListUsersRequest{
		Args:  args,
		Token: token,
	}
	response, err := e.ListUsersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListUsersResponse).Users, response.(ListUsersResponse).Err
}

// UpdateUser implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUser(ctx context.Context, token string, user service.User) (err error) {
	request := UpdateUserRequest{
		Token: token,
		User:  user,
	}
	response, err := e.UpdateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserResponse).Err
}

// ChangePassword implements Service. Primarily useful in a client.
func (e Endpoints) ChangePassword(ctx context.Context, authToken string, password string, oldPassword string) (err error) {
	request := ChangePasswordRequest{
		AuthToken:   authToken,
		OldPassword: oldPassword,
		Password:    password,
	}
	response, err := e.ChangePasswordEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ChangePasswordResponse).Err
}

// AddNode implements Service. Primarily useful in a client.
func (e Endpoints) AddNode(ctx context.Context, token string, node service.Node) (err error) {
	request := AddNodeRequest{
		Node:  node,
		Token: token,
	}
	response, err := e.AddNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddNodeResponse).Err
}

// GetNode implements Service. Primarily useful in a client.
func (e Endpoints) GetNode(ctx context.Context, token string, id string) (node service.Node, err error) {
	request := GetNodeRequest{
		Id:    id,
		Token: token,
	}
	response, err := e.GetNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetNodeResponse).Node, response.(GetNodeResponse).Err
}

// ListNodes implements Service. Primarily useful in a client.
func (e Endpoints) ListNodes(ctx context.Context, token string, region string) (nodes []Node, err error) {
	request := ListNodesRequest{
		Region: region,
		Token:  token,
	}
	response, err := e.ListNodesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListNodesResponse).Nodes, response.(ListNodesResponse).Err
}

// DeleteNode implements Service. Primarily useful in a client.
func (e Endpoints) DeleteNode(ctx context.Context, token string, id string) (err error) {
	request := DeleteNodeRequest{
		Id:    id,
		Token: token,
	}
	response, err := e.DeleteNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteNodeResponse).Err
}

// UpdateNode implements Service. Primarily useful in a client.
func (e Endpoints) UpdateNode(ctx context.Context, token string, id string, node service.Node) (err error) {
	request := UpdateNodeRequest{
		Id:    id,
		Node:  node,
		Token: token,
	}
	response, err := e.UpdateNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateNodeResponse).Err
}

// AddRegion implements Service. Primarily useful in a client.
func (e Endpoints) AddRegion(ctx context.Context, token string, region service.Region) (err error) {
	request := AddRegionRequest{
		Region: region,
		Token:  token,
	}
	response, err := e.AddRegionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddRegionResponse).Err
}

// ListRegions implements Service. Primarily useful in a client.
func (e Endpoints) ListRegions(ctx context.Context, token string) (regions []Region, err error) {
	request := ListRegionsRequest{Token: token}
	response, err := e.ListRegionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListRegionsResponse).Regions, response.(ListRegionsResponse).Err
}
