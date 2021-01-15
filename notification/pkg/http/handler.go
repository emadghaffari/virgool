package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"

	endpoint "github.com/emadghaffari/virgool/notification/pkg/endpoint"
)

// makeSMSHandler creates the handler logic
func makeSMSHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/sms", http1.NewServer(endpoints.SMSEndpoint, decodeSMSRequest, encodeSMSResponse, options...))
}

// decodeSMSRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSMSRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SMSRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSMSResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSMSResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSMSTHandler creates the handler logic
func makeSMSTHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/smst", http1.NewServer(endpoints.SMSTEndpoint, decodeSMSTRequest, encodeSMSTResponse, options...))
}

// decodeSMSTRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSMSTRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SMSTRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSMSTResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSMSTResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeEmailHandler creates the handler logic
func makeEmailHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/email", http1.NewServer(endpoints.EmailEndpoint, decodeEmailRequest, encodeEmailResponse, options...))
}

// decodeEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.EmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeVerifyHandler creates the handler logic
func makeVerifyHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/verify", http1.NewServer(endpoints.VerifyEndpoint, decodeVerifyRequest, encodeVerifyResponse, options...))
}

// decodeVerifyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.VerifyRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeVerifyResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeVerifyResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
