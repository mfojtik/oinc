package cmd

import (
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the OpenShift server in a container.",
	Long:  `Runs the OpenShift server in a container`,
	Run: func(cmd *cobra.Command, args []string) {
		setupLogging()
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
	addPersistentFlags(RunCmd)
}
