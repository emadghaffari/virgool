package service

import (
	"context"
	"fmt"

	"github.com/emadghaffari/virgool/auth/database/mysql"
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

	// Hash Password
	password, err := new(model.Bcrypt).HashPassword(Password)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// find the user role
	role := model.Role{}
	if err := tx.Table("roles").Where("name = ?", "user").First(&role).Error; err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// fill the user model
	user := model.User{
		Username: Username,
		Password: &password,
		Name:     Name,
		LastName: LastName,
		Phone:    Phone,
		Email:    Email,
		RoleID:   role.ID,
	}

	// try to store user with model
	if gm := tx.Create(&user); gm.Error != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// TODO add notification service for send SMS to user for register
	// send user_id to notif service - in notif service generate a code and send notif to client

	// commit a transaction
	tx.Commit()

	return user, err
}
func (b *basicAuthService) LoginUP(ctx context.Context, Username string, Password string) (Response model.User, err error) {

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// find the user user
	user := model.User{}
	if err := tx.Table("users").Where("username = ?", Username).First(&user).Error; err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// Check Hash Password
	if ok := new(model.Bcrypt).CheckPasswordHash(Password, *user.Password); !ok {
		return Response, fmt.Errorf("username or password not found")
	}

	// commit the transaction
	tx.Commit()

	return user, err
}
func (b *basicAuthService) LoginP(ctx context.Context, Phone string) (Response model.User, err error) {

	// TODO check phone number for sended code before {DB: Redis - Time: 2min}

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// find the user user
	user := model.User{}
	if err := tx.Table("users").Where("phone = ?", Phone).First(&user).Error; err != nil {
		tx.Rollback()
		return Response, fmt.Errorf(err.Error())
	}

	// TODO add notification service for send SMS to user for login
	// send user_id to notif service - in notif service generate code and send notif to client

	tx.Commit()

	return Response, err
}
func (b *basicAuthService) Verify(ctx context.Context, Token string, Type string, Device string) (Response model.User, err error) {

	// TODO store code in redis for every (10sec) for each requset to notif service
	// if code exists do not send code to notif service
	// if code not exists in {redis} then store into redis and send code to notif service

	// TODO send code{Token} to notif service
	// response for verify user something like [user_id:"1",code: "---", status:"VERIFY | BANDED | ..."]
	// then we can say user is verified or not!

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
