package main

import (
	"bufio"
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func NewCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// If this is a post request then create the category
		err := CreateCategory(r.FormValue("CatName"))
		if err != nil {
			// If there was an error resend the form and an error message
			p := PageStruct{Data: template.HTML("<center> <h1 style=\"color:red;\">" + err.Error() + "</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
			t, _ := template.New("indexTemplate").Parse(PageIndex)
			t.Execute(w, p)
			return
		}
		// if the category was successfully created then tell the user that and then redirect after three seconds to the edit category.
		p := PageStruct{Data: template.HTML("<center> <h1 style=\"color:green;\">The Category was created successfully!</h1><br><h3 style=\"color:white;\">You will be redirected to edit this category in 3 seconds.</h3><meta http-equiv=\"refresh\" content=\"3;url=/EditCategory/" + r.FormValue("CatName") + "/\" /> </center> <br>" + NewCatForm), ProjectName: ProgramName}
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
	// here we are performing simple checks to make sure an illegal category name isnt used.
	if name == "" {
		return errors.New("The Category name cannot be blank.")
	} else if strings.ToLower(name) == "newcategory" {
		return errors.New("You cannot name your category that.")
	} else if strings.ToLower(name) == "editcategory" {
		return errors.New("You cannot name your category that.")
	} else if strings.ToLower(name) == "favicon.ico" {
		return errors.New("You cannot name your category that.")
	} else {
		// checks to make sure that the category doesnt already exist.
		if _, err := os.Stat(ExecPath + "/data/" + name + ".csv"); os.IsNotExist(err) {
			f, err := os.Create(ExecPath + "/data/" + name + ".csv")
			if err != nil {
				// if there was an error creating the category say so here.
				return errors.New("There was an error creating the category.")
			}
			b := bufio.NewWriter(f)
			// write a simple example into the file here.
			b.WriteString("\"1\",\"item1\",\"100f\",100,10,\"This is a cool item, and it always will be.\"")
			b.Flush()
			f.Close()
			// return nil if the category was created successfully/
			return nil
		} else {
			// throw an error if the category already exists.
			return errors.New("That category already exists.")
		}
	}
}
