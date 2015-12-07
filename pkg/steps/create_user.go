package steps

import (
	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

type CreateUserStep struct {
	DefaultStep
}

func (*CreateUserStep) String() string { return "create-user" }

func (*CreateUserStep) Execute() error {
	log.Info("Creating 'test-user' with password 'test' ...")
	_, err := util.RunOAdm("policy", "add-role-to-user", "view", "test-admin")
	if err != nil {
		return err
	}
	out, err := util.RunOc("login", "https://"+util.GetHostIP()+":8443", "-u", "test-admin", "-p", "test", "--api-version=v1")
	if err != nil {
		return err
	}
	log.Info(out)
	out, _ = util.RunOc("project", "devel")
	log.Info(out)
	return nil
}
