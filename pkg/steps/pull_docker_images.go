package steps

import (
	"fmt"
	"strings"
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

// pullImages defines the Docker image to pre-pull to speed up OpenShift
// installation. Normally you would have to pull just openshift/origin, but
// then OpenShift server will pull them in when they are needed.
var pullImages = []string{
	"openshift/base-centos7",
	"openshift/origin-base",
	"openshift/origin-pod",
	"openshift/origin-deployer",
	"openshift/origin-docker-registry",
	"openshift/origin-haproxy-router",
	"openshift/origin-sti-builder",
	"openshift/origin",
}

type PullDockerImages struct {
	PullImages bool
	DefaultStep
}

func (*PullDockerImages) String() string { return "pull-images" }

func (s *PullDockerImages) Execute() error {
	time.Sleep(1 * time.Second)
	log.Info("Pulling OpenShift Docker images ...")
	uiprogress.Start()
	bar := uiprogress.AddBar(len(pullImages))
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("[%d/%d] ", b.Current(), len(pullImages)) + strings.TrimPrefix(pullImages[b.Current()-1], "openshift/")
	})
	index := 0
	errors := map[string]error{}
	for bar.Incr() {
		if !s.PullImages && isAvailable(pullImages[index]) {
			log.Debug("Image %q available locally, skipping pull ...", pullImages[index])
			index++
			continue
		}
		err := util.RunSudoCommand("docker", "pull", pullImages[index])
		if err != nil {
			errors[pullImages[index]] = err
		}
		index++
	}
	uiprogress.Stop()
	if len(errors) > 0 {
		for name, err := range errors {
			log.Error("Error pulling %q: %v", name, err)
		}
		return fmt.Errorf("Images failed to pull")
	}
	return nil
}

func isAvailable(name string) bool {
	return util.RunSudoCommand("docker", "inspect", name) == nil
}
