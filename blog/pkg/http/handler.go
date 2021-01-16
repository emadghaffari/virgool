package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/emadghaffari/virgool/blog/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeCreatePostHandler creates the handler logic
func makeCreatePostHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-post", http1.NewServer(endpoints.CreatePostEndpoint, decodeCreatePostRequest, encodeCreatePostResponse, options...))
}

// decodeCreatePostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreatePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreatePostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreatePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreatePostResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdatePostHandler creates the handler logic
func makeUpdatePostHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-post", http1.NewServer(endpoints.UpdatePostEndpoint, decodeUpdatePostRequest, encodeUpdatePostResponse, options...))
}

// decodeUpdatePostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdatePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdatePostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdatePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdatePostResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetPostHandler creates the handler logic
func makeGetPostHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-post", http1.NewServer(endpoints.GetPostEndpoint, decodeGetPostRequest, encodeGetPostResponse, options...))
}

// decodeGetPostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetPostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetPostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetPostResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeletePostHandler creates the handler logic
func makeDeletePostHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-post", http1.NewServer(endpoints.DeletePostEndpoint, decodeDeletePostRequest, encodeDeletePostResponse, options...))
}

// decodeDeletePostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeletePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeletePostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeletePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeletePostResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateTagHandler creates the handler logic
func makeCreateTagHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-tag", http1.NewServer(endpoints.CreateTagEndpoint, decodeCreateTagRequest, encodeCreateTagResponse, options...))
}

// decodeCreateTagRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateTagRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateTagResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateTagResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTagHandler creates the handler logic
func makeGetTagHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-tag", http1.NewServer(endpoints.GetTagEndpoint, decodeGetTagRequest, encodeGetTagResponse, options...))
}

// decodeGetTagRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTagRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTagResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTagResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTagHandler creates the handler logic
func makeUpdateTagHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-tag", http1.NewServer(endpoints.UpdateTagEndpoint, decodeUpdateTagRequest, encodeUpdateTagResponse, options...))
}

// decodeUpdateTagRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTagRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTagResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTagResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTagHandler creates the handler logic
func makeDeleteTagHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-tag", http1.NewServer(endpoints.DeleteTagEndpoint, decodeDeleteTagRequest, encodeDeleteTagResponse, options...))
}

// decodeDeleteTagRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTagRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTagResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTagResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUploadHandler creates the handler logic
func makeUploadHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/upload", http1.NewServer(endpoints.UploadEndpoint, decodeUploadRequest, encodeUploadResponse, options...))
}

// decodeUploadRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUploadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UploadRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUploadResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUploadResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
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
