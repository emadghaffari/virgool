package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	"github.com/emadghaffari/virgool/blog/env"
)

var seedCMD = cobra.Command{
	Use:  "seed database",
	Long: "seed database strucutures. This will seed tables",
	Run:  seed,
}

func seed(cmd *cobra.Command, args []string) {

	// Current working directory
	dir, err := os.Getwd()
	if err != nil {
		logrus.Warn(err.Error())
	}

	// read from file
	if err := env.LoadGlobalConfiguration(dir + "/blog/config.yaml"); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	conf.GlobalConfigs.MYSQL.Automigrate = false
	conf.GlobalConfigs.MYSQL.Logger = false

	if err := mysql.Database.Connect(&conf.GlobalConfigs, logrus.New()); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// db := mysql.Database.GetDatabase()
	// TODO add seeder
}
