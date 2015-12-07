package cmd

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var (
	LogLevel int
)

func setupLogging() {
	if LogLevel > 6 || LogLevel < 0 {
		fmt.Fprintf(os.Stderr, "Allowed log levels are between 0 and 6, you set %d", LogLevel)
	}
	logging.SetLevel(logging.Level(LogLevel), "")
}

func addPersistentFlags(c *cobra.Command) {
	c.PersistentFlags().IntVarP(&LogLevel, "loglevel", "v", 4, "Set the verbosity level (0-5), default: 4")
}
