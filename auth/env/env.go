package env

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/database/vault"
)

func localEnvironment(filename string) error {

	// name of config file (without extension)
	// REQUIRED if the config file does not have the extension in the name
	// path to look for the config file in
	viper.SetConfigFile(filename)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logrus.WithFields(logrus.Fields{
				"error": fmt.Sprintf("Config file not found; ignore error if desired: %s", err),
			}).Fatal(fmt.Sprintf("Config file not found; ignore error if desired: %s", err))
		} else {
			// Config file was found but another error was produced
			logrus.WithFields(logrus.Fields{
				"error": fmt.Sprintf("Config file was found but another error was produced: %s", err),
			}).Fatal(fmt.Sprintf("Config file was found but another error was produced: %s", err))
		}
		return err
	}

	if err := viper.Unmarshal(&conf.GlobalConfigs); err != nil {
		// Config file can not unmarshal to struct
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Config file can not unmarshal to struct: %s", err),
		}).Fatal(fmt.Sprintf("Config file can not unmarshal to struct: %s", err))

		return err
	}

	return nil
}

func vaultEnvironment() error {
	conf.GlobalConfigs.Vault.Address = os.Getenv("vault_address")
	conf.GlobalConfigs.Vault.Token = os.Getenv("vault_token")

	// Vault connection
	if err := vault.Database.New(conf.GlobalConfigs); err != nil {
		fmt.Fprintf(os.Stderr, ": %v\n", err)
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("Failed to connect to vault: %s", err),
		}).Fatal(fmt.Sprintf("Failed to connect to vault: %s", err))
		return err
	}
	return nil
}

// LoadGlobalConfiguration returns configs
func LoadGlobalConfiguration(filename string) error {
	if os.Getenv("environment") == "production" {
		return vaultEnvironment()
	}

	if err := localEnvironment(filename); err != nil {
		return err
	}

	return nil
}
