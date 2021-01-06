package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint1 "github.com/emadghaffari/virgool/auth/pkg/endpoint"
	http2 "github.com/emadghaffari/virgool/auth/pkg/http"
	service "github.com/emadghaffari/virgool/auth/pkg/service"
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
func New(instance string, options map[string][]http.ClientOption) (service.AuthService, error) {
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

	var loginUPEndpoint endpoint.Endpoint
	{
		loginUPEndpoint = http.NewClient("POST", copyURL(u, "/login-up"), encodeHTTPGenericRequest, decodeLoginUPResponse, options["LoginUP"]...).Endpoint()
	}

	var loginPEndpoint endpoint.Endpoint
	{
		loginPEndpoint = http.NewClient("POST", copyURL(u, "/login-p"), encodeHTTPGenericRequest, decodeLoginPResponse, options["LoginP"]...).Endpoint()
	}

	var verifyEndpoint endpoint.Endpoint
	{
		verifyEndpoint = http.NewClient("POST", copyURL(u, "/verify"), encodeHTTPGenericRequest, decodeVerifyResponse, options["Verify"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		LoginPEndpoint:   loginPEndpoint,
		LoginUPEndpoint:  loginUPEndpoint,
		RegisterEndpoint: registerEndpoint,
		VerifyEndpoint:   verifyEndpoint,
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

// decodeLoginUPResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeLoginUPResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.LoginUPResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeLoginPResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeLoginPResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.LoginPResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeVerifyResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeVerifyResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.VerifyResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
