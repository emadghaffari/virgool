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
	Redis             RedisConfiguration
	Vault             VaultConfiguration
	MultiInstanceMode bool `yaml:"multiInstanceMode"`
	Log               LoggingConfig
	Service           Service
	Jaeger            Jaeger
	Kafka             Kafka
	Notif             Notif
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
	Name string `yaml:"service.name"`

	// min code lenght
	MinCL int `yaml:"service.mincl"`
	MaxCl int `yaml:"service.maxcl"`

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

// Notif struct
type Notif struct {
	SMS   sms
	Email email
}

type sms struct {
	UserAPIKey string `yaml:"notif.sms.userApiKey"`
	SecretKey  string `yaml:"notif.sms.secretKey"`
	Token      struct {
		URL         string `yaml:"notif.sms.token.url"`
		ContentType string `yaml:"notif.sms.token.contentType"`
	}

	Send struct {
		TemplateURL string   `yaml:"notif.sms.send.templateURL"`
		BodyURL     string   `yaml:"notif.sms.send.bodyURL"`
		LineNumber  []string `yaml:"notif.sms.send.lineNumber"`
		Verify      struct {
			TemplateID  string `yaml:"notif.sms.send.verify.templateId"`
			ContentType string `yaml:"notif.sms.send.verify.contentType"`
		}
	}
}

type email struct {
	Driver   string `yaml:"notif.email.driver"`
	Host     string `yaml:"notif.email.host"`
	Port     int    `yaml:"notif.email.port"`
	Username string `yaml:"notif.email.username"`
	Password string `yaml:"notif.email.password"`
	Identity string `yaml:"notif.email.identity"`
	Send     struct {
		Template string `yaml:"notif.email.send.template"`
	}
}
