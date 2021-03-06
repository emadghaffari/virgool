package main

import (
	"fmt"
	"os"

	"github.com/emadghaffari/virgool/auth/cmd/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run command: %v\n", err)
		os.Exit(1)
	}
}
