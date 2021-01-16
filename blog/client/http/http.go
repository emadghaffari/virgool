package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"

	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"

	endpoint1 "github.com/emadghaffari/virgool/blog/pkg/endpoint"
	http2 "github.com/emadghaffari/virgool/blog/pkg/http"
	service "github.com/emadghaffari/virgool/blog/pkg/service"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.BlogService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var createPostEndpoint endpoint.Endpoint
	{
		createPostEndpoint = http.NewClient("POST", copyURL(u, "/create-post"), encodeHTTPGenericRequest, decodeCreatePostResponse, options["CreatePost"]...).Endpoint()
	}

	var updatePostEndpoint endpoint.Endpoint
	{
		updatePostEndpoint = http.NewClient("POST", copyURL(u, "/update-post"), encodeHTTPGenericRequest, decodeUpdatePostResponse, options["UpdatePost"]...).Endpoint()
	}

	var getPostEndpoint endpoint.Endpoint
	{
		getPostEndpoint = http.NewClient("POST", copyURL(u, "/get-post"), encodeHTTPGenericRequest, decodeGetPostResponse, options["GetPost"]...).Endpoint()
	}

	var deletePostEndpoint endpoint.Endpoint
	{
		deletePostEndpoint = http.NewClient("POST", copyURL(u, "/delete-post"), encodeHTTPGenericRequest, decodeDeletePostResponse, options["DeletePost"]...).Endpoint()
	}

	var createTagEndpoint endpoint.Endpoint
	{
		createTagEndpoint = http.NewClient("POST", copyURL(u, "/create-tag"), encodeHTTPGenericRequest, decodeCreateTagResponse, options["CreateTag"]...).Endpoint()
	}

	var getTagEndpoint endpoint.Endpoint
	{
		getTagEndpoint = http.NewClient("POST", copyURL(u, "/get-tag"), encodeHTTPGenericRequest, decodeGetTagResponse, options["GetTag"]...).Endpoint()
	}

	var updateTagEndpoint endpoint.Endpoint
	{
		updateTagEndpoint = http.NewClient("POST", copyURL(u, "/update-tag"), encodeHTTPGenericRequest, decodeUpdateTagResponse, options["UpdateTag"]...).Endpoint()
	}

	var deleteTagEndpoint endpoint.Endpoint
	{
		deleteTagEndpoint = http.NewClient("POST", copyURL(u, "/delete-tag"), encodeHTTPGenericRequest, decodeDeleteTagResponse, options["DeleteTag"]...).Endpoint()
	}

	var uploadEndpoint endpoint.Endpoint
	{
		uploadEndpoint = http.NewClient("POST", copyURL(u, "/upload"), encodeHTTPGenericRequest, decodeUploadResponse, options["Upload"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreatePostEndpoint: createPostEndpoint,
		CreateTagEndpoint:  createTagEndpoint,
		DeletePostEndpoint: deletePostEndpoint,
		DeleteTagEndpoint:  deleteTagEndpoint,
		GetPostEndpoint:    getPostEndpoint,
		GetTagEndpoint:     getTagEndpoint,
		UpdatePostEndpoint: updatePostEndpoint,
		UpdateTagEndpoint:  updateTagEndpoint,
		UploadEndpoint:     uploadEndpoint,
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

// decodeCreatePostResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreatePostResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreatePostResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdatePostResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdatePostResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdatePostResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetPostResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetPostResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetPostResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeletePostResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeletePostResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeletePostResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateTagResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateTagResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateTagResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTagResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTagResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTagResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTagResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTagResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTagResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTagResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTagResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTagResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUploadResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUploadResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UploadResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
