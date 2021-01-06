package vault

import (
	"sync"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/auth/conf"
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
	// config.Confs.Users.DebugAddr = *debugAddr
	// config.Confs.Users.HTTPAddr = *httpAddr
	// config.Confs.Users.GrpcAddr = *grpcAddr
	// config.Confs.Users.ThriftAddr = *thriftAddr

	once.Do(func() {
		confs := &api.Config{
			Address: config.Vault.Address,
		}

		var client *api.Client
		client, err = api.NewClient(confs)
		if err != nil {
			logrus.Warn(err.Error())
			return
		}
		client.SetToken(config.Vault.Token)

		v.DB = client.Logical()
	})

	// // Read notif path
	// notifs, err := v.DB.Read(config.Confs.Notifs.Path)
	// if err != nil {
	// 	logrus.Warn(err.Error())
	// 	return err
	// }
	// config.Confs.Notifs.Host = notifs.Data["grpc"].(string)

	// // Read jwt secret
	// jwt, err := v.DB.Read(config.Confs.JWT.Path)
	// if err != nil {
	// 	logrus.Warn(err.Error())
	// 	return err
	// }
	// config.Confs.JWT.Secret = jwt.Data["jwt"].(string)
	// config.Confs.JWT.RSecret = jwt.Data["rjwt"].(string)

	// Read jwt secret
	// rd, err := v.DB.Read(config.Confs.Redis.Path)
	// if err != nil {
	// 	logrus.Warn(err.Error())
	// 	return err
	// }
	// config.Confs.Redis.Host = rd.Data["host"].(string)
	// config.Confs.Redis.DB = rd.Data["db"].(string)

	// Write users Path
	// _, err = v.DB.Write(config.Confs.Users.Path, map[string]interface{}{
	// 	"debug":  config.Confs.Users.Host + config.Confs.Users.DebugAddr,
	// 	"http":   config.Confs.Users.Host + config.Confs.Users.HTTPAddr,
	// 	"grpc":   config.Confs.Users.Host + config.Confs.Users.GrpcAddr,
	// 	"thrift": config.Confs.Users.Host + config.Confs.Users.ThriftAddr,
	// })
	// if err != nil {
	// 	logrus.Warn(err.Error())
	// 	return err
	// }

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
