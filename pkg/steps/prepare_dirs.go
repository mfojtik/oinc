package steps

import (
	"os"
	"path/filepath"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

var BaseDir = filepath.Join(os.Getenv("HOME"), ".openshift", "oinc")
var OpenShiftVolumes = []string{
	"openshift.local.volumes", "openshift.local.config", "openshift.local.etcd",
}

type PrepareDirsStep struct {
	ParallelStep
}

func (*PrepareDirsStep) String() string { return "prepare-dir" }

func (*PrepareDirsStep) Execute() error {
	for _, path := range OpenShiftVolumes {
		path = filepath.Join(BaseDir, path)
		if _, err := os.Stat(path); err == nil {
			log.Info("Directory %q already exists. Skipping ...", path)
			continue
		}
		if err := os.MkdirAll(path, 0700); err != nil {
			return err
		}
		util.RunSudoCommand("chcon", "-R", "-t", "svirt_sandbox_file_t", path)
	}
	return nil
}
