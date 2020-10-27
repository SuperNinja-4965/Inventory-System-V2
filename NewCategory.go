package main

import (
	"bufio"
	"errors"
	"html/template"
	"net/http"
	"os"
)

func NewCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// If this is a post request then create the category
		err := CreateCategory(r.FormValue("CatName"))
		if err != nil {
			p := PageStruct{Data: template.HTML("<center> <h1 style=\"color:red;\">" + err.Error() + "</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
			t, _ := template.New("indexTemplate").Parse(PageIndex)
			t.Execute(w, p)
			return
		}
		p := PageStruct{Data: template.HTML("<center> <h1 style=\"color:green;\">The Category was created successfully!</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
		return // stops any further code from running
	}
	// Here we present a simple form asking which category to create.
	p := PageStruct{Data: template.HTML(NewCatForm), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}

func CreateCategory(name string) error {
	f, err := os.Create(ExecPath + "/data/" + name + ".csv")
	if err != nil {
		return errors.New("There was an error creating the category.")
	}
	b := bufio.NewWriter(f)
	b.WriteString("\"1\",\"item1\",\"100f\",100,10,\"This is a cool item, and it always will be.\"")
	b.Flush()
	f.Close()
	return nil
}
