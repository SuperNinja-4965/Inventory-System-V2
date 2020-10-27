package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// StartUp - This is run when the program begins. Checks to see if files and directories exist. Useful for deployment.
func StartUp() {
	// Calling this function will load all of the css and the html page into variables.
	FileStore()
	// checks to see if the data ditectory exists and if it doesnt create it with an example category with item inside.
	if _, err := os.Stat(ExecPath + "/data"); os.IsNotExist(err) {
		fmt.Printf("/data directory created.\n")
		err := os.Mkdir(ExecPath+"/data", 0755)
		check(err)
		// Creates an example category
		f, err := os.Create(ExecPath + "/data/Example.csv")
		b := bufio.NewWriter(f)
		b.WriteString("\"1\",\"item1\",\"100f\",100,10,\"This is a cool item, and it always will be.\"")
		b.Flush()
		f.Close()
	}
	// creates the https folder where the user can add the files needed for https.
	if _, err := os.Stat(ExecPath + "/HTTPS-key"); os.IsNotExist(err) {
		fmt.Printf("/HTTPS-key directory created.\n")
		err := os.Mkdir(ExecPath+"/HTTPS-key", 0755)
		check(err)
		f, err := os.Create(ExecPath + "/HTTPS-key/README.txt")
		b := bufio.NewWriter(f)
		b.WriteString("If you want to use HTTPS with this program then you will need to insert 2 files into this directory: A Server Certificate and a Server Private key.\n\nThe file which contains the private key should be called: server.key\nAnd the file containing the server certificate should be called: server.crt\n\nIf you do not want to use HTTPS then leave this directory empty.")
		b.Flush()
		f.Close()
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
