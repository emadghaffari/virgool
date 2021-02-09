package conf

import "time"

// GlobalConfigs var
var GlobalConfigs GlobalConfiguration

// GlobalConfiguration holds all the global configuration for gocommerce
type GlobalConfiguration struct {
	Environment string `yaml:"environment"`
	GRPC        struct {
		Host     string `yaml:"grpc.host"`
		Port     string `yaml:"grpc.port"`
		Endpoint string `yaml:"grpc.endpoint"`
	}
	HTTP struct {
		Host     string `yaml:"http.host"`
		Port     string `yaml:"http.port"`
		Endpoint string `yaml:"http.endpoint"`
	}
	DEBUG struct {
		Host     string `yaml:"debug.host"`
		Port     string `yaml:"debug.port"`
		Endpoint string `yaml:"debug.endpoint"`
	}
	MYSQL             DBConfiguration
	Redis             RedisConfiguration
	Vault             VaultConfiguration
	MultiInstanceMode bool `yaml:"multiInstanceMode"`
	Log               LoggingConfig
	Service           Service
	Jaeger            Jaeger
	Kafka             Kafka
	JWT               JWT
	ELK               ELK
}

// DBConfiguration struct
type DBConfiguration struct {
	Username    string `yaml:"mysql.username"`
	Password    string `yaml:"mysql.password"`
	Host        string `yaml:"mysql.host"`
	Schema      string `yaml:"mysql.schema"`
	Driver      string `yaml:"mysql.driver"`
	Automigrate bool   `yaml:"mysql.automigrate"`
	Logger      bool   `yaml:"mysql.logger"`
	Namespace   string
}

// RedisConfiguration struct
type RedisConfiguration struct {
	Username string `yaml:"redis.username"`
	Password string `yaml:"redis.password"`
	DB       int    `yaml:"redis.db"`
	Host     string `yaml:"redis.host"`
	Logger   bool   `yaml:"redis.logger"`
}

// VaultConfiguration struct
type VaultConfiguration struct {
	Address string `yaml:"vault.address"`
	Token   string `yaml:"vault.token"`
}

// LoggingConfig struct
type LoggingConfig struct {
	DisableColors    bool `json:"disable_colors" yaml:"log.disableColors"`
	QuoteEmptyFields bool `json:"quote_empty_fields" yaml:"log.quoteEmptyFields"`
}

// Service details
type Service struct {
	Name  string `yaml:"service.name"`
	Redis struct {
		SMSDuration         time.Duration `yaml:"service.redis.smsDuration"`
		SMSCodeVerification time.Duration `yaml:"service.redis.smsCodeVerification"`
		UserDuration        time.Duration `yaml:"service.redis.userDuration"`
	}
}

// Jaeger tracer
type Jaeger struct {
	HostPort string `yaml:"jaeger.hostPort"`
	LogSpans bool   `yaml:"jaeger.logSpans"`
}

// Kafka struct
type Kafka struct {
	Username string   `yaml:"kafka.username"`
	Password string   `yaml:"kafka.password"`
	Brokers  []string `yaml:"kafka.brokers"`
	Version  string   `yaml:"kafka.version"`
	Group    string   `yaml:"kafka.group"`
	Assignor string   `yaml:"kafka.assignor"`
	Oldest   bool     `yaml:"kafka.oldest"`
	Verbose  bool     `yaml:"kafka.verbose"`
	Topics   Topic    `yaml:"kafka.topics"`
	Auth     bool     `yaml:"kafka.auth"`
	Consumer bool     `yaml:"kafka.consumer"`
	Producer bool     `yaml:"kafka.producer"`
}

// Topic struct
type Topic struct {
	Notif string `yaml:"kafka.topics.notif"`
}

// JWT struct
type JWT struct {
	RSecret string `yaml:"jwt.rSecret"`
	Secret  string `yaml:"jwt.secret"`
}

// ELK , elasticsearch - kibana - logstash
type ELK struct {
	URLs            []string      `yaml:"elk.urls"`
	Username        string        `yaml:"elk.username"`
	Password        string        `yaml:"elk.password"`
	HealthCheckTime time.Duration `yaml:"elk.health_check"`
}
