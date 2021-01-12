package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/notification/conf"
	"github.com/emadghaffari/virgool/notification/model/notif"
	"github.com/emadghaffari/virgool/notification/pkg/service"
)

var (
	// Database var
	Database kf = &Client{}
	once     sync.Once
	err      error
)

// kf interface
type kf interface {
	Connect(conf *conf.GlobalConfiguration) error
	Consumer(ctx context.Context, brokers []string, topic string) (sarama.Consumer, error)
}

// Client struct
type Client struct {
	Config   *sarama.Config
	Messages chan *sarama.ConsumerMessage
	Errors   chan *sarama.ConsumerError
	ready    chan bool
	Count    int
}

// Connect to kafka
func (k *Client) Connect(conf *conf.GlobalConfiguration) error {
	if err := k.validate(conf); err != nil {
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

		k.Config = config
		k.Count = conf.Service.MinCL
	})

	return err
}

// Consumer meth
func (k *Client) Consumer(ctx context.Context, brokers []string, topic string) (sarama.Consumer, error) {
	client, err := sarama.NewConsumer(brokers, k.Config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Error in newConsumer: %s", err),
		}).Fatal(fmt.Sprintf("Error in newConsumer: %s", err))
		return nil, err
	}

	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)

	partitions, err := client.Partitions(topic)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	consumer, err := client.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	go func(topic string, consumer sarama.PartitionConsumer) {
		for {
			select {
			case consumerError := <-consumer.Errors():
				errors <- consumerError

			case msg := <-consumer.Messages():
				consumers <- msg
			}
		}
	}(topic, consumer)

	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case msg := <-consumers:
				id := fmt.Sprintf("%v-%d-%d", msg.Topic, msg.Partition, msg.Offset)

				var item notif.Notification
				if err := json.Unmarshal(msg.Value, &item); err != nil {
					logrus.WithFields(logrus.Fields{
						"error": fmt.Sprintf("cannot unmarshal data from kafka Error: %s - id: %s", err, id),
					}).Fatal(fmt.Sprintf("cannot unmarshal data from kafka Error: %s - id: %s", err, id))
				}

				service.Streamer.Store(context.Background(), k.Count, item)
				k.Count++

				if k.Count == conf.GlobalConfigs.Service.MaxCl {
					k.Count = conf.GlobalConfigs.Service.MinCL
				}
			case consumerError := <-errors:
				logrus.Warn(string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
			}
		}
	}()

	return client, nil
}
