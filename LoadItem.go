package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func LoadItem(w http.ResponseWriter, r *http.Request, Category string, ItemID string) {
	// Search for the item and store the results into a variable.
	ItemSearchresults, ItemName := SearchForItemInCategory(Category, ItemID)

	// Variable to store the data to be injected into the page
	var PageData string = "<center><h1 style=\"color:white;\">" + Category + ": " + ItemName + " - Information</h1><div class=\"container\"><br><table><thead><tr><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Total amount</th><th>Notes</th></tr></thead><tbody>" + ItemSearchresults + "</tbody></table></div><br><h2 style=\"color:white;\"><a href=\"javascript:history.back()\">Back</a></h2></center>"

	// Send the parsed html data to the user.
	p := PageStruct{Data: template.HTML(PageData), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}

func SearchForItemInCategory(Category string, IDToSearch string) (ReturnData string, ItemName string) {
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
		ItemName = record[1]
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

		// Check to see if the item is the item we are looking for.
		if IDToSearch == ItemID {
			ReturnData = "<tr><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + ItemAmountAvailable + "</td><td>" + ItemInUse + "</td><td>" + strconv.Itoa(ItemAmountTotal) + "</td><td>" + ItemNotes + "</td></tr>"
			csvfile.Close()
			return
		}
	}
	csvfile.Close()
	if ReturnData == "" {
		ReturnData = "<tr><td>No Items</td><td>ERR</td><td>ERR</td><td>ERR</td><td>ERR</td><td>ERR</td></tr>"
		ItemName = "No Item Found"
	}
	return
}
