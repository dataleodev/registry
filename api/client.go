package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func MakeClientEndpoints(instance string) (e Endpoints, err error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	var options []kithttp.ClientOption

	var authThingEndpoint endpoint.Endpoint
	{
		authThingEndpoint = kithttp.NewClient(
			http.MethodGet,
			tgt,
			encodeAuthThingRequest,
			decodeAuthThingResponse,
			options...).Endpoint()
	}

	var registerEndpoint endpoint.Endpoint

	var loginEndpoint endpoint.Endpoint

	var viewUserEndpoint endpoint.Endpoint

	var listUsersEndpoint endpoint.Endpoint

	var updateUserEndpoint endpoint.Endpoint

	var changePasswordEndpoint endpoint.Endpoint

	var addNodeEndpoint endpoint.Endpoint

	var getNodeEndpoint endpoint.Endpoint

	var listNodesEndpoint endpoint.Endpoint

	var deleteNodeEndpoint endpoint.Endpoint

	var updateNodeEndpoint endpoint.Endpoint

	var addRegionEndpoint endpoint.Endpoint

	var listRegionsEndpoint endpoint.Endpoint

	e, err = Endpoints{
		AuthThingEndpoint:      authThingEndpoint,
		RegisterEndpoint:       registerEndpoint,
		LoginEndpoint:          loginEndpoint,
		ViewUserEndpoint:       viewUserEndpoint,
		ListUsersEndpoint:      listUsersEndpoint,
		UpdateUserEndpoint:     updateUserEndpoint,
		ChangePasswordEndpoint: changePasswordEndpoint,
		AddNodeEndpoint:        addNodeEndpoint,
		GetNodeEndpoint:        getNodeEndpoint,
		ListNodesEndpoint:      listNodesEndpoint,
		DeleteNodeEndpoint:     deleteNodeEndpoint,
		UpdateNodeEndpoint:     updateNodeEndpoint,
		AddRegionEndpoint:      addRegionEndpoint,
		ListRegionsEndpoint:    listRegionsEndpoint,
	}, nil

	return
}

func encodeAuthThingRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/auth")
	req.URL.Path = "/auth"
	return encodeHTTPGenericRequest(ctx, req, request)
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
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
func decodeAuthThingResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AuthThingResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeRegisterResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeRegisterResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp RegisterResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeLoginResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeLoginResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp LoginResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeViewUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeViewUserResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ViewUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListUsersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListUsersResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ListUsersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp UpdateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeChangePasswordResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeChangePasswordResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ChangePasswordResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddNodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AddNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetNodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp GetNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListNodesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListNodesResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ListNodesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteNodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp DeleteNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateNodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp UpdateNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddRegionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddRegionResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AddRegionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListRegionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListRegionsResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ListRegionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
