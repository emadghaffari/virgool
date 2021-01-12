package notif

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// SMS struct
type SMS struct {
	TokenKey     string
	IsSuccessful bool
}

type params struct {
	Parameter      string
	ParameterValue interface{}
}

// Send meth
func (s *SMS) Send(ctx context.Context, notif Notification, code int) error {
	t, err := s.token()
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	body, err := json.Marshal(map[string]interface{}{
		"ParameterArray": []params{
			{Parameter: "Code", ParameterValue: code},
		},
		"Mobile":     "09355960597",
		"TemplateId": "22108",
	})
	if err != nil {
		logrus.Warn(err)
		return err
	}

	request, err := http.NewRequest("POST", "https://RestfulSms.com/api/UltraFastSend", bytes.NewBuffer(body))
	request.Header.Set("x-sms-ir-secure-token", t)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		logrus.Warn(err)
		return err
	}
	defer resp.Body.Close()

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		logrus.Warn(err)
		return err
	}

	return nil
}

func (s *SMS) token() (string, error) {
	req, err := json.Marshal(map[string]string{
		"UserApiKey": "dd198b76e5ea1d1ef31f8b76",
		"SecretKey":  "cp6teBC!@FeBC!YFuBC",
	})
	if err != nil {
		logrus.Warn(err)
		return "", err
	}

	resp, err := http.Post("https://RestfulSms.com/api/Token", "application/json", bytes.NewBuffer(req))
	if err != nil {
		logrus.Warn(err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Warn(err)
		return "", err
	}

	if err := json.Unmarshal(body, &s); err != nil {
		logrus.Warn(err)
		return "", err
	}

	return s.TokenKey, nil
}
