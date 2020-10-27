package main

import (
	"fmt"
	"net/http"
)

// This function will initialise the pages needed for the web server.
func InitPages() {
	// Define the index page that will be shown
	http.HandleFunc("/", DeterminPage)
	// define the css pages - these are static pages that are needed for the web page to load properly.
	http.HandleFunc("/assets/css/styles.css", stylesCss)
	http.HandleFunc("/assets/css/styles2.css", styles2Css)
	// handle a favicon request with a 404 cannot be found.
	http.HandleFunc("/favicon.ico", favicon)
}

// when called a css header is added to the page and then the css code is returned
func stylesCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, cssIndex)
}

// when called a css header is added to the page and then the css code is returned
func styles2Css(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, cssTwo)
}

// returns a 404 error when this is called - this can be used to return an icon in the future.
func favicon(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
