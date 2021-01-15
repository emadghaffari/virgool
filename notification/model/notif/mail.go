package notif

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"

	"github.com/emadghaffari/virgool/notification/conf"
	"github.com/emadghaffari/virgool/notification/model"
)

type email struct {
	isSuccessful bool
	templateID   string
	template     string
	from         string
	to           string
	subject      string
	body         string
}

func (e *email) SendWithTemplate(ctx context.Context, notif Notification, params []Params, templateID string) error {
	if err := e.validate(notif); err != nil {
		return err
	}

	return nil
}
func (e *email) SendWithBody(ctx context.Context, notif Notification, options ...Option) error {
	if err := e.validate(notif); err != nil {
		return err
	}
	td := struct {
		Subject string
		Body    string
	}{
		Body:    e.body,
		Subject: e.subject,
	}
	if err := e.parseTemplate(conf.GlobalConfigs.Notif.Email.Send.Template, td); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", e.to)
	m.SetHeader("Subject", e.subject)
	m.SetBody("text/html", e.template)
	d := gomail.NewDialer(conf.GlobalConfigs.Notif.Email.Host,
		conf.GlobalConfigs.Notif.Email.Port,
		conf.GlobalConfigs.Notif.Email.Username,
		conf.GlobalConfigs.Notif.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Send Email with smtp server: %v", err),
		}).Warn(fmt.Sprintf("Mail - Error in Send Email with smtp server: %v", err))
		return err
	}

	return nil
}

func (e *email) validate(notif Notification) error {
	// Validate from exists in notif
	from, ok := notif.Data.(map[string]interface{})["from"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Get {from} mail From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Mail - Error in Get {from} mail From Notif Data: %v", notif))
		return fmt.Errorf("Mail - {from} not exists")
	}
	if err := model.Validator.Get().Var(from, "required,email"); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Invalid From Vaule: %v", notif),
		}).Warn(fmt.Sprintf("Mail - Error in Invalid From Vaule: %v", notif))
		return fmt.Errorf("Mail - Invalid From Vaule")
	}
	e.from = from

	// Validate to exists in notif
	to, ok := notif.Data.(map[string]interface{})["to"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Get {to} mail From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Mail - Error in Get {to} mail From Notif Data: %v", notif))
		return fmt.Errorf("Mail - {to} not exists")
	}
	if err := model.Validator.Get().Var(to, "required,email"); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Invalid to Vaule: %v", notif),
		}).Warn(fmt.Sprintf("Mail - Error in Invalid to Vaule: %v", notif))
		return fmt.Errorf("Mail - Invalid to Vaule")
	}
	e.to = to

	// Validate subject exists in notif
	subject, ok := notif.Data.(map[string]interface{})["subject"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Get {subject} mail From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Mail - Error in Get {subject} mail From Notif Data: %v", notif))
		return fmt.Errorf("Mail - {subject} not exists")
	}
	e.subject = subject

	// Validate body exists in notif
	body, ok := notif.Data.(map[string]interface{})["body"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Get {body} mail From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Mail - Error in Get {body} mail From Notif Data: %v", notif))
		return fmt.Errorf("Mail - {body} not exists")
	}
	e.body = body

	return nil
}

func (e *email) parseTemplate(tfn string, data interface{}) error {
	if !conf.FileExists(tfn) {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - template file not found: %s", tfn),
		}).Warn(fmt.Sprintf("Mail - template file not found: %s", tfn))
		return fmt.Errorf("Mail - template file not found: %s", tfn)
	}

	t, err := template.ParseFiles(tfn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in ParseFiles from template: %s", tfn),
		}).Warn(fmt.Sprintf("Mail - Error in ParseFiles from template: %s", tfn))
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Mail - Error in Execute parsed file from template: %s", tfn),
		}).Warn(fmt.Sprintf("Mail - Error in Execute parsed file from template: %s", tfn))
		return err
	}
	e.template = buf.String()

	return nil
}
