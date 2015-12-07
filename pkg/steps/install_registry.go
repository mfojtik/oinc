package steps

import (
	"path/filepath"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

type InstallRegistryStep struct {
	DefaultStep
}

func (*InstallRegistryStep) String() string { return "install-registry" }

func (*InstallRegistryStep) Execute() error {
	log.Info("Installing Docker Registry ...")
	_, err := util.RunOAdm("registry",
		"--create",
		"--credentials", filepath.Join(util.MasterConfigPath, "openshift-registry.kubeconfig"),
	)
	return err
}
