package service

import (
	"context"
	"fmt"
	"time"

	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/database/kafka"
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

	// token: A4613Ac9...., Type: JWT_TOKEN, Code: ....
	Verify(ctx context.Context, Token, Type, Code string) (Response model.User, err error)
}

type basicAuthService struct{}

// Register meth
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
	if err := tx.Table("roles").Preload("Permissions").Where("name = ?", "user").First(&role).Error; err != nil {
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
		Role:     role,
	}

	// try to store user with model
	if gm := tx.Create(&user); gm.Error != nil {
		tx.Rollback()
		return Response, fmt.Errorf(err.Error())
	}

	// try to store phone for 2 min
	// IF Have Error
	if err := redis.Database.Set(context.Background(), user.Phone, "NOTIFICATION", time.Duration(conf.GlobalConfigs.Service.Redis.SMSDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	// generate new jwt token
	jwt, err := model.JWT.Generate(context.Background(), user)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// send user_id to notif service - in notif service generate a code and send notif to client
	err = kafka.Database.Producer(model.Notification{
		Data: model.Data{
			User:  user,
			Token: jwt.AccessToken,
		},
		Message: "SMS",

		KEY: user.Phone,
	}, conf.GlobalConfigs.Kafka.Topics.Notif)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// Store User into redis DB with uuid
	if err := redis.Database.Set(context.Background(), jwt.AccessUUID, user, time.Duration(conf.GlobalConfigs.Service.Redis.UserDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	// commit a transaction
	tx.Commit()

	return user, err
}

// Login with username and password
func (b *basicAuthService) LoginUP(ctx context.Context, Username string, Password string) (Response model.User, err error) {

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// find the user with username or email
	user := model.User{}
	if err := tx.Table("users").Preload("Role").Preload("Role.Permissions").Where("username = ? OR email = ?", Username, Username).First(&user).Error; err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	fmt.Println(user)

	// Check Hash Password
	if ok := new(model.Bcrypt).CheckPasswordHash(Password, *user.Password); !ok {
		return Response, fmt.Errorf("username or password not found")
	}

	// generate new jwt token
	jwt, err := model.JWT.Generate(context.Background(), user)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// Store User into redis DB with uuid
	if err := redis.Database.Set(context.Background(), jwt.AccessUUID, user, time.Duration(conf.GlobalConfigs.Service.Redis.UserDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	// commit the transaction
	tx.Commit()

	return user, err
}

// login with phone and sms
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
	if err := tx.Table("users").Preload("Role").Preload("Role.Permissions").Where("phone = ?", Phone).First(&user).Error; err != nil {
		tx.Rollback()
		return Response, fmt.Errorf(err.Error())
	}

	// try to store phone fo 2 min
	// IF Have Error
	if err := redis.Database.Set(context.Background(), user.Phone, "NOTIFICATION", time.Duration(conf.GlobalConfigs.Service.Redis.SMSDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	// generate new jwt token
	jwt, err := model.JWT.Generate(context.Background(), user)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// send user_id to notif service - in notif service generate a code and send notif to client
	err = kafka.Database.Producer(model.Notification{
		Data: model.Data{
			User:  user,
			Token: jwt.AccessToken,
		},
		Message: "SMS",
		KEY:     user.Phone,
	}, conf.GlobalConfigs.Kafka.Topics.Notif)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// Store User into redis DB with uuid
	if err := redis.Database.Set(context.Background(), jwt.AccessUUID, user, time.Duration(conf.GlobalConfigs.Service.Redis.UserDuration)); err != nil {
		tx.Rollback()
		return model.User{}, err
	}
	tx.Commit()

	return user, err
}

// verify the jwt{UUID} and get user
func (b *basicAuthService) Verify(ctx context.Context, Token string, Type string, Code string) (Response model.User, err error) {

	// generate new jwt token
	tk, err := model.JWT.Verify(Token)
	if err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	if err := model.JWT.Get(context.Background(), tk, &Response); err != nil {
		return Response, fmt.Errorf(err.Error())
	}

	// return item.(model.User), err
	return Response, nil
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
