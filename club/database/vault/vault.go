package vault

import (
	"fmt"
	"sync"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/club/conf"
)

var (
	// Database var
	Database Vault = &vault{}
	once     sync.Once
	err      error
)

// Vault interface
type Vault interface {
	New(config conf.GlobalConfiguration) error
	Read(path string) (*api.Secret, error)
	Write(path string, data map[string]interface{}) (*api.Secret, error)
}
type vault struct {
	DB *api.Logical
}

func (v *vault) New(config conf.GlobalConfiguration) error {
	// config.Confs.Notifs.Path = "blog/notificator"
	// config.Confs.Users.Host = "localhost"
	// config.Confs.Users.Path = "blog/users"
	// config.Confs.JWT.Path = "blog/jwt/secret"
	// config.Confs.Redis.Path = "blog/redis"

	once.Do(func() {
		confs := &api.Config{
			Address: config.Vault.Address,
		}

		var client *api.Client
		client, err = api.NewClient(confs)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": fmt.Sprintf("Failed to connect to vault: %s", err),
			}).Fatal(fmt.Sprintf("Failed to connect to vault: %s", err))
			return
		}
		client.SetToken(config.Vault.Token)

		v.DB = client.Logical()
	})
	return err
}

func (v *vault) Read(path string) (*api.Secret, error) {
	r, err := v.DB.Read(path)
	if err != nil {
		logrus.Warn(err.Error())
		return nil, err
	}

	return r, nil
}

func (v *vault) Write(path string, data map[string]interface{}) (*api.Secret, error) {
	w, err := v.DB.Write(path, data)
	if err != nil {
		logrus.Warn(err.Error())
		return nil, err
	}

	return w, nil
}
