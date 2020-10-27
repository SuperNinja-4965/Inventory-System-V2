package main

import (
	"fmt"
	"net/http"
)

func LoadItem(w http.ResponseWriter, r *http.Request, Category string, ItemID string) {
	fmt.Println(Category)
	fmt.Println(ItemID)
}
