package main

import (
	"log"
	"os"
	"strconv"
)

// Processing Shell Arguments
func processEnvVars() (int, string) {

	computerFile := DefaultComputerFilePath
	port := DefaultHTTPPort

	if os.Getenv(DefaultComputerFilePathEnvironmentName) != "" {
		computerFile = os.Getenv(DefaultComputerFilePathEnvironmentName)
		if !FileExists(computerFile) {
			log.Fatalf("Environmentvariable \"%s\" is set but Value is not a Path to an existing File: %s", DefaultComputerFilePathEnvironmentName, computerFile)
		}
	}

	if os.Getenv(DefaultHTTPPortEnvironmentVariableName) != "" {
		var err error
		if port, err = strconv.Atoi(os.Getenv(DefaultHTTPPortEnvironmentVariableName)); err != nil {
			log.Fatalf("Environmentvariable \"%s\" should be a integer", DefaultHTTPPortEnvironmentVariableName)
		}
	}
	return port, computerFile
}
