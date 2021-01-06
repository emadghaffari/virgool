package service

import (
	"context"

	"github.com/emadghaffari/virgool/auth/model"
)

// AuthService describes the service.
type AuthService interface {
	// Add your methods here
	Register(ctx context.Context, Username, Password, Name, LastName, Phone, Email string) (Response model.User, err error)

	// Username: Username, password password
	// Username: Phone, password password
	LoginUP(ctx context.Context, Username, Password string) (Response model.User, err error)

	// Phone: phone
	LoginP(ctx context.Context, Phone string) (Response model.User, err error)

	// token: 13649, Type: SMS_CODE, Device: Phone
	// token: A4613Ac9...., Type: JWT_TOKEN, Device: MACBOOK
	Verify(ctx context.Context, Token, Type, Device string) (Response model.User, err error)
}

type basicAuthService struct{}

func (b *basicAuthService) Register(ctx context.Context, Username string, Password string, Name string, LastName string, Phone string, Email string) (Response model.User, err error) {
	// TODO implement the business logic of Register
	return Response, err
}
func (b *basicAuthService) LoginUP(ctx context.Context, Username string, Password string) (Response model.User, err error) {
	// TODO implement the business logic of LoginUP
	return Response, err
}
func (b *basicAuthService) LoginP(ctx context.Context, Phone string) (Response model.User, err error) {
	// TODO implement the business logic of LoginP
	return Response, err
}
func (b *basicAuthService) Verify(ctx context.Context, Token string, Type string, Device string) (Response model.User, err error) {
	// TODO implement the business logic of Verify
	return Response, err
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
