package util

import (
	"strings"

	"github.com/mfojtik/oinc/pkg/log"
)

func GetHostIP() string {
	var out string
	var err error
	if IsDarwin() {
		out, err = RunCommand("ipconfig", "getifaddr", "en0")
	} else {
		out, err = RunCommand("hostname", "-I")
	}
	if err != nil {
		log.Error("Unable to obtain the host IP address: %v", err)
	}
	p := strings.Split(out, " ")
	return p[0]
}
