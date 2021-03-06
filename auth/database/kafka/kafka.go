package kafka

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/auth/conf"
)

var (
	// Database var
	Database Kafka = &kf{}
	once     sync.Once
	err      error
)

// Kafka interface
type Kafka interface {
	Connect(config *conf.GlobalConfiguration) error
	Validate(config *conf.GlobalConfiguration) error
	syncProducer() (sarama.SyncProducer, error)
	Producer(item interface{}, topic string) error
}

type kf struct {
	Config *sarama.Config
}

func (k *kf) Connect(conf *conf.GlobalConfiguration) error {
	if err := k.Validate(conf); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in GlobalConfig: %s", err),
		}).Fatal(fmt.Sprintf("Error in GlobalConfig: %s", err))
		return err
	}

	once.Do(func() {
		// newconfig
		config := sarama.NewConfig()

		// clientID is service name
		config.ClientID = conf.Service.Name

		// config.Net
		config.Net.MaxOpenRequests = 1

		if conf.Kafka.Consumer {
			// config.Consumer
			config.Consumer.Return.Errors = true
			config.Consumer.MaxProcessingTime = time.Second
			config.Consumer.Fetch.Max = 500
			config.Consumer.Fetch.Min = 10
			config.Consumer.Group.Heartbeat.Interval = time.Second * 5
			config.Consumer.Group.Session.Timeout = time.Second * 15
			switch conf.Kafka.Assignor {
			case "sticky":
				config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
			case "roundrobin":
				config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
			case "range":
				config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
			default:
				logrus.WithFields(logrus.Fields{
					"error": fmt.Sprintf("Unrecognized consumer group partition assignor: %s", conf.Kafka.Assignor),
				}).Fatal(fmt.Sprintf("Unrecognized consumer group partition assignor: %s", conf.Kafka.Assignor))
				err = fmt.Errorf("Unrecognized consumer group partition assignor: %s", conf.Kafka.Assignor)
			}

		}

		// if kafka has SASL auth
		if conf.Kafka.Auth {
			// Auth
			config.Net.SASL.Enable = true
			config.Net.SASL.Handshake = true
			config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
			config.Net.SASL.User = conf.Kafka.Username
			config.Net.SASL.Password = conf.Kafka.Password
			config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }

		}

		if conf.Kafka.Producer {
			// config.Producer
			config.Producer.Idempotent = true
			config.Producer.Return.Errors = true
			config.Producer.RequiredAcks = sarama.WaitForAll
			config.Producer.Return.Successes = true
			config.Producer.Retry.Backoff = time.Duration(time.Second * 5)
			config.Producer.Retry.Max = 5
			config.Producer.Compression = sarama.CompressionLZ4
			config.Producer.Timeout = time.Duration(time.Second * 50)
		}

		k.Config = config
	})

	return err
}

// SyncProducer func
func (k *kf) syncProducer() (sarama.SyncProducer, error) {
	syncProducer, err := sarama.NewSyncProducer(conf.GlobalConfigs.Kafka.Brokers, k.Config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": fmt.Sprintf("can not syncProducer kafka brokers: %s", conf.GlobalConfigs.Kafka.Brokers),
			"error":   fmt.Sprintf("Error: %s", err),
		}).Fatal(fmt.Sprintf("can not syncProducer kafka brokers: %s Error: %s", conf.GlobalConfigs.Kafka.Brokers, err))
		return nil, err
	}
	return syncProducer, nil
}

// Producer func
func (k *kf) Producer(item interface{}, topic string) error {
	bt, err := json.Marshal(item)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": fmt.Sprintf("can not marshal data: %s", item),
			"error":   fmt.Sprintf("Error: %s", err),
		}).Fatal(fmt.Sprintf("can not marshal data: %s", item))
		return fmt.Errorf("can not marshal data: %s", item)
	}

	syncProducer, err := k.syncProducer()
	if err != nil {
		return err
	}

	_, _, err = syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(bt)),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": fmt.Sprintf("failed to send message to: %s || Topics:%s", topic, conf.GlobalConfigs.Kafka.Topics),
			"error":   fmt.Sprintf("Error: %s", err),
		}).Fatal(fmt.Sprintf("Error: %s", err))
		return fmt.Errorf("failed to send message to: %s Error:%s", topic, err)
	}

	return syncProducer.Close()
}
