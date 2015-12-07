package main

import (
	"github.com/mfojtik/oinc/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	RootCmd.AddCommand(cmd.SetupCmd)
	RootCmd.AddCommand(cmd.RunCmd)
	RootCmd.AddCommand(cmd.CleanCmd)
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "oinc",
	Short: "OpenShift-in-Container installer",
	Long:  "Install and run the OpenShift server in Docker container.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
