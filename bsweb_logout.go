package main

import (
	"net/http"

	"github.com/bykovme/bslib"
)

func processLogout(w http.ResponseWriter, r *http.Request) {
	bsInstance := bslib.GetInstance()
	if err := bsInstance.Lock(); err != nil {
		ErrorPage(w, "Error", err.Error())
		return
	}
	http.Redirect(w, r, "/", 302)
}
