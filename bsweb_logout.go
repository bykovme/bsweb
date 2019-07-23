package main

import (
	"net/http"

	"github.com/bykovme/bslib"
)

func processLogout(w http.ResponseWriter, r *http.Request) {
	bsInstance := bykovstorage.GetInstance()
	bsInstance.Lock()
	http.Redirect(w, r, "/", 302)
}
