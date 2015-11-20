package log

import (
	"os"

	"github.com/op/go-logging"
)

const format = "%{color}%{time:15:04:05.000} %{level:.6s} ▶ %{color:reset} %{message}"

var (
	log           = logging.MustGetLogger("oinc")
	formatter     = logging.MustStringFormatter(format)
	isInitialized = false
)

func init() {
	if isInitialized {
		return
	}

	infoBackend := logging.NewLogBackend(os.Stderr, "", 0)
	errBackend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(infoBackend, formatter)

	// Send just errors to stderr
	errLeveled := logging.AddModuleLevel(errBackend)
	errLeveled.SetLevel(logging.ERROR, "oinc")
	logging.SetBackend(errLeveled, backendFormatter)
	isInitialized = true
}

// Logging helpers
func Debug(format string, args ...interface{})    { log.Debug(format, args...) }
func Info(format string, args ...interface{})     { log.Info(format, args...) }
func Notice(format string, args ...interface{})   { log.Notice(format, args...) }
func Warning(format string, args ...interface{})  { log.Warning(format, args...) }
func Error(format string, args ...interface{})    { log.Error(format, args...) }
func Critical(format string, args ...interface{}) { log.Critical(format, args...) }
func Panic(format string, args ...interface{})    { log.Panicf(format, args...) }
