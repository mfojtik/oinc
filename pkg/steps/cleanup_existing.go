package steps

import (
	"fmt"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

const OpenShiftContainerName = "origin-oinc"

type CleanupExistingStep struct {
	DefaultStep
}

func (*CleanupExistingStep) String() string { return "cleanup-existing" }

func (*CleanupExistingStep) Execute() error {
	r, err := util.GetSudoCommandOutput("docker", "inspect", "-f", "{{.State.Running}}", OpenShiftContainerName)
	if err != nil {
		log.Debug("Container %q not found (%q), skipping cleanup", OpenShiftContainerName, r)
		return nil
	}
	if r == "true" {
		log.Notice("The %q container is running, skipping cleanup")
		log.Notice("Run $ docker rm %s", OpenShiftContainerName)
		return fmt.Errorf("Running %q container found", OpenShiftContainerName)
	}
	return util.RunSudoCommand("docker", "rm", "-f", "-v", OpenShiftContainerName)
}
