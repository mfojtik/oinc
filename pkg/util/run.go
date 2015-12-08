package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mfojtik/oinc/pkg/log"
)

var (
	BaseDir          = filepath.Join(os.Getenv("HOME"), ".openshift", "oinc")
	MasterConfigPath = filepath.Join(BaseDir, "openshift.local.config", "master")
)

func BaseDirPath(cmd string) string {
	return filepath.Join(BaseDir, "bin", cmd)
}

func RunSudoCommand(path string, args ...string) error {
	_, err := GetSudoCommandOutput(path, args...)
	return err
}

func GetSudoCommandOutput(path string, args ...string) (string, error) {
	// Don't do sudo on OSX
	if runtime.GOOS == "darwin" {
		return RunCommand(path, args...)
	}
	sudoPath, err := exec.LookPath("sudo")
	if err != nil {
		return "", err
	}
	return RunCommand(sudoPath, append([]string{path}, args...)...)
}

func RunCommand(path string, args ...string) (string, error) {
	c := exec.Command(path, args...)
	log.Debug(path + " " + strings.Join(args, " "))
	out, err := c.CombinedOutput()
	if len(strings.TrimSpace(string(out))) > 0 {
		log.Debug("%q returned %q", path+" "+strings.Join(args, " "), string(out))
	}
	return strings.TrimSpace(string(out)), err
}

func RunOAdm(args ...string) (string, error) {
	args = append(args, []string{"--config", filepath.Join(MasterConfigPath, "admin.kubeconfig")}...)
	return GetSudoCommandOutput(BaseDirPath("oadm"), args...)
}

func RunAdminOc(args ...string) (string, error) {
	args = append(args, []string{"--config", filepath.Join(MasterConfigPath, "admin.kubeconfig")}...)
	return RunOc(args...)
}

func RunOc(args ...string) (string, error) {
	return RunCommand(BaseDirPath("oc"), args...)
}
