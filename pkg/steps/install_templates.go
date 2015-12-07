package steps

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gosuri/uiprogress"
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

var jbossVersion = "ose-v1.1.0"
var templates = []string{
	"https://raw.githubusercontent.com/openshift/origin/master/examples/image-streams/image-streams-rhel7.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/jboss-image-streams.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mongodb-ephemeral-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mongodb-persistent-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mysql-ephemeral-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mysql-persistent-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/postgresql-ephemeral-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/postgresql-persistent-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/jenkins/jenkins-ephemeral-template.json",
	"https://raw.githubusercontent.com/openshift/origin/master/examples/jenkins/jenkins-persistent-template.json",
	"https://raw.githubusercontent.com/openshift/nodejs-ex/master/openshift/templates/nodejs-mongodb.json",
	"https://raw.githubusercontent.com/openshift/nodejs-ex/master/openshift/templates/nodejs.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-amq-persistent-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-amq-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-basic-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-https-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-mongodb-persistent-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-mongodb-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-mysql-persistent-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-mysql-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-postgresql-persistent-s2i.json",
	"https://raw.githubusercontent.com/jboss-openshift/application-templates/" + jbossVersion + "/eap/eap64-postgresql-s2i.json",
}

type InstallTemplatesStep struct {
	DefaultStep
}

func (*InstallTemplatesStep) String() string { return "install-templates" }

func downloadAndInstall(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	t, err := ioutil.TempFile("", "template")
	defer os.Remove(t.Name())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(t.Name(), []byte(body), 0660)
	if err != nil {
		return err
	}
	_, err = util.RunAdminOc("create", "-n", "openshift", "-f", t.Name())
	return err
}

func (*InstallTemplatesStep) Execute() error {
	log.Info("Installing %d templates ...", len(templates))
	uiprogress.Start()
	bar := uiprogress.AddBar(len(templates) - 1)
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		name := templates[b.Current()]
		name = name[strings.LastIndex(name, "/")+1:]
		name = strings.TrimSuffix(name, ".json")
		return fmt.Sprintf("[%s]", name)
	})
	errors := []error{}
	var wg sync.WaitGroup
	for _, url := range templates {
		wg.Add(1)
		go func(u string) {
			if err := downloadAndInstall(u); err != nil {
				errors = append(errors, fmt.Errorf("%q failed to download: %v", u, err))
			}
			bar.Incr()
			wg.Done()
		}(url)
	}
	wg.Wait()
	if len(errors) > 0 {
		for _, e := range errors {
			log.Error("%v", e)
		}
		return fmt.Errorf("Some templates failed to download")
	}
	uiprogress.Stop()
	return nil
}
