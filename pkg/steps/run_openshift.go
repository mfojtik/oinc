package steps

import (
	"path/filepath"
	"time"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

type RunOpenShiftStep struct {
	DefaultStep
}

func (*RunOpenShiftStep) String() string { return "restart-docker" }

func volume(path string, private bool) string {
	if private {
		return filepath.Join(BaseDir, path) + ":" + "/var/lib/origin/" + path + ":z"
	}
	return filepath.Join(BaseDir, path) + ":" + filepath.Join(BaseDir, path) + ":z"
}

func runOpenShift(args ...string) error {
	return util.RunSudoCommand("docker", "run", "-d", "--name", OpenShiftContainerName,
		"--privileged",
		"--net", "host",
		"--pid", "host",
		"-v", "/:/rootfs:ro",
		"-v", "/sys:/sys:ro",
		"-v", "/var/run:/var/run:rw",
		"-v", "/var/lib/docker:/var/lib/docker:rw",
		"-v", volume(OpenShiftPublicVolumes[0], false),
		"-v", volume(OpenShiftPrivateVolumes[0], true),
		"-v", volume(OpenShiftPrivateVolumes[1], true),
		"openshift/origin", "start",
		"--nodes=127.0.0.1",
		"--hostname=localhost",
		"--volume-dir", filepath.Join(BaseDir, OpenShiftPublicVolumes[0]),
		"--cors-allowed-origins=.*",
	)
}

func (*RunOpenShiftStep) Execute() error {
	if util.RunSudoCommand("docker", "inspect", OpenShiftContainerName) == nil {
		log.Info("OpenShift container is already running")
		return nil
	}
	// When an error occurs, display logs and remove the failed container
	log.Info("Starting OpenShift server at https://%s:8443/console ...", util.GetHostIP())
	err := runOpenShift()
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
	} else {
		// FIXME: This should be really a health-check
		time.Sleep(15 * time.Second)
	}
	return err
}
