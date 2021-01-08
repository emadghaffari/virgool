package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"

	model "github.com/emadghaffari/virgool/auth/model"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}
}

// FIXME add middlewares for service

func (l loggingMiddleware) Register(ctx context.Context, Username string, Password string, Name string, LastName string, Phone string, Email string) (Response model.User, err error) {
	defer func() {
		l.logger.Log("method", "Register", "Username", Username, "Password", Password, "Name", Name, "LastName", LastName, "Phone", Phone, "Email", Email, "Response", Response, "err", err)
	}()

	err = model.Validator.Get().Struct(model.User{Username: Username, Password: &Password, Name: Name, LastName: LastName, Phone: Phone, Email: Email})
	if err != nil {
		return model.User{}, fmt.Errorf("Error: %s", err.Error())
	}

	return l.next.Register(ctx, Username, Password, Name, LastName, Phone, Email)
}
func (l loggingMiddleware) LoginUP(ctx context.Context, Username string, Password string) (Response model.User, err error) {
	defer func() {
		l.logger.Log("method", "LoginUP", "Username", Username, "Password", Password, "Response", Response, "err", err)
	}()
	return l.next.LoginUP(ctx, Username, Password)
}
func (l loggingMiddleware) LoginP(ctx context.Context, Phone string) (Response model.User, err error) {
	defer func() {
		l.logger.Log("method", "LoginP", "Phone", Phone, "Response", Response, "err", err)
	}()
	return l.next.LoginP(ctx, Phone)
}
func (l loggingMiddleware) Verify(ctx context.Context, Token string, Type string, Code string) (Response model.User, err error) {
	defer func() {
		l.logger.Log("method", "Verify", "Token", Token, "Type", Type, "Code", Code, "Response", Response, "err", err)
	}()
	return l.next.Verify(ctx, Token, Type, Code)
}
