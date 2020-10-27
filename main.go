package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	JG "github.com/joshuag1000/GoEssentials"
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

	// Call the StartUp function to check if the needed files and directories exist and they are where they need to be.
	StartUp()

	// Call the InitPages function to initialise the pages.
	InitPages()

	// Print the Server IP to the console as well as the server ports.
	fmt.Println("The server ip is: " + JG.GetServerIP(0))
	fmt.Println("The http port is: " + NonHttpsPort + "\nThe https port is: " + SitePort)

	// Checks the setting to see if the user wants the browser to open when the sever starts. If they do open it.
	if openBrowserOnLoad == true {
		JG.OpenBrowser("http://localhost:" + NonHttpsPort)
		JG.OpenBrowser("https://localhost:" + SitePort)
	}
	// Uses my GoEssentials Library to start the StartWebServer
	JG.StartWebServer(NonHttpsPort, SitePort)
}
