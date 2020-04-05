package main

import (
	"log"
	"os"
	"strconv"
)

// Processing Shell Arguments
func processEnvVars() (int, string, string, string) {

	computerFile := DefaultComputerFilePath
	port := DefaultHTTPPort
	username := UsernameForShutdown
    password := PasswordForShutdown

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

	if os.Getenv(DefaultComputerUsernameEnvironmentName) != "" {
    	username = os.Getenv(DefaultComputerUsernameEnvironmentName)
    }

    if os.Getenv(DefaultComputerPasswordEnvironmentName) != "" {
        password = os.Getenv(DefaultComputerPasswordEnvironmentName)
    }

	return port, computerFile, username, password
}
