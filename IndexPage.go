package main

import (
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	// if something is posted to the server it will be handled here.
	if r.Method == http.MethodPost {

		// if this if statement was run then do not run the code below.
		return
	}
	// if nothing was posted display the index page.

}
