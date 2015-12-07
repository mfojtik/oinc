package steps

import (
	"path/filepath"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

type CleanupStep struct {
	DefaultStep
}

func (*CleanupStep) String() string { return "cleanup" }

func (*CleanupStep) Execute() error {
	if err := util.RunSudoCommand("docker", "kill", OpenShiftContainerName); err != nil {
		log.Debug("Failed to kill %q container: %v", OpenShiftContainerName, err)
	}
	rm := &CleanupExistingStep{}
	if err := rm.Execute(); err != nil {
		return err
	}
	for _, path := range append(OpenShiftPublicVolumes, OpenShiftPrivateVolumes...) {
		if err := util.RunSudoCommand("rm", "-r", "-f", filepath.Join(BaseDir, path)); err != nil {
			log.Debug("Failed to remove %q: %v", path, err)
		}
	}
	return nil
}
