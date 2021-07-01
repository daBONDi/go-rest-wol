/* Processing and handling Shell Arguments */

package main

import (
	"flag"
	"log"
)

// Processing Shell Arguments
func processShellArgs() (int, string) {
	// Reading Shell Args
	port := flag.Int("port", DefaultHTTPPort,
		"Define the port on which the webserver will listen to")
	filePath := flag.String("file", DefaultComputerFilePath,
		"Path to the CSV file containing the list of hosts ")

	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name]=true } )


	// Try to Stat the File if not Fail Fatal
	if flagset["file"] && !FileExists(*filePath) {
		log.Fatalf("Could not find or access Computerlist --file Argument Value: \"%s\"", *filePath)
	}

	return *port, *filePath
}
