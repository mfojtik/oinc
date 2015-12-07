package steps

import (
	"path/filepath"

	"github.com/mfojtik/oinc/pkg/util"
)

type InstallRegistryStep struct {
	DefaultStep
}

func (*InstallRegistryStep) String() string { return "install-registry" }

func (*InstallRegistryStep) Execute() error {
	_, err := util.RunOAdm("registry",
		"--create",
		"--credentials", filepath.Join(util.MasterConfigPath, "openshift-registry.kubeconfig"),
	)
	return err
}
