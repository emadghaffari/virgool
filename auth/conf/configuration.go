package conf

// GlobalConfigs var
var GlobalConfigs GlobalConfiguration

// GlobalConfiguration holds all the global configuration for gocommerce
type GlobalConfiguration struct {
	GRPC struct {
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
	Vault             VaultConfiguration
	MultiInstanceMode bool `yaml:"multiInstanceMode"`
	Log               LoggingConfig
	Service           Service
	Jaeger            Jaeger
}

// DBConfiguration struct
type DBConfiguration struct {
	Username    string `yaml:"mysql.username"`
	Password    string `yaml:"mysql.password"`
	Host        string `yaml:"mysql.host"`
	Schema      string `yaml:"mysql.schema"`
	Driver      string `yaml:"mysql.driver"`
	Automigrate bool   `yaml:"mysql.automigrate"`
	Namespace   string
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
