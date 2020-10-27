package main

import (
	"fmt"
	"net/http"
)

func LoadCategory(w http.ResponseWriter, r *http.Request, Category string) {
	fmt.Println(Category)
}
