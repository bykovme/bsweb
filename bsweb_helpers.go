package main

import (
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

func getAction(r *http.Request) string {
	r.ParseForm()
	action, err := getValueByName(r, "action")
	if err != nil {
		return ""
	}
	return action
}

func checkLockAndStandardActions(w http.ResponseWriter, r *http.Request) bool {

	action := getAction(r)
	bsInstance := bykovstorage.GetInstance()
	if action == "logout" {
		processLogout(w, r)
	} else if bsInstance.IsNew() == true {
		log.Println("New storage, start sign up process")
		processRegister(w, r)
	} else if bsInstance.IsLocked() == true {
		log.Println("Storage is locked, show login page")
		processLogin(w, r)
	} else {
		return false
	}
	return true
}
