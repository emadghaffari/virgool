package conf

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	once sync.Once
)

// ConfigureLogging func
func ConfigureLogging(config *LoggingConfig) {

	once.Do(func() {

		// always use the full timestamp
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    true,
			DisableTimestamp: false,
			TimestampFormat:  time.RFC3339Nano,
			DisableColors:    config.DisableColors,
			QuoteEmptyFields: config.QuoteEmptyFields,
		})

		logrus.SetLevel(logrus.InfoLevel)

		if GlobalConfigs.Environment == "production" {
			// Log as JSON instead of the default ASCII formatter.
			f, _ := os.OpenFile(fmt.Sprintf("%s/logs/%s.log", GlobalConfigs.Service.Name ,time.Now().Local().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.SetOutput(f)
			return
		}

		// The TextFormatter is default, you don't actually have to do this.
		logrus.SetFormatter(&logrus.TextFormatter{})

		// Output to stdout instead of the default stderr
		logrus.SetOutput(os.Stdout)

	})
}
