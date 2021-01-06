package endpoint

import (
	"context"

	model "github.com/emadghaffari/virgool/auth/model"
	service "github.com/emadghaffari/virgool/auth/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	Response model.User `json:"response"`
	Err      error      `json:"err"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		Response, err := s.Register(ctx, req.Username, req.Password, req.Name, req.LastName, req.Phone, req.Email)
		return RegisterResponse{
			Err:      err,
			Response: Response,
		}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.Err
}

// LoginUPRequest collects the request parameters for the LoginUP method.
type LoginUPRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUPResponse collects the response parameters for the LoginUP method.
type LoginUPResponse struct {
	Response model.User `json:"response"`
	Err      error      `json:"err"`
}

// MakeLoginUPEndpoint returns an endpoint that invokes LoginUP on the service.
func MakeLoginUPEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginUPRequest)
		Response, err := s.LoginUP(ctx, req.Username, req.Password)
		return LoginUPResponse{
			Err:      err,
			Response: Response,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginUPResponse) Failed() error {
	return r.Err
}

// LoginPRequest collects the request parameters for the LoginP method.
type LoginPRequest struct {
	Phone string `json:"phone"`
}

// LoginPResponse collects the response parameters for the LoginP method.
type LoginPResponse struct {
	Response model.User `json:"response"`
	Err      error      `json:"err"`
}

// MakeLoginPEndpoint returns an endpoint that invokes LoginP on the service.
func MakeLoginPEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginPRequest)
		Response, err := s.LoginP(ctx, req.Phone)
		return LoginPResponse{
			Err:      err,
			Response: Response,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginPResponse) Failed() error {
	return r.Err
}

// VerifyRequest collects the request parameters for the Verify method.
type VerifyRequest struct {
	Token  string `json:"token"`
	Type   string `json:"type"`
	Device string `json:"device"`
}

// VerifyResponse collects the response parameters for the Verify method.
type VerifyResponse struct {
	Response model.User `json:"response"`
	Err      error      `json:"err"`
}

// MakeVerifyEndpoint returns an endpoint that invokes Verify on the service.
func MakeVerifyEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyRequest)
		Response, err := s.Verify(ctx, req.Token, req.Type, req.Device)
		return VerifyResponse{
			Err:      err,
			Response: Response,
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

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, Username string, Password string, Name string, LastName string, Phone string, Email string) (Response model.User, err error) {
	request := RegisterRequest{
		Email:    Email,
		LastName: LastName,
		Name:     Name,
		Password: Password,
		Phone:    Phone,
		Username: Username,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).Response, response.(RegisterResponse).Err
}

// LoginUP implements Service. Primarily useful in a client.
func (e Endpoints) LoginUP(ctx context.Context, Username string, Password string) (Response model.User, err error) {
	request := LoginUPRequest{
		Password: Password,
		Username: Username,
	}
	response, err := e.LoginUPEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginUPResponse).Response, response.(LoginUPResponse).Err
}

// LoginP implements Service. Primarily useful in a client.
func (e Endpoints) LoginP(ctx context.Context, Phone string) (Response model.User, err error) {
	request := LoginPRequest{Phone: Phone}
	response, err := e.LoginPEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginPResponse).Response, response.(LoginPResponse).Err
}

// Verify implements Service. Primarily useful in a client.
func (e Endpoints) Verify(ctx context.Context, Token string, Type string, Device string) (Response model.User, err error) {
	request := VerifyRequest{
		Device: Device,
		Token:  Token,
		Type:   Type,
	}
	response, err := e.VerifyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(VerifyResponse).Response, response.(VerifyResponse).Err
}
