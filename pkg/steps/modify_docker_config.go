package steps

import "github.com/mfojtik/oinc/pkg/log"

// FIXME: This will be different on Windows
const SystemDockerConfigPath = "/etc/sysconfig/docker"

type ModifyDockerConfigStep struct {
	DefaultStep
}

func (*ModifyDockerConfigStep) String() string { return "modify-docker-config" }

func (*ModifyDockerConfigStep) Execute() error {
	log.Info("Copying %q to %q", SystemDockerConfigPath, SystemDockerConfigPath+".backup")
	return nil
}
