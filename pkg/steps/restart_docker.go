package steps

import "github.com/mfojtik/oinc/pkg/util"

type RestartDockerStep struct {
	DefaultStep
}

func (*RestartDockerStep) String() string { return "restart-docker" }

func (*RestartDockerStep) Execute() error {
	if util.IsDarwin() {
		return nil
	}
	return util.RunSudoCommand("systemctl", "restart", "docker")
}
