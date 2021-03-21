package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daBONDi/go-rest-wol/internal/repository"
	"github.com/daBONDi/go-rest-wol/pkg/cmd"
	"github.com/daBONDi/go-rest-wol/pkg/http/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	httpPort := cmd.DefaultHTTPPort
	computerFilePath := cmd.DefaultComputerFilePath

	// Start Processing Shell Arguments or use Default Values defined i const.go
	httpPort, computerFilePath = cmd.ProcessShellArgs()

	// Process Environment Variables
	httpPort, computerFilePath = cmd.ProcessEnvVars()

	// Loading Computer CSV File to Memory File in Memory
	var loadComputerCSVFileError error
	if repository.ComputerList, loadComputerCSVFileError = repository.LoadComputerList(computerFilePath); loadComputerCSVFileError != nil {
		log.Fatalf("Error on loading Computerlist File \"%s\" check File access and formating", computerFilePath)
	}

	// Init HTTP Router - mux
	router := mux.NewRouter()

	// Define Home Route
	router.HandleFunc("/", handler.RenderHomePage).Methods("GET")

	// Define Wakeup Api functions with a Computer Name
	router.HandleFunc("/api/wakeup/computer/{computerName}", handler.RestWakeUpWithComputerName).Methods("GET")
	router.HandleFunc("/api/wakeup/computer/{computerName}/", handler.RestWakeUpWithComputerName).Methods("GET")

	// Setup Webserver
	httpListen := fmt.Sprint(":", httpPort)
	log.Printf("Startup Webserver on \"%s\"", httpListen)

	log.Fatal(http.ListenAndServe(httpListen, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router)))
}
