package cmd

import (
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/spf13/cobra"
)

var PullImages bool

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up the system for running the OpenShift server",
	Long:  `Sets up the host system for running the OpenShift server in a container.`,
	Run: func(cmd *cobra.Command, args []string) {
		setupLogging()
		preConfig := &steps.PreConfigStep{}
		if err := preConfig.Execute(); err != nil {
			log.Critical("%s", err)
		}

		images := &steps.ImagesStep{PullImages: PullImages}
		if err := images.Execute(); err != nil {
			log.Critical("%s", err)
		}
	},
}

func init() {
	addPersistentFlags(SetupCmd)
	SetupCmd.PersistentFlags().BoolVarP(&PullImages, "pull", "p", true, "Pull Docker images if they do not exist locally")
}
