package conf

import (
	"fmt"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	once sync.Once
)

// ConfigureLogging func
func ConfigureLogging(config *LoggingConfig) {
	once.Do(func() {

		// always use the full timestamp
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:    true,
			DisableTimestamp: false,
			TimestampFormat:  time.RFC3339Nano,
			DisableColors:    config.DisableColors,
			QuoteEmptyFields: config.QuoteEmptyFields,
		})

		log.SetLevel(log.InfoLevel)

		if GlobalConfigs.Environment == "production" {
			// Log as JSON instead of the default ASCII formatter.
			f, _ := os.OpenFile(fmt.Sprintf("blog/logs/%s.log", time.Now().Local().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
			log.SetFormatter(&log.JSONFormatter{})
			log.SetOutput(f)
			return
		}

		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})

		// Output to stdout instead of the default stderr
		log.SetOutput(os.Stdout)

	})
}
