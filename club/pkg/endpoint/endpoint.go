package endpoint

import (
	"context"
	service "github.com/emadghaffari/virgool/club/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Result int   `json:"result"`
	Err    error `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.ClubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		result, err := s.Get(ctx, req.Id, req.Token)
		return GetResponse{
			Err:    err,
			Result: result,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, id string, token string) (result int, err error) {
	request := GetRequest{
		Id:    id,
		Token: token,
	}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Result, response.(GetResponse).Err
}
