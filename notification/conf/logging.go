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

// ConfigureLogging func
func ConfigureLogging(config *LoggingConfig) {

	once.Do(func() {
		Logger = logrus.New()

		// always use the full timestamp
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    true,
			DisableTimestamp: false,
			TimestampFormat:  time.RFC3339Nano,
			DisableColors:    config.DisableColors,
			QuoteEmptyFields: config.QuoteEmptyFields,
		})

		Logger.SetLevel(logrus.InfoLevel)

		if viper.GetString("environment") == "production" {
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
