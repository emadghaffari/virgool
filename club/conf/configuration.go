package conf

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
	Vault             VaultConfiguration
	MultiInstanceMode bool `yaml:"multiInstanceMode"`
	Log               LoggingConfig
	Service           Service
	Jaeger            Jaeger
	Kafka             Kafka
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
