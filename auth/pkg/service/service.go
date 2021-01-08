package service

import (
	"context"
	"fmt"
	"time"

	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/database/mysql"
	"github.com/emadghaffari/virgool/auth/database/redis"
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

	// token: 09355960597, Type: SMS_CODE, Code: 31649
	// token: A4613Ac9...., Type: JWT_TOKEN, Code: ....
	Verify(ctx context.Context, Token, Type, Code string) (Response model.User, err error)
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
		tx.Rollback()
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
		tx.Rollback()
		return Response, fmt.Errorf(err.Error())
	}

	// try to store phone fo 2 min
	// IF Have Error
	// TODO Change the way the code is sent. Send by email
	if err := redis.Database.Set(context.Background(), user.Phone, "NOTIFICATION", time.Duration(conf.GlobalConfigs.Service.Redis.SMSDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
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

	// find the user with username or email
	user := model.User{}
	if err := tx.Table("users").Where("username = ? OR email = ?", Username, Username).First(&user).Error; err != nil {
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

	// check phone number for sended code before {DB: Redis - Time: 2min}
	var dst string
	if err := redis.Database.Get(context.Background(), Phone, &dst); err == nil && dst == "NOTIFICATION" {
		return model.User{}, fmt.Errorf("We have sent you an SMS. Please check your number {%s}", Phone)
	}

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// find the user user
	user := model.User{}
	if err := tx.Table("users").Where("phone = ?", Phone).First(&user).Error; err != nil {
		tx.Rollback()
		return Response, fmt.Errorf(err.Error())
	}

	// try to store phone fo 2 min
	// IF Have Error
	// TODO Change the way the code is sent. Send by email
	if err := redis.Database.Set(context.Background(), user.Phone, "NOTIFICATION", time.Duration(conf.GlobalConfigs.Service.Redis.SMSDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	// TODO add notification service for send SMS to user for login
	// send user_id to notif service - in notif service generate code and send notif to client

	tx.Commit()

	return user, err
}
func (b *basicAuthService) Verify(ctx context.Context, Token string, Type string, Code string) (Response model.User, err error) {
	var dst string

	// check the token exists in redis or not
	// example:
	// {Token == Phone} and Phone exists in Redis then you can check SMS Status
	if err := redis.Database.Get(context.Background(), Token, &dst); err != nil {
		return model.User{}, fmt.Errorf("please check your identity")
	}

	// check for sended Notification before
	// if code exists do not send code to notif service
	if err := redis.Database.Get(context.Background(), Token+"_"+Type, &dst); err == nil && dst == "NOTIFICATION" {
		return model.User{}, fmt.Errorf("You have tried before, please wait a minute")
	}

	// if code not exists in {redis} then store into redis and send code to notif service
	// store code in redis for every (10sec) for each requset to notif service
	if err := redis.Database.Set(context.Background(), Token+"_"+Type, "NOTIFICATION", time.Duration(conf.GlobalConfigs.Service.Redis.SMSCodeVerification)); err != nil {
		return model.User{}, err
	}

	// TODO code{Token} to notif service
	// response for verify user something like [code: "---", status:"VERIFY | BANDED | ..."]
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
