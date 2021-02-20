package service

import "context"

// ClubService describes the service.
type ClubService interface {
	// Add your methods here
	Get(ctx context.Context, id string, token string) (result int, err error)
}

type basicClubService struct{}

func (b *basicClubService) Get(ctx context.Context, id string, token string) (result int, err error) {
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
