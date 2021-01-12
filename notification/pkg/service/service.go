package service

import (
	"context"

	"github.com/emadghaffari/virgool/notification/database/redis"
)

// NotificationService describes the service.
type NotificationService interface {
	// Add your methods here
	SMS(ctx context.Context, to, body string, data interface{}) (message, status string, err error)
	Email(ctx context.Context, to, body string, data interface{}) (message, status string, err error)
	Verify(ctx context.Context, phone, code string) (message, status string, data interface{}, err error)
}

type basicNotificationService struct{}

// FIXME fix SMS
func (b *basicNotificationService) SMS(ctx context.Context, to string, body string, data interface{}) (message string, status string, err error) {
	// TODO implement the business logic of SMS
	return message, status, err
}

// FIXME fix Email
func (b *basicNotificationService) Email(ctx context.Context, to string, body string, data interface{}) (message string, status string, err error) {
	// TODO implement the business logic of Email
	return message, status, err
}
func (b *basicNotificationService) Verify(ctx context.Context, phone string, code string) (message string, status string, data interface{}, err error) {

	if err := redis.Database.Get(context.Background(), phone+"_"+code, &data); err != nil {
		return "FAILD", "ERROR", data, err
	}

	redis.Database.Del(context.Background(), phone+"_"+code)

	message = "OK"
	status = "OK"

	return message, status, data, err
}

// NewBasicNotificationService returns a naive, stateless implementation of NotificationService.
func NewBasicNotificationService() NotificationService {
	return &basicNotificationService{}
}

// New returns a NotificationService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificationService {
	var svc NotificationService = NewBasicNotificationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
