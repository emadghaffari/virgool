package notif

import (
	"context"
	"fmt"
)

// Notifier interface
type Notifier interface {
	SendWithTemplate(ctx context.Context, notif Notification, params []SMSParams, template string) error
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

// Option func for SMS
type Option func(*Notification)

// GetNotifier func
func GetNotifier(t string) (Notifier, error) {
	switch t {
	case "SMS":
		return new(SMS), nil
	default:
		return nil, fmt.Errorf("Not Found")
	}
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
