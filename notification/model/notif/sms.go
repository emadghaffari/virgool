package notif

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/notification/conf"
)

// SMS struct
type SMS struct {
	TokenKey     string
	IsSuccessful bool
}

// Params struct
type Params struct {
	Parameter      string
	ParameterValue interface{}
}

// SendWithTemplate meth
// Need Phone in Data Params
func (s *SMS) SendWithTemplate(ctx context.Context, notif Notification, params []Params, template string) error {
	// Validate Phone number exists in notif
	mobile, ok := notif.Data.(map[string]interface{})["phone"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Get Phone Number From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Error in Get Phone Number From Notif Data: %v", notif))
		return fmt.Errorf("Phone not exists")
	}

	// Marshal Data We Need
	body, err := json.Marshal(map[string]interface{}{
		"ParameterArray": params,
		"Mobile":         mobile,
		"TemplateId":     template,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Marshal Send SMS Data: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Marshal Send SMS Data: %s", err.Error()))
		return err
	}

	headers := []struct {
		Header string
		Vaule  string
	}{
		{
			"Content-Type",
			conf.GlobalConfigs.Notif.SMS.Send.Verify.ContentType,
		},
	}

	if err := s.send("POST", conf.GlobalConfigs.Notif.SMS.Send.TemplateURL, headers, body); err != nil {
		return err
	}

	return nil
}

// SendWithBody meth
// Need Phone,Body in Data Params
func (s *SMS) SendWithBody(ctx context.Context, notif Notification, options ...Option) error {
	// Validate Phone number exists in notif
	mobile, ok := notif.Data.(map[string]interface{})["phone"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Get Phone Number From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Error in Get Phone Number From Notif Data: %v", notif))
		return fmt.Errorf("Phone not exists")
	}

	// FIXME change phone into body
	// Validate Phone number exists in notif
	txt, ok := notif.Data.(map[string]interface{})["phone"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Get Phone Number From Notif Data: %v", notif),
		}).Warn(fmt.Sprintf("Error in Get Phone Number From Notif Data: %v", notif))
		return fmt.Errorf("Phone not exists")
	}

	for _, option := range options {
		option(&notif)
	}

	// Marshal Data We Need
	body, err := json.Marshal(map[string]interface{}{
		"Messages":                 []string{txt},
		"MobileNumbers":            []string{mobile},
		"LineNumber":               notif.line,
		"SendDateTime":             notif.sendDateTime,
		"CanContinueInCaseOfError": "false",
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Marshal Send SMS Data: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Marshal Send SMS Data: %s", err.Error()))
		return err
	}

	headers := []struct {
		Header string
		Vaule  string
	}{
		{
			"Content-Type",
			conf.GlobalConfigs.Notif.SMS.Send.Verify.ContentType,
		},
	}
	if err := s.send("POST", conf.GlobalConfigs.Notif.SMS.Send.BodyURL, headers, body); err != nil {
		return err
	}

	return nil
}

// Get Token We Need for every SMS We Send from sms.ir
func (s *SMS) token() error {
	// try to Marshal Api key, Secret key
	body, err := json.Marshal(map[string]string{
		"UserApiKey": conf.GlobalConfigs.Notif.SMS.UserAPIKey,
		"SecretKey":  conf.GlobalConfigs.Notif.SMS.SecretKey,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Marshal Data sms.go: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Marshal Data sms.go: %s", err.Error()))
		return err
	}

	headers := []struct {
		Header string
		Vaule  string
	}{
		{
			"Content-Type",
			conf.GlobalConfigs.Notif.SMS.Send.Verify.ContentType,
		},
	}

	// New Client
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// Start POST Request
	request, err := http.NewRequest("POST", conf.GlobalConfigs.Notif.SMS.Token.URL, bytes.NewBuffer(body))
	for _, v := range headers {
		request.Header.Set(v.Header, v.Vaule)
	}

	// Client try To Send
	resp, err := client.Do(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Send SMS: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Send SMS: %s", err.Error()))
		return err
	}
	defer resp.Body.Close()

	// Try To Read Response
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Read Response Body After Send SMS: %s", err.Error()),
			"msg":   string(bt),
		}).Warn(fmt.Sprintf("Error in Read Response Body After Send SMS: %s", err.Error()))
		return err
	}

	// Try To Unmarshal Data
	if err := json.Unmarshal(bt, &s); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Unmarshal Response Data For Sended SMS: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Unmarshal Response Data For Sended SMS: %s", err.Error()))
		return err
	}

	// Check if IsSuccessful or Not!
	if s.IsSuccessful == false {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Send SMS: %s", string(bt)),
			"body":  string(bt),
		}).Warn(fmt.Sprintf("Error in Send SMS: %s", string(bt)))
		return err
	}

	return nil
}

func (s *SMS) send(method, url string, headers []struct {
	Header string
	Vaule  string
}, body []byte) error {

	// Get Token From Service if TokenKey is empty
	if err := s.token(); err != nil {
		return err
	}

	// New Client
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// Start POST Request
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	for _, v := range headers {
		request.Header.Set(v.Header, v.Vaule)
	}
	request.Header.Set("x-sms-ir-secure-token", s.TokenKey)

	// Client try To Send
	resp, err := client.Do(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Send SMS: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Send SMS: %s", err.Error()))
		return err
	}
	defer resp.Body.Close()

	// Try To Read Response
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Read Response Body After Send SMS: %s", err.Error()),
			"msg":   string(bt),
		}).Warn(fmt.Sprintf("Error in Read Response Body After Send SMS: %s", err.Error()))
		return err
	}

	// Try To Unmarshal Data
	if err := json.Unmarshal(bt, &s); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Unmarshal Response Data For Sended SMS: %s", err.Error()),
		}).Warn(fmt.Sprintf("Error in Unmarshal Response Data For Sended SMS: %s", err.Error()))
		return err
	}

	// Check if IsSuccessful or Not!
	if s.IsSuccessful == false {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Send SMS: %s", string(bt)),
		}).Warn(fmt.Sprintf("Error in Send SMS"))
		return err
	}

	return nil
}
