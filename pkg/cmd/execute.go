package cmd

import (
	"fmt"
	"os"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var ExecuteCmd = &cobra.Command{
	Use:   "execute",
	Short: "oinc is fully automated oppenshift-in-container installer",
	Long: `Configure the host system to run OpenShift v3 in container and bootstrap OpenShift
server to be ready to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		if LogLevel > 6 || LogLevel < 0 {
			fmt.Fprintf(os.Stderr, "Allowed log levels are between 0 and 6, you set %d", LogLevel)
		}
		logging.SetLevel(logging.Level(LogLevel), "")

		server := &steps.RunOpenShiftStep{}
		if err := server.Execute(); err != nil {
			log.Critical("%s", err)
		}

		registry := &steps.InstallRegistryStep{}
		if err := registry.Execute(); err != nil {
			log.Critical("%s", err)
		}

	},
}

func init() {
	ExecuteCmd.PersistentFlags().IntVarP(&LogLevel, "loglevel", "v", 4, "Set the verbosity level (0-5), default: 4")
}
