package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/gojsonq"

	model "github.com/emadghaffari/virgool/club/model"
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

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return "user not found", err
	}

	// get user from context
	usr, ok := ctx.Value(model.User).(map[string]interface{})
	if !ok {
		return "user not found", fmt.Errorf(fmt.Sprintf("error user not found: %s", err.Error()))
	}

	if usr["id"] != id {
		jq := gojsonq.New().FromString(fmt.Sprintf("%s", user))
		if cnt := jq.Where("Permissions", "=", "admin").Count(); cnt == 0 {
			return "you have not permisstion", fmt.Errorf("permisstion error")
		}
	}

	return l.next.Get(context.WithValue(ctx, model.User, user), id, token)
}

func (l loggingMiddleware) Index(ctx context.Context, from int32, size int32, filter []*model.Query, token string) (results []model.Point, err error) {
	defer func() {
		l.logger.Log("method", "Index", "from", from, "size", size, "filter", filter, "token", token, "results", results, "err", err)
	}()

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return nil, err
	}

	return l.next.Index(ctx, from, size, filter, token)
}
