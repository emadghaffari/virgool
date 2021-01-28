package kf

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/segmentio/kafka-go"

	"github.com/emadghaffari/virgool/notification/conf"
)

var (
	// Database var
	Database kf = &Client{}
	once     sync.Once
	err      error
)

// kf interface
type kf interface {
	Connect(ctx context.Context,conf *conf.GlobalConfiguration,topic string,partition int) error
	Consumer(ctx context.Context, reader *kafka.Reader, action func(kafka.Message))
	NewReader(groupID,topic string, partition int) (*kafka.Reader,error)
	TopicList() (map[string]struct{},error)
}

// Client struct
type Client struct {
	Config   *kafka.Conn
	Messages chan *sarama.ConsumerMessage
	Errors   chan *sarama.ConsumerError
	ready    chan bool
	Count    int
}

// Connect is basic connection to kafka
func (k *Client) Connect(ctx context.Context,conf *conf.GlobalConfiguration,topic string,partition int) error {
	once.Do(func() {
		k.Config, err = kafka.DialLeader(ctx, "tcp", conf.Kafka.Brokers[0], topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}
	})
	
	return err
}

// NewReader func
func (k *Client) NewReader(groupID,topic string, partition int) (*kafka.Reader,error) {
	r:=  kafka.NewReader(kafka.ReaderConfig{
		Brokers:   conf.GlobalConfigs.Kafka.Brokers,
		Topic:     topic,
		GroupID:   groupID,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	return r,nil
}

// Consumer for a topic 
func (k *Client) Consumer(ctx context.Context, reader *kafka.Reader, action func(kafka.Message)){
	for{
		// m,err := reader.ReadMessage(ctx)
		m,err := reader.FetchMessage(ctx)
		if err != nil {
			break
		}
		action(m)
		if err := reader.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// TopicList get the list of topics for a broker
func (k *Client) TopicList() (map[string]struct{},error) {
	if k.Config == nil {
		return nil,fmt.Errorf("Error in get Config")
	}
	partitions, err := k.Config.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	return m,nil
}