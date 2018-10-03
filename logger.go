package main

import (
	"fmt"
	"os"

	logging "github.com/op/go-logging"
)

func init() {
	logFile, err := os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE, 0666)

	err = logFile.Truncate(0)
	_, err = logFile.Seek(0, 0)

	if err != nil {
		fmt.Println("Could not load errors.log file")
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

	// Set the backends to be used.
	logging.SetBackend(backendConsole, backendLogfileLeveled)
}
