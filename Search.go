package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	JG "github.com/joshuag1000/GoEssentials"
)

func SearchOrAsk(w http.ResponseWriter, r *http.Request) {
	// if nothing was posted display the index page. "r.URL.Path" is used to determin the url path
	// Removes the / at the beginning of the string to help with parsing.
	URL := strings.Replace(r.URL.Path, "/Search/", "", 1)
	// checks to see if there is a / at the end of the string and if there is removes it.
	// doing this ensures the program gets the correct amount of splits when the string is parsed.
	if len(URL) >= 1 && URL[len(URL)-1] == '/' {
		URL = JG.Reverse(strings.Replace(JG.Reverse(URL), "/", "", 1))
	}
	UrlPaths := strings.Split(URL, "/")

	// Check to see if the URL is on / (the root directory)
	if URL == "" {
		// Ask what to search and where.
		DisplayAdvancedSearch(w, r, "")
		return // Stop any other code running
	} else if len(UrlPaths) == 2 {
		// Search for the item in the category specified.
		SearchForItem(w, r, UrlPaths[0], UrlPaths[1])
		return // Stop any other code running
	} else {
		// If the if statement ends up to here the url is not a valid format hence return a 404.
		DisplayAdvancedSearch(w, r, "There was an error handling your request. <br>Are you sure you did the search correctly?")
		return // Stop any other code running
	}
}

