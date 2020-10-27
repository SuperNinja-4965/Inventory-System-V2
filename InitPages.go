package main

import (
	"fmt"
	"net/http"
)

// This function will initialise the pages needed for the web server.
func InitPages() {
	// define the css pages - these are static pages that are needed for the web page to load properly.
	http.HandleFunc("/assets/css/styles.css", stylesCss)
	http.HandleFunc("/assets/css/styles2.css", styles2Css)
}

func stylesCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, cssIndex)
}

func styles2Css(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, cssTwo)
}
