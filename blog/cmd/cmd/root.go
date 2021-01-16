package cmd

import (
	"github.com/spf13/cobra"
)

var configFile = ""

// rootCmd will run the log streamer
var rootCmd = cobra.Command{
	Use:  "virgool",
	Long: "A service that will validate restful transactions and send them to stripe.",
	Run: func(cmd *cobra.Command, args []string) {
		// service.Run()
	},
}

// RootCmd will add flags and subcommands to the different commands
func RootCmd() *cobra.Command {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "The configuration file")
	rootCmd.AddCommand(&seedCMD)
	return &rootCmd
}
