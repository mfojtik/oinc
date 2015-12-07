package util

import (
	"strings"

	"github.com/mfojtik/oinc/pkg/log"
)

func GetHostIP() string {
	out, err := RunCommand("hostname", "-I")
	if err != nil {
		log.Error("Unable to obtain host IP address: %v", err)
	}
	p := strings.Split(out, " ")
	return p[0]
}
