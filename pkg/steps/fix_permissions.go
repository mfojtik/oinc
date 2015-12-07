package steps

import (
	"os"

	"github.com/mfojtik/oinc/pkg/util"
)

type FixPermissionsStep struct {
	DefaultStep
}

func (*FixPermissionsStep) String() string { return "fix-permissions" }

func (*FixPermissionsStep) Execute() error {
	return util.RunSudoCommand("chown", "-R", os.Getenv("USER"), util.BaseDir)
}
