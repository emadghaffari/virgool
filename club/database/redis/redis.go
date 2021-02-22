package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/club/conf"
)

var (
	// Database var
	Database Redis = &rd{}

	// sync
	once sync.Once

	// error
	err error
)

// Redis interface
type Redis interface {
	Connect(config *conf.GlobalConfiguration) error
	GetDatabase() *redis.Client
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Del(ctx context.Context, key ...string) error
}

type rd struct {
	db *redis.Client
}

func (s *rd) Connect(config *conf.GlobalConfiguration) error {
	once.Do(func() {
		s.db = redis.NewClient(&redis.Options{
			DB:       config.Redis.DB,
			Addr:     config.Redis.Host,
			Username: config.Redis.Username,
			Password: config.Redis.Password,
		})

		if err = s.db.Ping(context.Background()).Err(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": fmt.Sprintf("Config redis database Error: %s", err),
			}).Fatal(fmt.Sprintf("Config redis database Error: %s", err))
		}
	})

	return err
}
func (s *rd) GetDatabase() *redis.Client {
	return s.db
}

// Set meth a new key,value
func (s *rd) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Marshal Error for Set New item in Redis: %s", err),
		}).Fatal(fmt.Sprintf("Marshal Error for Set New item in Redis: %s", err))
		return err
	}
	return s.db.Set(ctx, key, p, duration).Err()
}

// Get meth, get value with key
func (s *rd) Get(ctx context.Context, key string, dest interface{}) error {
	p, err := s.db.Get(ctx, key).Result()

	if p == "" {
		return fmt.Errorf("Value Not Found")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in Get value from Redis: %s", err),
		}).Fatal(fmt.Sprintf("Error in Get value from Redis: %s", err))
		return err
	}

	return json.Unmarshal([]byte(p), &dest)
}

func (s *rd) Del(ctx context.Context, key ...string) error {
	_, err := s.db.Del(ctx, key...).Result()
	if err != nil {
		return err
	}
	return nil
}
