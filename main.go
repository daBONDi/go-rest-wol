package main

import (
	"fmt"
	"log"
	"net/http"
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

	fmt.Print("Loading Computer CSV File")
	ComputerList := LoadComputerListCSVFile("./computer.csv")
	fmt.Print("Done")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		paramComputerName := r.URL.Query().Get("computer")
		if paramComputerName != "" {
			result, err := WakeUp(paramComputerName, ComputerList)

			if err != "" {
				fmt.Fprintf(w, "<p>WOL Result: %s ERROR: %s</p>", paramComputerName, err)
			} else {
				fmt.Fprintf(w, "<p>Computer: %s %s</p>", paramComputerName, result)
			}

		}

		// Building Table Output here
		fmt.Fprint(w, GetComputerListTable(ComputerList))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
