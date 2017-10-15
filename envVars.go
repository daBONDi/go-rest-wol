package main

import (
	"log"
	"os"
)

// Processing Shell Arguments
func processEnvVars() string {

	computerFile := DefaultComputerFilePath

	if os.Getenv(DefaultComputerFilePathEnvironmentName) != "" {
		computerFile = os.Getenv(DefaultComputerFilePathEnvironmentName)
	}

	// Try to Stat the Filepath from Env Variable if not Fail Fatal
	if _, err := os.Stat(computerFile); os.IsNotExist(err) {
		log.Fatalf("Could not find or access Computerlist Environment Variable WOLFILE: \"%s\"", computerFile)
	}

	return computerFile
}
