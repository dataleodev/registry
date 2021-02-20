package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dataleodev/registry"
	endpoint1 "github.com/dataleodev/registry/api/endpoint"
	http2 "github.com/dataleodev/registry/api/http"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (registry.Service, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var registerEndpoint endpoint.Endpoint
	{
		registerEndpoint = http.NewClient("POST", copyURL(u, "/register"), encodeHTTPGenericRequest, decodeRegisterResponse, options["Register"]...).Endpoint()
	}

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = http.NewClient("POST", copyURL(u, "/login"), encodeHTTPGenericRequest, decodeLoginResponse, options["Login"]...).Endpoint()
	}

	var viewUserEndpoint endpoint.Endpoint
	{
		viewUserEndpoint = http.NewClient("POST", copyURL(u, "/view-user"), encodeHTTPGenericRequest, decodeViewUserResponse, options["ViewUser"]...).Endpoint()
	}

	var listUsersEndpoint endpoint.Endpoint
	{
		listUsersEndpoint = http.NewClient("POST", copyURL(u, "/list-users"), encodeHTTPGenericRequest, decodeListUsersResponse, options["ListUsers"]...).Endpoint()
	}

	var updateUserEndpoint endpoint.Endpoint
	{
		updateUserEndpoint = http.NewClient("POST", copyURL(u, "/update-user"), encodeHTTPGenericRequest, decodeUpdateUserResponse, options["UpdateUser"]...).Endpoint()
	}

	var changePasswordEndpoint endpoint.Endpoint
	{
		changePasswordEndpoint = http.NewClient("POST", copyURL(u, "/change-password"), encodeHTTPGenericRequest, decodeChangePasswordResponse, options["ChangePassword"]...).Endpoint()
	}

	var addNodeEndpoint endpoint.Endpoint
	{
		addNodeEndpoint = http.NewClient("POST", copyURL(u, "/add-node"), encodeHTTPGenericRequest, decodeAddNodeResponse, options["AddNode"]...).Endpoint()
	}

	var getNodeEndpoint endpoint.Endpoint
	{
		getNodeEndpoint = http.NewClient("POST", copyURL(u, "/get-node"), encodeHTTPGenericRequest, decodeGetNodeResponse, options["GetNode"]...).Endpoint()
	}

	var listNodesEndpoint endpoint.Endpoint
	{
		listNodesEndpoint = http.NewClient("POST", copyURL(u, "/list-nodes"), encodeHTTPGenericRequest, decodeListNodesResponse, options["ListNodes"]...).Endpoint()
	}

	var deleteNodeEndpoint endpoint.Endpoint
	{
		deleteNodeEndpoint = http.NewClient("POST", copyURL(u, "/delete-node"), encodeHTTPGenericRequest, decodeDeleteNodeResponse, options["DeleteNode"]...).Endpoint()
	}

	var updateNodeEndpoint endpoint.Endpoint
	{
		updateNodeEndpoint = http.NewClient("POST", copyURL(u, "/update-node"), encodeHTTPGenericRequest, decodeUpdateNodeResponse, options["UpdateNode"]...).Endpoint()
	}

	var addRegionEndpoint endpoint.Endpoint
	{
		addRegionEndpoint = http.NewClient("POST", copyURL(u, "/add-region"), encodeHTTPGenericRequest, decodeAddRegionResponse, options["AddRegion"]...).Endpoint()
	}

	var listRegionsEndpoint endpoint.Endpoint
	{
		listRegionsEndpoint = http.NewClient("POST", copyURL(u, "/list-regions"), encodeHTTPGenericRequest, decodeListRegionsResponse, options["ListRegions"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		AddNodeEndpoint:        addNodeEndpoint,
		AddRegionEndpoint:      addRegionEndpoint,
		ChangePasswordEndpoint: changePasswordEndpoint,
		DeleteNodeEndpoint:     deleteNodeEndpoint,
		GetNodeEndpoint:        getNodeEndpoint,
		ListNodesEndpoint:      listNodesEndpoint,
		ListRegionsEndpoint:    listRegionsEndpoint,
		ListUsersEndpoint:      listUsersEndpoint,
		LoginEndpoint:          loginEndpoint,
		RegisterEndpoint:       registerEndpoint,
		UpdateNodeEndpoint:     updateNodeEndpoint,
		UpdateUserEndpoint:     updateUserEndpoint,
		ViewUserEndpoint:       viewUserEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeAuthThingResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAuthThingResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AuthThingResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeRegisterResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeRegisterResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.RegisterResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeLoginResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeLoginResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.LoginResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeViewUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeViewUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ViewUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListUsersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListUsersResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListUsersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeChangePasswordResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeChangePasswordResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ChangePasswordResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListNodesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListNodesResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListNodesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddRegionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddRegionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddRegionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListRegionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListRegionsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListRegionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
