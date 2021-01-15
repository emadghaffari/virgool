package service

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"github.com/emadghaffari/virgool/notification/conf"
	"github.com/emadghaffari/virgool/notification/database/redis"
	"github.com/emadghaffari/virgool/notification/model/notif"
)

// NotificationService describes the service.
type NotificationService interface {
	// Add your methods here
	SMS(ctx context.Context, to, body string) (message, status string, err error)
	SMST(ctx context.Context, to string, params map[string]string, template string) (message, status string, err error)
	Email(ctx context.Context, to, body string, data interface{}) (message, status string, err error)
	Verify(ctx context.Context, phone, code string) (message, status string, data interface{}, err error)
}

type basicNotificationService struct{}

// send sms with body
func (b *basicNotificationService) SMS(ctx context.Context, to string, body string) (message string, status string, err error) {
	var tracer opentracing.Tracer
	var span opentracing.Span

	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer = opentracing.GlobalTracer(); tracer != nil {
			span = tracer.StartSpan("SendWithBody", opentracing.ChildOf(pctx))
			defer span.Finish()
		}
	}

	objs := make(map[string]interface{}, 2)
	objs["phone"] = to
	objs["body"] = body

	ntf := notif.Notification{
		Data: objs,
		Type: "SMS",
	}

	if notifire, err := notif.GetNotifier(notif.SMS); err == nil {
		err := notifire.SendWithBody(ctx, ntf, notif.Line(conf.GlobalConfigs.Notif.SMS.Send.LineNumber[0]))
		if err != nil {
			return "ERROR IN SEND SMS", "FAILED", err
		}
	}

	return "SUCCESS", "OK", err
}

func (b *basicNotificationService) SMST(ctx context.Context, to string, params map[string]string, template string) (message string, status string, err error) {
	var tracer opentracing.Tracer
	var span opentracing.Span

	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer = opentracing.GlobalTracer(); tracer != nil {
			span = tracer.StartSpan("SendWithTemplate", opentracing.ChildOf(pctx))
			defer span.Finish()
		}
	}

	dt := make(map[string]interface{}, 1)
	dt["phone"] = to

	ntf := notif.Notification{
		Data: dt,
		Type: "SMS",
	}
	ps := make([]notif.Params, len(params))
	cc := 0
	for k, v := range params {
		ps[cc] = notif.Params{Parameter: k, ParameterValue: v}
		cc++
	}

	if notifire, err := notif.GetNotifier(notif.SMS); err == nil {
		err := notifire.SendWithTemplate(ctx, ntf, ps, template)
		if err != nil {
			return "ERROR IN SEND SMS", "FAILED", err
		}
	}
	return message, status, err
}

// FIXME fix Email
func (b *basicNotificationService) Email(ctx context.Context, to string, body string, data interface{}) (message string, status string, err error) {
	var tracer opentracing.Tracer
	var span opentracing.Span

	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer = opentracing.GlobalTracer(); tracer != nil {
			span = tracer.StartSpan("Email", opentracing.ChildOf(pctx))
			defer span.Finish()
		}
	}

	objs := make(map[string]interface{}, 2)
	objs["to"] = to
	objs["body"] = body
	objs["from"] = "mail@gmail.com"
	objs["subject"] = "subjct"

	fmt.Println("//////////")
	fmt.Println(objs)
	fmt.Println("//////////")

	ntf := notif.Notification{
		Data: objs,
		Type: "SMS",
	}

	if notifire, err := notif.GetNotifier(notif.EMAIL); err == nil {
		err := notifire.SendWithBody(ctx, ntf, notif.Line(conf.GlobalConfigs.Notif.SMS.Send.LineNumber[0]))
		if err != nil {
			return "ERROR IN SEND SMS", "FAILED", err
		}
	}

	return "SUCCESS", "OK", err
}
func (b *basicNotificationService) Verify(ctx context.Context, phone string, code string) (message string, status string, data interface{}, err error) {
	var tracer opentracing.Tracer
	var span opentracing.Span

	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer = opentracing.GlobalTracer(); tracer != nil {
			span = tracer.StartSpan("verify", opentracing.ChildOf(pctx))
			defer span.Finish()
		}
	}

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
