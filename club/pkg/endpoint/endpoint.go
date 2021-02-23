package endpoint

import (
	"context"

	model "github.com/emadghaffari/virgool/club/model"
	service "github.com/emadghaffari/virgool/club/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

type GetRequest struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type GetResponse struct {
	Result string `json:"result"`
	Err    error  `json:"err"`
}

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

func (r GetResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) Get(ctx context.Context, id string, token string) (result string, err error) {
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

type IndexRequest struct {
	From   int32       `json:"from"`
	Size   int32       `json:"size"`
	Filter model.Query `json:"filter"`
	Token  string      `json:"token"`
}

type IndexResponse struct {
	Results []model.Point `json:"results"`
	Err     error         `json:"err"`
}

func MakeIndexEndpoint(s service.ClubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IndexRequest)
		results, err := s.Index(ctx, req.From, req.Size, req.Filter, req.Token)
		return IndexResponse{
			Err:     err,
			Results: results,
		}, nil
	}
}

func (r IndexResponse) Failed() error {
	return r.Err
}

func (e Endpoints) Index(ctx context.Context, from int32, size int32, filter model.Query, token string) (results []model.Point, err error) {
	request := IndexRequest{
		Filter: filter,
		From:   from,
		Size:   size,
		Token:  token,
	}
	response, err := e.IndexEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(IndexResponse).Results, response.(IndexResponse).Err
}
