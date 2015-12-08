package util

import "runtime"

func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}
