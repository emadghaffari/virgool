package notif

import (
	"context"
	"fmt"
)

// Notifier interface
type Notifier interface {
	Send(ctx context.Context, notif Notification, code int) error
}

// Notification struct
type Notification struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	KEY     string      `json:"key"`
	Type    string      `json:"type"`
}

// GetNotifier func
func GetNotifier(t string) (Notifier, error) {
	switch t {
	case "SMS":
		return new(SMS), nil
	default:
		return nil, fmt.Errorf("Not Found")
	}
}
