package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/emadghaffari/virgool/club/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeGetHandler creates the handler logic
func makeGetHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get", http1.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...))
}

// decodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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

// makeIndexHandler creates the handler logic
func makeIndexHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/index", http1.NewServer(endpoints.IndexEndpoint, decodeIndexRequest, encodeIndexResponse, options...))
}

// decodeIndexRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.IndexRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeIndexResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeIndexResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
