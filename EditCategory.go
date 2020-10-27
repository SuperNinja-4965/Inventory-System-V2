package main

import (
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
		EditCategory(w, r)
		return // Stop any other code running
	} else {
		// If the if statement ends up to here the url is not a valid format hence return a 404.
		http.NotFound(w, r)
		return // Stop any other code running
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {

}

func ChooseCatToEdit(w http.ResponseWriter, r *http.Request) {

}
