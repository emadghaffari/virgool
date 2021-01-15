package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	service "github.com/emadghaffari/virgool/notification/pkg/service"
)

// SMSRequest collects the request parameters for the SMS method.
type SMSRequest struct {
	To   string      `json:"to"`
	Body string      `json:"body"`
	Data interface{} `json:"data"`
}

// SMSResponse collects the response parameters for the SMS method.
type SMSResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeSMSEndpoint returns an endpoint that invokes SMS on the service.
func MakeSMSEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SMSRequest)
		message, status, err := s.SMS(ctx, req.To, req.Body, req.Data)
		return SMSResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r SMSResponse) Failed() error {
	return r.Err
}

// EmailRequest collects the request parameters for the Email method.
type EmailRequest struct {
	To   string      `json:"to"`
	Body string      `json:"body"`
	Data interface{} `json:"data"`
}

// EmailResponse collects the response parameters for the Email method.
type EmailResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeEmailEndpoint returns an endpoint that invokes Email on the service.
func MakeEmailEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EmailRequest)
		message, status, err := s.Email(ctx, req.To, req.Body, req.Data)
		return EmailResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r EmailResponse) Failed() error {
	return r.Err
}

// VerifyRequest collects the request parameters for the Verify method.
type VerifyRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// VerifyResponse collects the response parameters for the Verify method.
type VerifyResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Err     error       `json:"err"`
}

// MakeVerifyEndpoint returns an endpoint that invokes Verify on the service.
func MakeVerifyEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyRequest)
		message, status, data, err := s.Verify(ctx, req.Phone, req.Code)
		return VerifyResponse{
			Data:    data,
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r VerifyResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SMS implements Service. Primarily useful in a client.
func (e Endpoints) SMS(ctx context.Context, to string, body string, data interface{}) (message string, status string, err error) {
	request := SMSRequest{
		Body: body,
		Data: data,
		To:   to,
	}
	response, err := e.SMSEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SMSResponse).Message, response.(SMSResponse).Status, response.(SMSResponse).Err
}

// Email implements Service. Primarily useful in a client.
func (e Endpoints) Email(ctx context.Context, to string, body string, data interface{}) (message string, status string, err error) {
	request := EmailRequest{
		Body: body,
		Data: data,
		To:   to,
	}
	response, err := e.EmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(EmailResponse).Message, response.(EmailResponse).Status, response.(EmailResponse).Err
}

// Verify implements Service. Primarily useful in a client.
func (e Endpoints) Verify(ctx context.Context, phone string, code string) (message string, status string, data interface{}, err error) {
	request := VerifyRequest{
		Code:  code,
		Phone: phone,
	}
	response, err := e.VerifyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(VerifyResponse).Message, response.(VerifyResponse).Status, response.(VerifyResponse).Data, response.(VerifyResponse).Err
}

// SMSTRequest collects the request parameters for the SMST method.
type SMSTRequest struct {
	To       string            `json:"to"`
	Params   map[string]string `json:"params"`
	Template string            `json:"template"`
	Time     string            `json:"time"`
	Data     interface{}       `json:"data"`
}

// SMSTResponse collects the response parameters for the SMST method.
type SMSTResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeSMSTEndpoint returns an endpoint that invokes SMST on the service.
func MakeSMSTEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SMSTRequest)
		message, status, err := s.SMST(ctx, req.To, req.Params, req.Template, req.Time, req.Data)
		return SMSTResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r SMSTResponse) Failed() error {
	return r.Err
}

// SMST implements Service. Primarily useful in a client.
func (e Endpoints) SMST(ctx context.Context, to string, params map[string]string, template string, time string, data interface{}) (message string, status string, err error) {
	request := SMSTRequest{
		Data:     data,
		Params:   params,
		Template: template,
		Time:     time,
		To:       to,
	}
	response, err := e.SMSTEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SMSTResponse).Message, response.(SMSTResponse).Status, response.(SMSTResponse).Err
}
