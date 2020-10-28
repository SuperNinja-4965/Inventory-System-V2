package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	JG "github.com/joshuag1000/GoEssentials"
)

func EditOrAsk(w http.ResponseWriter, r *http.Request) {
	// if nothing was posted display the index page. "r.URL.Path" is used to determin the url path
	// Removes the / at the beginning of the string to help with parsing.
	URL := strings.Replace(r.URL.Path, "/EditCategory/", "", 1)
	// checks to see if there is a / at the end of the string and if there is removes it.
	// doing this ensures the program gets the correct amount of splits when the string is parsed.
	if len(URL) >= 1 && URL[len(URL)-1] == '/' {
		URL = JG.Reverse(strings.Replace(JG.Reverse(URL), "/", "", 1))
	}
	UrlPaths := strings.Split(URL, "/")

	// Check to see if the URL is on / (the root directory)
	if URL == "" {
		// Ask which category to edit.
		ChooseCatToEdit(w, r)
		return // Stop any other code running
	} else if len(UrlPaths) == 1 {
		// Go to edit that Category
		EditCategory(w, r, UrlPaths[0])
		return // Stop any other code running
	} else {
		// If the if statement ends up to here the url is not a valid format hence return a 404.
		http.NotFound(w, r)
		return // Stop any other code running
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request, Category string) {
	// This is run when a category is selected to be edited.
	// If the cateory does not exist redirect the user to the select category page
	if _, err := os.Stat(ExecPath + "/data/" + Category + ".csv"); os.IsNotExist(err) {
		fmt.Println("There was a request to modify a non existant category.")
		http.Redirect(w, r, "/EditCategory/", 303)
		return
	} else {
		// if the category does exist then load the data so it can be edited.
		// check that this isn't a post request because if it is then data needs to be saved.
		var Report string = ""
		if r.Method == http.MethodPost {
			Report = EditCategoryPost(w, r, Category)
			if Report == "#CATDELETED#" {
				return
			}
		}
		var OutputData string = ""
		// open the file and load the data into the table.
		csvfile, err := os.Open(ExecPath + "/data/" + Category + ".csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		// Parse the file
		r := csv.NewReader(csvfile)
		var count int = 0
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
			// Form the html for the table.
			var RowDelTemplate string = "<tr><td><input type=\"checkbox\" id=\"" + strconv.Itoa(count) + "-0\" name=\"" + strconv.Itoa(count) + "-0\" value=\"yes\" disabled>"
			var RowTemplate1 string = "</td><td><input disabled id=\"" + record[0] + "-ID\" name=\"" + record[0] + "-ID\" size=\"4%\" style=\"background-color:transparent;color:white;border:0;\" value=\"" + record[0]
			var RowTemplate2 string = "\"></td><td><input id=\"" + record[0] + "-Name\" name=\"" + record[0] + "-Name\" size=\"10%\" style=\"background-color:transparent;color:white;border:0;\" value=\"" + record[1]
			var RowTemplate3 string = "\"></td><td><input id=\"" + record[0] + "-Value\" name=\"" + record[0] + "-Value\" size=\"6%\" style=\"background-color:transparent;color:white;border:0;\" value=\"" + record[2]
			var RowTemplate4 string = "\"></td><td><input id=\"" + record[0] + "-AmountAvailable\" name=\"" + record[0] + "-AmountAvailable\" size=\"6%\" style=\"background-color:transparent;color:white;border:0;\" value=\"" + record[3]
			var RowTemplate5 string = "\"></td><td><input id=\"" + record[0] + "-AmountInUse\" name=\"" + record[0] + "-AmountInUse\" size=\"6%\" style=\"background-color:transparent;color:white;border:0;\" value=\"" + record[4]
			var RowTemplate6 string = "\"></td><td><input id=\"" + record[0] + "-Note\" name=\"" + record[0] + "-Note\" size=\"66%\" style=\"background-color:transparent;color:white;border:0;\" value=\"" + record[5]
			var RowTemplateEnd string = "\"></td></tr>"
			// conbined all the variables into one.
			OutputData = OutputData + RowDelTemplate + RowTemplate1 + RowTemplate2 + RowTemplate3 + RowTemplate4 + RowTemplate5 + RowTemplate6 + RowTemplateEnd
			count = count + 1
		}
		// Closes the CSV File
		csvfile.Close()

		// Sending the data to the user
		// Compiles the first and last part of the html including the table head.
		var templatePart1 string = "<script>function EnableBoxes() {  var inputs = document.getElementsByTagName(\"input\");for(var i = 0; i < inputs.length; i++) {    if(inputs[i].type == \"checkbox\") {        inputs[i].disabled = false;     }  }} function ConfirmDelete() {  if (confirm(\"Are you sure you want to delete this category?\")) {  	document.getElementById(\"DeleteCat\").setAttribute(\"value\", \"Yes\");    document.getElementById(\"CatForm\").submit();  } else {  	document.getElementById(\"DeleteCat\").setAttribute(\"value\", \"No\"); }}</script><center><h1 style=\"color:white;\">Editing Category: " + Category + "</h1><div class=\"container\"><br><form id=\"CatForm\" method=\"POST\"><h2 style=\"color:white;\">Record Count: " + strconv.Itoa(count) + "</h2><br>" + Report + "<input type=\"hidden\" id=\"Count\" name=\"Count\" value=\"" + strconv.Itoa(count) + "\"><table><thead><tr><th>Del?</th><th>ID</th><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Notes</th></tr></thead><tbody>"
		var templatePartEnd string = "</tbody></table><br><button name=\"AddRow\" type=\"submit\" value=\"Yes\">Add Item</button>  <button type=\"button\" id=\"EnableBoxesButton\" onclick=\"EnableBoxes()\">Enable Del Checkboxes</button><br><br><button style=\"color:red;\" type=\"submit\" id=\"DeleteCat\" name=\"DeleteCat\" onclick=\"ConfirmDelete()\" value=\"No\">Delete Category</button><br><br><input type=\"submit\" value=\"Save Changes\"></form></div></center>"
		// sends the data to the user.
		p := PageStruct{Data: template.HTML(templatePart1 + OutputData + templatePartEnd), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
	}
}

func EditCategoryPost(w http.ResponseWriter, r *http.Request, Category string) string {
	// if the Delete category is set to yes then delete the category and return the result.
	if r.FormValue("DeleteCat") == "Yes" {
		var err = os.Remove(ExecPath + "/data/" + Category + ".csv")
		var response string
		// checks to see if the deletion was successful.
		if err != nil {
			fmt.Println(err)
			response = "<center><h1 style=\"color:red;\">There was an error deleting that file.</h1></center>"
		} else {
			response = "<center><h1 style=\"color:red;\">Category " + Category + " Deleted.</h1></center><meta http-equiv=\"refresh\" content=\"1;url=/EditCategory/\" />"
		}
		// send the response to the user.
		p := PageStruct{Data: template.HTML(response), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
		return "#CATDELETED#"
	}
	// if the DeleteCat form value isn't set to yes then proceed to save the data as requested.

	return ""
}

func ChooseCatToEdit(w http.ResponseWriter, r *http.Request) {
	// when the data is posted to the page (the Category to load) this block will redirect the user.
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/EditCategory/"+r.FormValue("SelectedCategory"), 303)
		return
	}
	// These variables define the beginnign and end of the html block that will be injected.
	var PreSelectHTML string = "<center><h1 style=\"color:white;\">Select a Category</h1><div class=\"container\"><br><form method=\"POST\"><table><thead><tr><th><h2 style=\"color:white;\">Type</h2></th><th><h2 style=\"color:white;\">Value</h2></th></tr></thead><tbody><tr><td><h3 style=\"color:white;\">Category:</h3></td><td><select id=\"SelectedCategory\" name=\"SelectedCategory\">"
	var PostSelectHTML string = "</select></td></tr></tbody></table><br><input type=\"submit\" value=\"Select Category\"> </form></div></center>"
	var MiddleSelectHTML string = ""
	// Refreshed the categories array.
	GetCategories()
	// loops through all the categories and adds them to the option list for the user to select from.
	if len(Categories) != 0 {
		for i := 0; i <= len(Categories)-1; i++ {
			MiddleSelectHTML = MiddleSelectHTML + " <option value=\"" + Categories[i] + "\">" + Categories[i] + "</option>"
		}
	} else {
		// if there aren't any categories then set the middle section to mothing.
		MiddleSelectHTML = ""
	}
	// Sends the parsed data (beginning, middle and end) to the user.
	p := PageStruct{Data: template.HTML(PreSelectHTML + MiddleSelectHTML + PostSelectHTML), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}
