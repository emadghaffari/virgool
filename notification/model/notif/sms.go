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

// SMSParams struct
type SMSParams struct {
	Parameter      string
	ParameterValue interface{}
}

// Send meth
func (s *SMS) Send(ctx context.Context, notif Notification, params []SMSParams, to, template string) error {
	fmt.Println(to)
	fmt.Println(template)
	t, err := s.token()
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	body, err := json.Marshal(map[string]interface{}{
		"ParameterArray": params,
		"Mobile":         to,
		"TemplateId":     template,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Marshal Send SMS Data: %s", err),
		}).Warn(fmt.Sprintf("Error in Marshal Send SMS Data: %s", err))
		return err
	}

	request, err := http.NewRequest("POST", conf.GlobalConfigs.Notif.SMS.Send.TemplateURL, bytes.NewBuffer(body))
	request.Header.Set("x-sms-ir-secure-token", t)
	request.Header.Set("Content-Type", conf.GlobalConfigs.Notif.SMS.Send.Verify.ContentType)

	resp, err := client.Do(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Send SMS: %s", err),
		}).Warn(fmt.Sprintf("Error in Send SMS: %s", err))
		return err
	}
	defer resp.Body.Close()

	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Read Response Body After Send SMS: %s", err),
			"msg":   string(bt),
		}).Warn(fmt.Sprintf("Error in Read Response Body After Send SMS: %s", err))
		return err
	}

	if err := json.Unmarshal(bt, &s); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Unmarshal Response Data For Sended SMS: %s", err),
		}).Warn(fmt.Sprintf("Error in Unmarshal Response Data For Sended SMS: %s", err))
		return err
	}

	if s.IsSuccessful == false {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Send SMS: %s", err),
			"body":  string(bt),
		}).Warn(fmt.Sprintf("Error in Send SMS: %s", err))
		return err
	}

	return nil
}

func (s *SMS) token() (string, error) {
	req, err := json.Marshal(map[string]string{
		"UserApiKey": conf.GlobalConfigs.Notif.SMS.UserAPIKey,
		"SecretKey":  conf.GlobalConfigs.Notif.SMS.SecretKey,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Marshal Data sms.go: %s", err),
		}).Warn(fmt.Sprintf("Error in Marshal Data sms.go: %s", err))
		return "", err
	}

	resp, err := http.Post(conf.GlobalConfigs.Notif.SMS.Token.URL, conf.GlobalConfigs.Notif.SMS.Token.ContentType, bytes.NewBuffer(req))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Post Data For Get SMS Token: %s", err),
		}).Warn(fmt.Sprintf("Error in Post Data For Get SMS Token: %s", err))
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error ioutil Read Data: %s", err),
		}).Warn(fmt.Sprintf("Error ioutil Read Data: %s", err))
		return "", err
	}

	if err := json.Unmarshal(body, &s); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Unmarshal Response Post Token SMS: %s", err),
		}).Warn(fmt.Sprintf("Error in Unmarshal Response Post Token SMS: %s", err))
		return "", err
	}

	if s.IsSuccessful == false {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Get Send SMS Token: %s", err),
			"body":  string(body),
		}).Warn(fmt.Sprintf("Error in Get Send SMS Token: %s", err))
		return "", err
	}

	return s.TokenKey, nil
}
