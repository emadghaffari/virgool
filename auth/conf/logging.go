package conf

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once sync.Once

	// Logger var
	Logger *logrus.Logger
)

// LoggingConfig struct
type LoggingConfig struct {
	Level            string                 `mapstructure:"log_level" json:"log_level"`
	File             string                 `mapstructure:"log_file" json:"log_file"`
	DisableColors    bool                   `mapstructure:"disable_colors" split_words:"true" json:"disable_colors"`
	QuoteEmptyFields bool                   `mapstructure:"quote_empty_fields" split_words:"true" json:"quote_empty_fields"`
	TSFormat         string                 `mapstructure:"ts_format" json:"ts_format"`
	Fields           map[string]interface{} `mapstructure:"fields" json:"fields"`
	UseNewLogger     bool                   `mapstructure:"use_new_logger",split_words:"true"`
}

// ConfigureLogging func
func ConfigureLogging(config *LoggingConfig) {

	once.Do(func() {
		Logger = logrus.New()

		tsFormat := time.RFC3339Nano
		if config.TSFormat != "" {
			tsFormat = config.TSFormat
		}
		// always use the full timestamp
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    true,
			DisableTimestamp: false,
			TimestampFormat:  tsFormat,
			DisableColors:    config.DisableColors,
			QuoteEmptyFields: config.QuoteEmptyFields,
		})

		Logger.SetLevel(logrus.InfoLevel)

		if viper.GetString("Environment") == "production" {
			// Log as JSON instead of the default ASCII formatter.
			f, _ := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Local().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
			Logger.SetFormatter(&logrus.JSONFormatter{})
			Logger.SetOutput(f)
		}

		// The TextFormatter is default, you don't actually have to do this.
		Logger.SetFormatter(&logrus.TextFormatter{})

		// Output to stdout instead of the default stderr
		Logger.SetOutput(os.Stdout)

	})
}
