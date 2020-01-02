package main

import (
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

func processMain(w http.ResponseWriter, _ *http.Request) {
	bsInstance := bslib.GetInstance()
	items, err := bsInstance.ReadAllItems()

	if err != nil {
		log.Println(err.Error())
		return
	}

	err = renderHTMLTemplate(w, "items", items)
	if err != nil {
		ErrorPage(w, "Error", err.Error())
	}
}
