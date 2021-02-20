package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ClubService) ClubService

type loggingMiddleware struct {
	logger log.Logger
	next   ClubService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ClubService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ClubService) ClubService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context, id string, token string) (result string, err error) {
	defer func() {
		l.logger.Log("method", "Get", "id", id, "token", token, "result", result, "err", err)
	}()
	return l.next.Get(ctx, id, token)
}
