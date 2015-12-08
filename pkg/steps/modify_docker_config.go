package steps

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

// FIXME: This will be different on Windows
const SystemDockerConfigPath = "/etc/sysconfig/docker"
const RegistryAddress = "172.30.0.0/16"

type ModifyDockerConfigStep struct {
	DefaultStep
}

func (*ModifyDockerConfigStep) String() string { return "modify-docker-config" }

func (*ModifyDockerConfigStep) Execute() error {
	if util.IsDarwin() {
		return nil
	}
	file, err := os.Open(SystemDockerConfigPath)
	if err != nil {
		return err
	}
	defer file.Close()

	newFile, err := ioutil.TempFile("", "oinc.docker")
	if err != nil {
		return err
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(file)
	isUpdated := false
	for scanner.Scan() {
		data := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(data, "INSECURE_REGISTRY=") {
			if strings.Contains(data, RegistryAddress) {
				log.Debug("Docker Registry %q already configured in %q", RegistryAddress, SystemDockerConfigPath)
				return nil
			} else {
				// Deal with different suffixes
				suffix := ""
				if strings.HasSuffix(data, "'") {
					suffix = "'"
				}
				if strings.HasSuffix(data, `'`) {
					suffix = `'`
				}
				data = strings.TrimSuffix(data, suffix)
				data += " --insecure-registry " + RegistryAddress + suffix
				log.Debug("Using %q", data)
			}
		}
		isUpdated = true
		newFile.WriteString(data + "\n")
	}

	// There is no INSECURE_REGISTRY defined yet, so append one on the bottom
	if !isUpdated {
		log.Debug("Adding INSECURE_REGISTRY entry for " + RegistryAddress)
		newFile.WriteString("INSECURE_REGISTRY='--insecure-registry " + RegistryAddress + "'\n")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return util.RunSudoCommand("mv", "-f", newFile.Name(), SystemDockerConfigPath)
}
