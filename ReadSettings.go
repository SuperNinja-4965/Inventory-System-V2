package main

import (
	"bufio"
	"os"
	"strconv"
)

// Default Settings - These are the default settings that the program uses.
// If the settings file does not exist these are the values the program will use.
var DefaultSettings string = "// Settings file for Inventory System each setting must have a space after the colon or it will be ignored.\n\nProgram-Name: Inventory System\nHTTPS-PORT: 8443\nHTTP-PORT: 8080\n\n// options are true or false\nOpenBrowser: false"

// Public variables to store the settings
var ProgramName string = "Inventory System (Failed to load settings)"
var openBrowserOnLoad bool = true
var SitePort string = "443"
var NonHttpsPort string = "80"

// When called the function will fetch the settings from the settings file
func ReadSettings() {
	// Check if the settings file exists.
	if _, err := os.Stat(ExecPath + "/settings.preferences"); os.IsNotExist(err) {
		f, _ := os.Create(ExecPath + "/settings.preferences")
		b := bufio.NewWriter(f)
		// If the file doesn't exist write the default settings.
		b.WriteString(DefaultSettings)
		b.Flush()
		f.Close()
	}

	// Open the settings file.
	readFile, err := os.Open(ExecPath + "/settings.preferences")
	if err != nil {
		panic("failed to open file: " + err.Error())
	}

	// Open a filescanner to scan the settings file for the settings.
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var count int = 1
	// Begin scanning the file.
	for fileScanner.Scan() {
		value := fileScanner.Text()
		// Checking to see if the line is blank
		if value != "" {
			// Checking to see if the line inside the settings folder is commented out.
			if value[0:2] != "//" {
				// each count represents a setting. The settings must be in the same order as they are here.
				if count == 1 {
					// stores the setting into the variable only saves the setting and not the define inside the settings file. (the bit after the equals sign)
					ProgramName = value[14:len(value)]
					count = count + 1
				} else if count == 2 {
					SitePort = value[12:len(value)]
					count = count + 1
				} else if count == 3 {
					NonHttpsPort = value[11:len(value)]
					count = count + 1
				} else if count == 4 {
					// this settings converts the string value to boolean as the value for open browser can only be true or false.
					openBrowserOnLoad, _ = strconv.ParseBool(value[13:len(value)])
					count = count + 1
				}
			}
		}
	} // Close the settings file.
	readFile.Close()
}
