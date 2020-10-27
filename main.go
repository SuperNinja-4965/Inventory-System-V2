package main

import (
	"log"
	"os"
	"path/filepath"
)

// Import my own library of useful functions
// Define the Program Wide variables.
// ExecPath - the path where the program is stored.
var ExecPath string

// The main function. The program will enter here.
func main() {
	// Getting the programs executable path.
	var err2 error
	ExecPath, err2 = filepath.Abs(filepath.Dir(os.Args[0]))
	// check if there was an error with the program. If there was an error then kill the program
	if err2 != nil {
		log.Fatal(err2)
	}

	// Call the read settings function to get the programs settings from the settings file.
	ReadSettings()

}
