package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/dataleodev/registry/api/endpoint"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/register").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)))
}

// decodeRegisterRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRegisterRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRegisterResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRegisterResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/login").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)))
}

// decodeLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeViewUserHandler creates the handler logic
func makeViewUserHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/view-user").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ViewUserEndpoint, decodeViewUserRequest, encodeViewUserResponse, options...)))
}

// decodeViewUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeViewUserRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ViewUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeViewUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeViewUserResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListUsersHandler creates the handler logic
func makeListUsersHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/list-users").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListUsersEndpoint, decodeListUsersRequest, encodeListUsersResponse, options...)))
}

// decodeListUsersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListUsersRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListUsersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListUsersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListUsersResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserHandler creates the handler logic
func makeUpdateUserHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/update-user").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, encodeUpdateUserResponse, options...)))
}

// decodeUpdateUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangePasswordHandler creates the handler logic
func makeChangePasswordHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/change-password").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ChangePasswordEndpoint, decodeChangePasswordRequest, encodeChangePasswordResponse, options...)))
}

// decodeChangePasswordRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangePasswordRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ChangePasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangePasswordResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangePasswordResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddNodeHandler creates the handler logic
func makeAddNodeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/add-node").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AddNodeEndpoint, decodeAddNodeRequest, encodeAddNodeResponse, options...)))
}

// decodeAddNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddNodeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AddNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddNodeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetNodeHandler creates the handler logic
func makeGetNodeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/get-node").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetNodeEndpoint, decodeGetNodeRequest, encodeGetNodeResponse, options...)))
}

// decodeGetNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetNodeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetNodeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListNodesHandler creates the handler logic
func makeListNodesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/list-nodes").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListNodesEndpoint, decodeListNodesRequest, encodeListNodesResponse, options...)))
}

// decodeListNodesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListNodesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListNodesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListNodesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListNodesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteNodeHandler creates the handler logic
func makeDeleteNodeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/delete-node").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteNodeEndpoint, decodeDeleteNodeRequest, encodeDeleteNodeResponse, options...)))
}

// decodeDeleteNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteNodeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteNodeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateNodeHandler creates the handler logic
func makeUpdateNodeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/update-node").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateNodeEndpoint, decodeUpdateNodeRequest, encodeUpdateNodeResponse, options...)))
}

// decodeUpdateNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateNodeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateNodeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddRegionHandler creates the handler logic
func makeAddRegionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/add-region").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AddRegionEndpoint, decodeAddRegionRequest, encodeAddRegionResponse, options...)))
}

// decodeAddRegionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddRegionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AddRegionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddRegionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddRegionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListRegionsHandler creates the handler logic
func makeListRegionsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/list-regions").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListRegionsEndpoint, decodeListRegionsRequest, encodeListRegionsResponse, options...)))
}

// decodeListRegionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListRegionsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListRegionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListRegionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListRegionsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
