package util

import (
	"os/exec"
	"strings"

	"github.com/mfojtik/oinc/pkg/log"
)

func RunSudoCommand(path string, args ...string) error {
	sudoPath, err := exec.LookPath("sudo")
	if err != nil {
		return err
	}
	return RunCommand(sudoPath, append([]string{path}, args...)...)
}

func RunCommand(path string, args ...string) error {
	c := exec.Command(path, args...)
	log.Debug(path + " " + strings.Join(args, " "))
	out, err := c.CombinedOutput()
	if len(strings.TrimSpace(string(out))) > 0 {
		log.Debug("%q returned %q", path+" "+strings.Join(args, " "), string(out))
	}
	return err
}
