package main

import (
	"github.com/mfojtik/oinc/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	RootCmd.AddCommand(cmd.ExecuteCmd)
	RootCmd.AddCommand(cmd.InstallCmd)
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use: "oinc",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
