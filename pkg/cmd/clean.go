package cmd

import (
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/spf13/cobra"
)

var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up the oinc directories.",
	Long:  `Clean up the oinc directories and containers`,
	Run: func(cmd *cobra.Command, args []string) {
		setupLogging()
		cleanup := &steps.CleanupStep{}
		if err := cleanup.Execute(); err != nil {
			log.Critical("Error: %v", err)
		}
	},
}

func init() {
	addPersistentFlags(CleanCmd)
}
