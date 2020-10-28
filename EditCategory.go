package main

import (
	"html/template"
	"net/http"
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
	// This is run when a category is selected to be exited.
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
