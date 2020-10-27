package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func LoadCategory(w http.ResponseWriter, r *http.Request, Category string) {
	fmt.Println(Category)

}

// Categories - where all the categories are stored.
var Categories []string

// GetCategories - gets all of the Categories in /data
func GetCategories() {
	// deletes any data in the array
	Categories = nil
	count := 0
	// walks through the entire data directory and gets all of the categories.
	err := filepath.Walk(ExecPath+"/data/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			count = count + 1
			//fmt.Println(count)
			// depending on the os a different format of path may be needed. I am not sure why. This switch removes the path from the files in the directory eg result will look like test.csv
			// when this is first run there is no file (the path of the directory that is being walked is returned) so to stop that being stored the first run is ignored.
			if count > 1 {
				switch runtime.GOOS {
				case "linux":
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				case "windows":
					path = strings.ReplaceAll(path, ExecPath+"\\data\\", "")
				case "darwin":
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				default:
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				}
				// the .csv ending is then removed from the file name.
				path = strings.ReplaceAll(path, ".csv", "")
				// checks to ensure that there is actually a path left.
				if path == "" {
				} else if path == "\n" {
				} else {
					// adds the result to the category array.
					Categories = append(Categories, path)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	} else {
		//fmt.Println("Loaded Categories")
	}
}
