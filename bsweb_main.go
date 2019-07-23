package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

func processMain(w http.ResponseWriter, r *http.Request) {
	bsInstance := bykovstorage.GetInstance()
	items, err := bsInstance.ReadAllItems()

	if err != nil {
		log.Println(err.Error())
		return
	}

	t, _ := template.ParseFiles("templates/items.html")
	t.Execute(w, items)
}
