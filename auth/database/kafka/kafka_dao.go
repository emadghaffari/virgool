package kafka

import (
	"fmt"

	"github.com/emadghaffari/virgool/auth/conf"
)

func (k *kf) Validate(config *conf.GlobalConfiguration) error {
	if len(config.Kafka.Brokers) == 0 {
		return fmt.Errorf("kafka need Brokers, please set the Brokers")
	}
	if config.Kafka.Version == "" {
		return fmt.Errorf("kafka need Version, please set the Version")
	}
	if config.Kafka.Topics.Notif == "" {
		return fmt.Errorf("kafka need Topics, please set the Topics")
	}
	if config.Kafka.Assignor == "" {
		return fmt.Errorf("kafka need Assignor, please set the Assignor")
	}
	return nil
}
