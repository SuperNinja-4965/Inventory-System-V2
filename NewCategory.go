package main

import (
	"bufio"
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"

	JG "github.com/joshuag1000/GoEssentials"
)

func CreateByUrlOrForm(w http.ResponseWriter, r *http.Request) {
	// if nothing was posted display the index page. "r.URL.Path" is used to determin the url path
	// Removes the / at the beginning of the string to help with parsing.
	URL := strings.Replace(r.URL.Path, "/NewCategory/", "", 1)
	// checks to see if there is a / at the end of the string and if there is removes it.
	// doing this ensures the program gets the correct amount of splits when the string is parsed.
	if len(URL) >= 1 && URL[len(URL)-1] == '/' {
		URL = JG.Reverse(strings.Replace(JG.Reverse(URL), "/", "", 1))
	}
	UrlPaths := strings.Split(URL, "/")

	// Check to see if the URL is on / (the root directory)
	if URL == "" {
		// Ask which category to Create.
		NewCategory(w, r)
		return // Stop any other code running
	} else if len(UrlPaths) == 1 {
		// Create the category specified.
		// If this is a post request then create the category
		err := CreateCategory(UrlPaths[0])
		if err != nil {
			// If there was an error resend the form and an error message
			p := PageStruct{Data: template.HTML("<center> <h1 style=\"color:red;\">" + err.Error() + "</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
			t, _ := template.New("indexTemplate").Parse(PageIndex)
			t.Execute(w, p)
			return
		}
		// Display a modified form that has the category name filled into the box.
		NewCatFormV2 := "<center><h1 style=\"color:white;\">New Category</h1>                <div class=\"container\">                    <br>                    <form method=\"POST\">                        <table>                            <thead>                                <tr>                                    <th>                                        <h2 style=\"color:white;\">Type</h2></th>                                    <th>                                        <h2 style=\"color:white;\">Value</h2></th>                                </tr>                            </thead>                            <tbody>                                <tr>                                    <td>                                        <h3 style=\"color:white;\">Category Name:</h3></td>                                    <td>                                        <input type=\"textbox\" name=\"CatName\" id=\"CatName\" placeholder=\"Category Name\" value=\"" + UrlPaths[0] + "\" style=\"display: inline-block;\">                                    </td>                                </tr>                            </tbody>                        </table>                        <br>                        <input type=\"submit\" value=\"Create Category\"> </form> 						</div><br><h2 style=\"color:white;\"><a href=\"javascript:history.back()\">Back</a></h2></center>"
		// if the category was successfully created then tell the user that and then redirect after three seconds to the edit category.
		p := PageStruct{Data: template.HTML("<center> <h1 style=\"color:green;\">The Category was created successfully!</h1><br><h3 style=\"color:white;\">You will be redirected to edit this category in 3 seconds.</h3><meta http-equiv=\"refresh\" content=\"3;url=/EditCategory/" + UrlPaths[0] + "/\" /> </center> <br>" + NewCatFormV2), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
		return // stops any further code from running
	} else {
		// If the if statement ends up to here the url is not a valid format hence return a 404.
		http.NotFound(w, r)
		return // Stop any other code running
	}
}

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
