package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func LoadCategory(w http.ResponseWriter, r *http.Request, Category string) {
	// This Function is called when the information of a category is requested.
	// GetCategories - gets all the categories in the categories directory
	GetCategories()
	// Variable to store the data to be injected into the page
	var PageData string = "<center><h1 style=\"color:white;\">" + Category + " - Information</h1><div class=\"container\"><br><table><thead><tr><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Total amount</th><th>Notes</th></tr></thead><tbody>" + GetItemsInCategory(Category) + "</tbody></table></div><br><h2 style=\"color:white;\"><a href=\"javascript:history.back()\">Back</a></h2></center>"

	// Send the parsed html data to the user.
	p := PageStruct{Data: template.HTML(PageData), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}

// Categories - where all the categories are stored.
var Categories []string

// GetCategories - gets all of the Categories in /data
func GetCategories() {
	// deletes any data in the array
	Categories = nil
	count := 0
	// walks through the entire data directory and gets all of the categories.
	err := filepath.Walk(ExecPath+"/data/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			count = count + 1
			//fmt.Println(count)
			// depending on the os a different format of path may be needed. I am not sure why. This switch removes the path from the files in the directory eg result will look like test.csv
			// when this is first run there is no file (the path of the directory that is being walked is returned) so to stop that being stored the first run is ignored.
			if count > 1 {
				switch runtime.GOOS {
				case "linux":
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				case "windows":
					path = strings.ReplaceAll(path, ExecPath+"\\data\\", "")
				case "darwin":
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				default:
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				}
				// the .csv ending is then removed from the file name.
				path = strings.ReplaceAll(path, ".csv", "")
				// checks to ensure that there is actually a path left.
				if path == "" {
				} else if path == "\n" {
				} else {
					// adds the result to the category array.
					Categories = append(Categories, path)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	} else {
		//fmt.Println("Loaded Categories")
	}
}

func GetItemsInCategory(Category string) (ReturnData string) {
	// Open the csv file
	csvfile, err := os.Open(ExecPath + "/data/" + Category + ".csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// store the items data into variables
		var ItemID string = record[0]
		var ItemName string = record[1]
		var ItemValue string = record[2]
		var ItemAmountAvailable string = record[3]
		if ItemAmountAvailable == "" {
			ItemAmountAvailable = "0"
		}
		var ItemInUse string = record[4]
		if ItemInUse == "" {
			ItemInUse = "0"
		}
		var ItemNotes string = record[5]

		// Converting the Item amounts to Int and adding them
		ItemAmountAvailableInt, _ := strconv.Atoi(ItemAmountAvailable)
		ItemInUseInt, _ := strconv.Atoi(ItemInUse)
		ItemAmountTotal := ItemAmountAvailableInt + ItemInUseInt

		ReturnData = ReturnData + "<tr onclick=\"window.location='/" + Category + "/" + ItemID + "';\"><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + ItemAmountAvailable + "</td><td>" + ItemInUse + "</td><td>" + strconv.Itoa(ItemAmountTotal) + "</td><td>" + ItemNotes + "</td></tr>"
	}
	csvfile.Close()
	if ReturnData == "" {
		ReturnData = "<tr><td>No Items</td><td>ERR</td><td>ERR</td><td>ERR</td><td>ERR</td><td>ERR</td></tr>"
	}
	return
}
