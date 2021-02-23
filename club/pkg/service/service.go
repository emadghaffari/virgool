package service

import (
	"context"

	"github.com/emadghaffari/virgool/club/model"
)

// ClubService describes the service.
type ClubService interface {
	// Add your methods here
	Get(ctx context.Context, id string, token string) (result string, err error)
	Index(ctx context.Context, from, size int32, filter model.Query, token string) (results []model.Point, err error)
}

type basicClubService struct{}

func (b *basicClubService) Get(ctx context.Context, id string, token string) (result string, err error) {
	// TODO implement the business logic of Get
	return result, err
}

// NewBasicClubService returns a naive, stateless implementation of ClubService.
func NewBasicClubService() ClubService {
	return &basicClubService{}
}

// New returns a ClubService with all of the expected middleware wired in.
func New(middleware []Middleware) ClubService {
	var svc ClubService = NewBasicClubService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicClubService) Index(ctx context.Context, from int32, size int32, filter model.Query, token string) (results []model.Point, err error) {
	// TODO implement the business logic of Index
	return results, err
}
