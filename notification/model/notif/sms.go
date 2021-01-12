package notif

import (
	"context"
	"fmt"
	"time"
)

// SMS struct
type SMS struct{}

// Send meth
func (s *SMS) Send(ctx context.Context, notif Notification) error {
	fmt.Println("START SEND SMS ....")
	time.Sleep(time.Second)
	fmt.Println("SMS SENDED")

	return nil
}
