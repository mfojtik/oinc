package steps

import (
	"path/filepath"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

type RunOpenShiftStep struct {
	DefaultStep
}

func (*RunOpenShiftStep) String() string { return "restart-docker" }

func (*RunOpenShiftStep) Execute() error {
	err := util.RunSudoCommand("docker", "run", "-d", "--name", OpenShiftContainerName, "--privileged",
		"--net", "host", "--pid", "host",
		"-v", "/:/rootfs:ro",
		"-v", "/sys:/sys:ro",
		"-v", "/var/run:/var/run:rw",
		"-v", "/var/lib/docker:/var/lib/docker:rw",
		"-v", filepath.Join(BaseDir, OpenShiftVolumes[0])+":/var/lib/origin/"+OpenShiftVolumes[0]+":z",
		"-v", filepath.Join(BaseDir, OpenShiftVolumes[1])+":/var/lib/origin/"+OpenShiftVolumes[1]+":z",
		"-v", filepath.Join(BaseDir, OpenShiftVolumes[2])+":/var/lib/origin/"+OpenShiftVolumes[2]+":z",
		"openshift/origin", "start",
		"master", "--etcd-dir", "/var/lib/origin/"+OpenShiftVolumes[2],
		"--cors-allowed-origins=.*",
	)
	// When an error occurs, display logs and remove the failed container
	if err != nil {
		out, logsErr := util.GetSudoCommandOutput("docker", "logs", OpenShiftContainerName)
		if logsErr != nil {
			log.Error("Unable to get logs from %q container: %v", OpenShiftContainerName, logsErr)
			return err
		}
		log.Debug(out)
		rmErr := util.RunSudoCommand("docker", "rm", "-f", "-v", OpenShiftContainerName)
		if rmErr != nil {
			log.Error("Unable to remove %q container: %v", OpenShiftContainerName, rmErr)
		}
	}
	return err
}
