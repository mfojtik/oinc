package steps

import "github.com/mfojtik/oinc/pkg/log"
import "time"

type PullDockerImages struct {
	ParallelStep
}

func (*PullDockerImages) String() string { return "pull-images" }

func (*PullDockerImages) Execute() error {
	log.Info("Pulling Docker images ...")
	time.Sleep(5 * time.Second)
	return nil
}
