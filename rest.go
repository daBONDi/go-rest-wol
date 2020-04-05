// Rest API Implementations

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"os/exec"

	"github.com/gorilla/mux"
)

//restWakeUpWithComputerName - REST Handler for Processing URLS /api/computer/<computerName>
func restWakeUpWithComputerName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	computerName := vars["computerName"]

	var result WakeUpResponseObject
	result.Success = false

	// Ensure computerName is not empty
	if computerName == "" {
		result.Message = "Empty Computername is not allowed"
		result.ErrorObject = nil
		w.WriteHeader(http.StatusBadRequest)
		// Computername is empty
	} else {

		// Get Computer from List
		for _, c := range ComputerList {
			if c.Name == computerName {

				// We found the Computername
				if err := SendMagicPacket(c.Mac, c.BroadcastIPAddress, ""); err != nil {
					// We got an internal Error on SendMagicPacket
					w.WriteHeader(http.StatusInternalServerError)
					result.Success = false
					result.Message = "Internal error on Sending the Magic Packet"
					result.ErrorObject = err
				} else {
					// Horray we send the WOL Packet succesfully
					result.Success = true
					result.Message = fmt.Sprintf("Succesfully Wakeup Computer %s with Mac %s on Broadcast IP %s", c.Name, c.Mac, c.BroadcastIPAddress)
					result.ErrorObject = nil
				}
			}
		}

		if result.Success == false && result.ErrorObject == nil {
			// We could not find the Computername
			w.WriteHeader(http.StatusNotFound)
			result.Message = fmt.Sprintf("Computername %s could not be found", computerName)
		}
	}
	json.NewEncoder(w).Encode(result)
}

//restShutdownWithComputerName - REST Handler for Processing URLS /api/computer/<computerName>
func restShutdownWithComputerName(w http.ResponseWriter, r *http.Request) {

	username := UsernameForShutdown
	password := PasswordForShutdown

	// Process Environment Variables
	_, _, username, password = processEnvVars()

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	computerName := vars["computerName"]

	var result WakeUpResponseObject
	result.Success = false

	// Ensure computerName is not empty
	if computerName == "" {
		result.Message = "Empty Computername is not allowed"
		result.ErrorObject = nil
		w.WriteHeader(http.StatusBadRequest)
		// Computername is empty
	} else {

		// Get Computer from List
		for _, c := range ComputerList {
			if c.Name == computerName {

				// We found the Computername
				//
				var namepluspass = username + "%" + password
                   cmd := exec.Command("net", "rpc", "shutdown","-I",c.IPAddress, "-U", namepluspass)
                   	cmd.Stdout = os.Stdout
                   	cmd.Stderr = os.Stderr
                   	err := cmd.Run()
                   	if err != nil {
                   		log.Fatalf("cmd.Run() failed with %s\n", err)
                   	}

                   result.Success = true
                   result.Message = "Shutdown successfull!"
                   result.ErrorObject = err


			}
		}

		if result.Success == false && result.ErrorObject == nil {
			// We could not find the Computername
			w.WriteHeader(http.StatusNotFound)
			result.Message = fmt.Sprintf("Computername %s could not be found", computerName)
		}
	}
	json.NewEncoder(w).Encode(result)
}
