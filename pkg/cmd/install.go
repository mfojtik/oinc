package cmd

import (
	"fmt"
	"os"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var (
	LogLevel   int
	PullImages bool
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Download and install OpenShift v3",
	Long: `Configure the host system to run OpenShift v3 in container and bootstrap OpenShift
server to be ready to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		if LogLevel > 6 || LogLevel < 0 {
			fmt.Fprintf(os.Stderr, "Allowed log levels are between 0 and 6, you set %d", LogLevel)
		}
		logging.SetLevel(logging.Level(LogLevel), "")

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
	InstallCmd.PersistentFlags().IntVarP(&LogLevel, "loglevel", "v", 4, "Set the verbosity level (0-5), default: 4")
	InstallCmd.PersistentFlags().BoolVarP(&PullImages, "pull", "p", true, "Pull Docker images if they does not exists locally")
}
