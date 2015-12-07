package steps

import (
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

type DisableFirewalldStep struct {
	DefaultStep
}

func (*DisableFirewalldStep) String() string { return "disable-firewalld" }

func (*DisableFirewalldStep) Execute() error {
	log.Info("Stopping the firewalld daemon")
	return util.RunSudoCommand("systemctl", "stop", "firewalld")
}
