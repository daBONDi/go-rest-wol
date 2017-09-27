package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// LoadComputerListCSVFile loads the Computer into ComputerList
func LoadComputerListCSVFile(computerCsvFilePath string) []Computer {

	computerCsvFilePointer, computerCsvFileError := os.Open("./computer.csv")
	if computerCsvFileError != nil {
		log.Fatal(computerCsvFileError)
	}
	defer computerCsvFilePointer.Close() // gets executed until the function finished, like a garbage collector ?? :-)

	var computers []Computer

	if err := gocsv.UnmarshalFile(computerCsvFilePointer, &computers); err != nil { // Load clients from file
		fmt.Println("Parsing Error, pls check File")
		panic(err)
	}

	return computers
}
