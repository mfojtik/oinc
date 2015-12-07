package cmd

import (
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/spf13/cobra"
)

var PullImages bool

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setups the pre-requires for running OpenShift server",
	Long:  `Setups the host system to run OpenShift server in container.`,
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
	SetupCmd.PersistentFlags().BoolVarP(&PullImages, "pull", "p", true, "Pull Docker images if they does not exists locally")
}
