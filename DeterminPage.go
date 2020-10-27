package main

import (
	"net/http"
	"strings"

	JG "github.com/joshuag1000/GoEssentials"
)

func DeterminPage(w http.ResponseWriter, r *http.Request) {
	// if something is posted to the server it will be handled here.
	if r.Method == http.MethodPost {
		IndexPagePost(w, r)
		// if this if statement was run then do not run the code below.
		return
	}
	if r.FormValue("search") != "" {
		// sometimes the search value is passed without a post so this needs to be checked to see if what was the case.
		IndexPagePost(w, r)
		return
	}
	// if nothing was posted display the index page. "r.URL.Path" is used to determin the url path
	// Removes the / at the beginning of the string to help with parsing.
	URL := strings.Replace(r.URL.Path, "/", "", 1)
	// checks to see if there is a / at the end of the string and if there is removes it.
	// doing this ensures the program gets the correct amount of splits when the string is parsed.
	if len(URL) >= 1 && URL[len(URL)-1] == '/' {
		URL = JG.Reverse(strings.Replace(JG.Reverse(URL), "/", "", 1))
	}
	UrlPaths := strings.Split(URL, "/")

	// Check to see if the URL is on / (the root directory)
	if URL == "" {
		// Load the Index page showing all the categories.
		IndexPage(w, r)
		return // Stop any other code running
	} else if len(UrlPaths) == 1 {
		// Call the load category function (we can pass through the category in the url to make life easier as well)
		LoadCategory(w, r, UrlPaths[0])
		return // Stop any other code running
	} else if len(UrlPaths) == 2 {
		// Load the item info of the item in the category. This can also be done using another function.
		LoadItem(w, r, UrlPaths[0], UrlPaths[1])
		return // Stop any other code running
	} else {
		// If the if statement ends up to here the url is not a valid format hence return a 404.
		http.NotFound(w, r)
		return // Stop any other code running
	}
}
