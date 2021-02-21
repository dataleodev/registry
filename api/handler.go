package api

import (
	"context"
	"encoding/json"
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func MakeHTTPHandler(svc registry.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(svc)

	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(ErrorEncoder),
	}

	//POST /register
	r.Methods(http.MethodPost,http.MethodPut).Path("/register").Handler(kithttp.NewServer(
		e.RegisterEndpoint,
		decodeRegisterRequest,
		encodeRegisterResponse,
		options...,
	))

	//GET /login
	r.Methods(http.MethodGet).Path("/login").Handler(kithttp.NewServer(
		e.LoginEndpoint,
		decodeLoginRequest,
		encodeLoginResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/auth").Handler(kithttp.NewServer(
		e.AuthThingEndpoint,
		decodeAuthThingRequest,
		encodeAuthThingResponse,
		options ...,
	))


	//GET
	r.Methods(http.MethodGet).Path("/users/{id}").Handler(kithttp.NewServer(
		e.ViewUserEndpoint,
		decodeViewUserRequest,
		encodeViewUserResponse,
		options...,
	))

	//GET
	r.Methods(http.MethodGet).Path("/users").Handler(kithttp.NewServer(
		e.ListUsersEndpoint,
		decodeListUsersRequest,
		encodeListUsersResponse,
		options...,
	))

	return r
}

// decodeRegisterRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRegisterResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	uuid, password, ok := r.BasicAuth()
	if !ok{
		return LoginRequest{}, errors.New("unauthorized")
	}
	req := LoginRequest{
		Uuid:     uuid,
		Password: password,
	}
	return req, nil
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}


// decodeAuthThingRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAuthThingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := AuthThingRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAuthThingResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAuthThingResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeViewUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeViewUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok{
		return ViewUserRequest{}, errors.New("err bad routing")
	}
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	req := ViewUserRequest{
		Token: reqToken,
		Id:    id,
	}
	return req, nil
}

// encodeViewUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeViewUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

type regionReq struct {
	Region string `json:"region"`
}


// decodeListUsersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	userRegion :=  regionReq{}
	err := json.NewDecoder(r.Body).Decode(&userRegion)
	if err != nil {
		return nil, err
	}
	args := map[string]string{"region":userRegion.Region}
	req := ListUsersRequest{
		Token: reqToken,
		Args:  args,
	}

	return req, nil
}

// encodeListUsersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListUsersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}



func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

