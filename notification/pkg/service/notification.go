package service

import (
	"context"
	"strconv"

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
	data := item.Data.(map[string]interface{})

	if notifire, err := notif.GetNotifier(item.Type); err == nil {
		err := notifire.SendWithTemplate(ctx, item, []notif.SMSParams{
			{Parameter: "Code", ParameterValue: code},
		}, conf.GlobalConfigs.Notif.SMS.Send.Verify.TemplateID)
		if err != nil {
			return err
		}
	}

	// count is code we sms to clients
	key := data["phone"].(string) + "_" + strconv.Itoa(code)
	redis.Database.Set(context.Background(), key, item, conf.GlobalConfigs.Service.Redis.SMSDuration)

	return err
}
