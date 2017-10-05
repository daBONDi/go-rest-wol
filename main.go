package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ComputerList contains all Computers who we can use to work with
var ComputerList []Computer

// Find and Wakeup Computer
func WakeUp(ComputerName string, List []Computer) (string, string) {

	for _, c := range List {
		if c.Name == ComputerName {
			err := SendMagicPacket(c.Mac, c.BroadcastIPAddress, "")
			if err != nil {
				// TODO Refactor with Error Object
				return "ERROR", "WOL Module Error"
			}

			return "OK", ""
		}
	}

	return "ERROR", "Invalid Computer Name"
}

func main() {

	// Start Processing Shell Arguments or use Default Values defined i const.go
	httpPort, computerFilePath := processShellArgs()

	// Loading Computer CSV File to Memory File in Memory
	var loadComputerCSVFileError error
	if ComputerList, loadComputerCSVFileError = loadComputerList(computerFilePath); loadComputerCSVFileError != nil {
		log.Fatalf("Error on loading Computerlist File \"%s\" check File access and formating", computerFilePath)
	}

	// Init HTTP Router - mux
	router := mux.NewRouter()

	// Define Home Route
	router.HandleFunc("/", renderHomePage).Methods("GET")

	// Define Wakeup Api functions with a Computer Name
	router.HandleFunc("/api/wakeup/computer/{computerName}", restWakeUpWithComputerName).Methods("GET")
	router.HandleFunc("/api/wakeup/computer/{computerName}/", restWakeUpWithComputerName).Methods("GET")

	// Setup Webserver
	httpListen := fmt.Sprint(":", httpPort)
	log.Printf("Startup Webserver on \"%s\"", httpListen)

	log.Fatal(http.ListenAndServe(httpListen, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router)))
}
