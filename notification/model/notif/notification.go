package notif

import (
	"context"
	"fmt"
)

type notification string

const (
	// EMAIL for get Email object
	EMAIL notification = "EMAIL"

	// SMS for get sms object
	SMS notification = "SMS"
)

// Notifier interface
type Notifier interface {
	SendWithTemplate(ctx context.Context, notif Notification, params []Params, template string) error
	SendWithBody(ctx context.Context, notif Notification, options ...Option) error
}

// Notification struct
type Notification struct {
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	KEY          string      `json:"key"`
	Type         string      `json:"type"`
	sendDateTime string      `default:""`
	line         string      `default:""`
}

// Params struct
type Params struct {
	Parameter      string
	ParameterValue interface{}
}

// Option func for SMS
type Option func(*Notification)

// GetNotifier func
func GetNotifier(t notification) (Notifier, error) {
	switch t {
	case "SMS":
		return new(sms), nil
	case "EMAIL":
		return new(email), nil
	default:
		return nil, fmt.Errorf("Not Found")
	}
}

// GetType for types we use when need to send a notif
func GetType(t string) notification {
	switch t {
	case "SMS":
		return SMS
	case "EMAIL":
		return EMAIL
	}
	return ""
}

// SendDateTime Option func
// you can specifically send a notif in specific time
func SendDateTime(time string) Option {
	return func(n *Notification) {
		n.sendDateTime = time
	}
}

// Line Option func
// you can specifically send a notif from a line
// examples: PHONE_NUMBER, EMAIL_ADDRESS,....
func Line(line string) Option {
	return func(n *Notification) {
		n.line = line
	}
}
