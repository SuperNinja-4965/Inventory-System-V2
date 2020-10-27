package main

import (
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	// This function is called when there is a get request for the index page.

}

func IndexPagePost(w http.ResponseWriter, r *http.Request) {
	// This function is called when there is a post request.
	if r.FormValue("search") != "" {
		// Call the search and pass the values.
		SearchForItem(r.FormValue("search"), "all")
	} else {

	}
}
