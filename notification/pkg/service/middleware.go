package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"

	"github.com/emadghaffari/virgool/notification/model"
)

// Middleware describes a service middleware.
type Middleware func(NotificationService) NotificationService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificationService) NotificationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SMS(ctx context.Context, to string, body string) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "SMS", "to", to, "body", body, "message", message, "status", status, "err", err)
	}()
	return l.next.SMS(ctx, to, body)
}

func (l loggingMiddleware) SMST(ctx context.Context, to string, params map[string]string, template string) (message, status string, err error) {
	defer func() {
		l.logger.Log("method", "SMS", "to", to, "params", params, "template", template, "message", message, "status", status, "err", err)
	}()
	return l.next.SMST(ctx, to, params, template)
}

func (l loggingMiddleware) Email(ctx context.Context, to string, body string, data interface{}) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "Email", "to", to, "body", body, "data", data, "message", message, "status", status, "err", err)
	}()
	return l.next.Email(ctx, to, body, data)
}
func (l loggingMiddleware) Verify(ctx context.Context, phone string, code string) (message string, status string, data interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Verify", "phone", phone, "code", code, "message", message, "status", status, "data", data, "err", err)
	}()

	if err := model.Validator.Get().Var(phone, "required"); err != nil {
		return fmt.Sprintf("Error: %s", err.Error()), "Error", "", fmt.Errorf("Error: %s", err.Error())
	}

	if err := model.Validator.Get().Var(code, "required"); err != nil {
		return fmt.Sprintf("Error: %s", err.Error()), "Error", "", fmt.Errorf("Error: %s", err.Error())
	}

	return l.next.Verify(ctx, phone, code)
}
