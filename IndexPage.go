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

type PageStruct struct {
	Data        template.HTML
	ProjectName string
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	// This function is called when there is a get request for the index page.
	// GetCategories - gets all the categories in the categories directory
	GetCategories()
	// Variable to store the data to be injected into the page
	var PageData string
	// Prepares the html to be injected into the page.
	if len(Categories) != 0 {
		for i := 0; i <= len(Categories)-1; i++ {
			var count int = 0
			// Opens the first csv file in the
			csvfile, err := os.Open(ExecPath + "/data/" + Categories[i] + ".csv")
			if err != nil {
				log.Fatalln("Couldn't open the csv file", err)
			}
			// parse the csv so we can read each line.
			r := csv.NewReader(csvfile)
			// for each line add one to the count variable
			for {
				_, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				count = count + 1
			}
			// close the csv file and add the info to the page data variable.
			csvfile.Close()
			PageData = PageData + ItemView("/"+Categories[i], Categories[i], "Amount of items: "+strconv.Itoa(count))
		}
	} else {
		PageData = ItemView("/", "No Cats Found", "I Cannot find any cats. :(")
	}

	// Send the parsed html data to the user.
	p := PageStruct{Data: template.HTML(PageData), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}

func IndexPagePost(w http.ResponseWriter, r *http.Request) {
	// This function is called when there is a post request.
	if r.FormValue("search") != "" {
		// Call the search and pass the values.
		SearchForItem(r.FormValue("search"), "all")
	} else {

	}
}

// Simple Function that can be used to format the items/GetCategories
func ItemView(link string, name string, details string) string {
	return "<li class=\"folders\"><a href=\"" + link + "\" title=\"files/\"" + name + "\" class=\"folders\"><span class=\"icon folder full\"></span><span class=\"name\">" + name + "</span><span class=\"details\">" + details + "</span></a></li>"
}
