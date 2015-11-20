package steps

import (
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/mfojtik/oinc/pkg/log"
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
	ParallelStep
}

func (*PullDockerImages) String() string { return "pull-images" }

func (*PullDockerImages) Execute() error {
	log.Info("Pulling %d Docker images ...", len(pullImages))
	uiprogress.Start()
	bar := uiprogress.AddBar(len(pullImages))
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return pullImages[b.Current()-1]
	})
	index := 0
	for bar.Incr() {
		// _ := pullImages[index]
		time.Sleep(2 * time.Second)
		index++
	}
	uiprogress.Stop()
	return nil
}
