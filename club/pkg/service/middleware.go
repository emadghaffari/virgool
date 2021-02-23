package service

import (
	"context"

	model "github.com/emadghaffari/virgool/club/model"
	log "github.com/go-kit/kit/log"
)

type Middleware func(ClubService) ClubService

type loggingMiddleware struct {
	logger log.Logger
	next   ClubService
}

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

func (l loggingMiddleware) Index(ctx context.Context, from int32, size int32, filter model.Query, token string) (results []model.Point, err error) {
	defer func() {
		l.logger.Log("method", "Index", "from", from, "size", size, "filter", filter, "token", token, "results", results, "err", err)
	}()
	return l.next.Index(ctx, from, size, filter, token)
}