func SearchForItem(w http.ResponseWriter, r *http.Request, Category string, ItemToFind string) {
	if Category == "all" {
		SearchForItemEverywhere(w, r, ItemToFind)
		return
	}
	// Define the Variable to store the search results as a html table
	var ResultPartial string
	var ResultMatch string
	var ReturnData string
	// Open the csv file
	csvfile, err := os.Open(ExecPath + "/data/" + Category + ".csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	CSVr := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := CSVr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// store the items data into variables
		//var ItemID string = record[0]
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

		// Check to see if the item is the item we are looking for. Or if it is a partial match
		if strings.ToLower(ItemName) == strings.ToLower(ItemToFind) {
			ResultMatch = ResultMatch + "<tr><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + ItemAmountAvailable + "</td><td>" + ItemInUse + "</td><td>" + strconv.Itoa(ItemAmountTotal) + "</td><td>" + ItemNotes + "</td></tr>"
		} else if strings.Contains(strings.ToLower(ItemName), strings.ToLower(ItemToFind)) {
			ResultPartial = ResultPartial + "<tr><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + ItemAmountAvailable + "</td><td>" + ItemInUse + "</td><td>" + strconv.Itoa(ItemAmountTotal) + "</td><td>" + ItemNotes + "</td></tr>"
		}
	}
	csvfile.Close()
	ReturnData = ResultMatch + ResultPartial
	if ReturnData == "" {
		ReturnData = "<tr><td>No Items found in the search</td><td>N/A</td><td>N/A</td><td>N/A</td><td>N/A</td><td>N/A</td></tr>"
	}

	// Variable to store the data to be injected into the page
	var PageData string = "<center><h1 style=\"color:white;\">" + Category + ": " + ItemToFind + " - Search Results</h1><div class=\"container\"><br><table><thead><tr><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Total amount</th><th>Notes</th></tr></thead><tbody>" + ReturnData + "</tbody></table></div><br><h2 style=\"color:white;\"><a href=\"javascript:history.back()\">Back</a></h2></center>"

	// Send the parsed html data to the user.
	p := PageStruct{Data: template.HTML(PageData), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}

func SearchForItemEverywhere(w http.ResponseWriter, r *http.Request, ItemToFind string) {
	var PartailResultsAll string = ""
	var ResultsAll string = ""
	GetCategories()

	if len(Categories) != 0 {
		for i := 0; i <= len(Categories)-1; i++ {
			//Table Formatting
			var TableBegin string = "<h2 style=\"color:white;\">Category: " + Categories[i] + "</h2><div class=\"container\"><br><table><thead><tr><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Total amount</th><th>Notes</th></tr></thead><tbody>"
			var TableEnd string = "</tbody></table></div>"
			// Define the Variable to store the search results as a html table
			var ResultPartial string = TableBegin
			var ResultMatch string = TableBegin
			// Open the csv file
			csvfile, err := os.Open(ExecPath + "/data/" + Categories[i] + ".csv")
			if err != nil {
				log.Fatalln("Couldn't open the csv file", err)
			}
			// Parse the file
			CSVr := csv.NewReader(csvfile)
			// Iterate through the records
			for {
				// Read each record from csv
				record, err := CSVr.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				// store the items data into variables
				//var ItemID string = record[0]
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

				// Check to see if the item is the item we are looking for. Or if it is a partial match
				if strings.ToLower(ItemName) == strings.ToLower(ItemToFind) {
					ResultMatch = ResultMatch + "<tr><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + ItemAmountAvailable + "</td><td>" + ItemInUse + "</td><td>" + strconv.Itoa(ItemAmountTotal) + "</td><td>" + ItemNotes + "</td></tr>"
				} else if strings.Contains(strings.ToLower(ItemName), strings.ToLower(ItemToFind)) {
					ResultPartial = ResultPartial + "<tr><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + ItemAmountAvailable + "</td><td>" + ItemInUse + "</td><td>" + strconv.Itoa(ItemAmountTotal) + "</td><td>" + ItemNotes + "</td></tr>"
				}
			}
			csvfile.Close()
			if ResultMatch != TableBegin {
				ResultsAll = ResultsAll + ResultMatch + TableEnd + "<br><br>"
			}
			if ResultPartial != TableBegin {
				PartailResultsAll = PartailResultsAll + ResultPartial + TableEnd + "<br><br>"
			}
		}
	}

	// Variable to store the data to be injected into the page
	var BeginningPageData string = "<center><h1 style=\"color:white;\">Searching everywhere for: " + ItemToFind + "</h1>"
	var BeginningPageDataPartialResults string = "<center><h1 style=\"color:white;\">Partial Results for: " + ItemToFind + "</h1>"
	var EndPageData string = "<br><h2 style=\"color:white;\"><a href=\"javascript:history.back()\">Back</a></h2></center>"

	// Send the parsed html data to the user.
	// This will not show the partial results or full results section if one is not needed. If no results are found then show the advanced search box and a message saying so.
	if ResultsAll != "" && PartailResultsAll != "" {
		p := PageStruct{Data: template.HTML(BeginningPageData + ResultsAll + EndPageData + "<br><br>" + BeginningPageDataPartialResults + PartailResultsAll + EndPageData), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
	} else if ResultsAll != "" {
		p := PageStruct{Data: template.HTML(BeginningPageData + ResultsAll + EndPageData), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
	} else if PartailResultsAll != "" {
		p := PageStruct{Data: template.HTML(BeginningPageDataPartialResults + PartailResultsAll + EndPageData), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
	} else {
		// Having the extra option on the extra option here allows us to easily show the no results message along with a search box.
		DisplayAdvancedSearch(w, r, "No results were found. :(")
	}
}

// Having the extra option on the extra option here allows us to easily show the no results message along with a search box.
func DisplayAdvancedSearch(w http.ResponseWriter, r *http.Request, ErrorMessage string) {
	var searchOptionsHTML string = "<script>function AdvancedSearchData() {if (document.getElementById(\"SearchFor\").value != \"\") {if (document.getElementById(\"CategoryToSeachIn\").value != \"\") {window.location.replace(window.location.origin + \"/Search/\" + document.getElementById(\"CategoryToSeachIn\").value + \"/\" + document.getElementById(\"SearchFor\").value);} else {document.getElementById(\"ErrorMessage\").innerHTML = \"In Box cannot be blank.\";}} else {document.getElementById(\"ErrorMessage\").innerHTML = \"Search Box cannot be blank.\";}}</script><center><h1 style=\"color:white;\">New Search</h1><br><h2 style=\"color:white;\">To search everywhere select \"all\" on the In Box.</h2><br><h1 id=\"ErrorMessage\" style=\"color:red;\">"
	var searchOptionsHTML2 string = "</h1><div class=\"container\"><br><form><table><thead><tr><th><h2 style=\"color:white;\">Type</h2></th><th><h2 style=\"color:white;\">Value</h2></th></tr></thead><tbody><tr><td><h3 style=\"color:white;\">Search For:</h3></td><td><input type=\"textbox\" style=\"display: inline-block;\" placeholder=\"Search For\" name=\"SearchFor\" id=\"SearchFor\"></td></tr><tr><td><h3 style=\"color:white;\">In:</h3></td><td><select name=\"CategoryToSeachIn\" id=\"CategoryToSeachIn\" style=\"display: inline-block;\">"
	var searchOptionsMiddle = "<option value=\"all\">All</option>"
	var searchOptionsEnd string = "</select></td></tr></tbody></table><br><h2 style=\"color:white;\"><a href=\"#\" onclick=\"AdvancedSearchData()\">Search</a></h2></form></div></center>"
	if len(Categories) != 0 {
		for i := 0; i <= len(Categories)-1; i++ {
			searchOptionsMiddle = searchOptionsMiddle + "<option value=\"" + Categories[i] + "\">" + Categories[i] + "</option>"
		}
	}

	// Send the parsed html data to the user.
	p := PageStruct{Data: template.HTML(searchOptionsHTML + ErrorMessage + searchOptionsHTML2 + searchOptionsMiddle + searchOptionsEnd), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}
