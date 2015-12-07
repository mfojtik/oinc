package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mfojtik/oinc/pkg/steps"
	"github.com/spf13/cobra"
)

var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Print environment settings for running CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout,
			"export PATH="+path.Join(steps.BaseDir, "bin")+":$PATH\n",
		)
	},
}
