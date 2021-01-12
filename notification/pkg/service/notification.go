package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/notification/conf"
	"github.com/emadghaffari/virgool/notification/database/redis"
	"github.com/emadghaffari/virgool/notification/model/notif"
)

var (
	// Streamer var
	Streamer StreamNotificationService = &streamNotificationService{}
)

// StreamNotificationService interface
type StreamNotificationService interface {
	// Store(ctx context.Context, code int, data map[string]interface{}) (err error)
	Store(ctx context.Context, code int, item notif.Notification) (err error)
}

type streamNotificationService struct{}

// func (s *streamNotificationService) Store(ctx context.Context, code int, item map[string]interface{}) (err error) {
func (s *streamNotificationService) Store(ctx context.Context, code int, item notif.Notification) (err error) {
	bt, err := json.Marshal(item.Data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("cannot marshal data key from consumed data Error: %s ", err),
		}).Fatal(fmt.Sprintf("cannot marshal data key from consumed data Error: %s ", err))

	}

	if notifire, err := notif.GetNotifier(item.Type); err == nil {
		notifire.Send(ctx, item)
	}

	data := item.Data.(map[string]interface{})

	// count is code we sms to clients
	key := data["phone"].(string) + "_" + strconv.Itoa(code)
	redis.Database.Set(context.Background(), key, string(bt), conf.GlobalConfigs.Service.Redis.SMSDuration)

	return err
}
