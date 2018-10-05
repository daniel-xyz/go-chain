package main

import (
	"os"

	logging "github.com/op/go-logging"
)

const initFailed = "Could not initialize logging. This is a bad sign!"

func init() {
	logFile, err := os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE, 0600)

	if err != nil {
		panic(initFailed)
	}

	if err := logFile.Truncate(0); err != nil {
		panic(initFailed)
	}

	if _, err := logFile.Seek(0, 0); err != nil {
		panic(initFailed)
	}

	consoleFormatter := logging.MustStringFormatter(
		`%{color}%{message}%{color:reset}`,
	)

	logfileFormatter := logging.MustStringFormatter(
		`%{time:15:04:05.000} %{longpkg} %{shortfunc}() ▶ %{level} %{id:03x} %{message}`,
	)

	backendConsole := logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), consoleFormatter)
	backendLogfile := logging.NewBackendFormatter(logging.NewLogBackend(logFile, "", 0), logfileFormatter)

	backendLogfileLeveled := logging.AddModuleLevel(backendLogfile)
	backendLogfileLeveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backendConsole, backendLogfileLeveled)
}
