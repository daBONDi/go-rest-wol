/* Processing and handling Shell Arguments */

package main

import (
	"log"
	"os"
	"strconv"
)

// Processing Shell Arguments
func processShellArgs() (int, string) {
	// Reading Shell Args
	shellArgs := os.Args[1:]

	filepath := DefaultComputerFilePath
	port := DefaultHTTPPort

	for i, arg := range shellArgs {

		// Process Shell Args --port
		if arg == "--port" {

			// Ensure Port Argument has a following up Argument
			if i+1 <= len(shellArgs) {

				// Ensure we can Parse the Port Argument Value
				var err error
				if port, err = strconv.Atoi(shellArgs[i+1]); err != nil {
					log.Fatalf("Could not find or parse ArgumentValue --port to Integer \"%s\"", shellArgs[i+1])
				}
			}
		}

		// Process Shell Args --file
		if arg == "--file" {
			// Ensure we got a shell Argument afterwards
			if i+1 <= len(shellArgs) {
				// Ensure passed File Argument Value is something
				if len(shellArgs[i+1]) > 0 {
					filepath = shellArgs[i+1]
				}
			}
		}
	}

	// Try to Stat the File if not Fail Fatal
	if !FileExists(filepath) {
		log.Fatalf("Could not find or access Computerlist --file Argument Value: \"%s\"", filepath)
	}

	return port, filepath
}
