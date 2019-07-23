package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

// LoginPage - registration page structure
type LoginPage struct {
	ErrorText string
}

func processLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		checkPassword(w, r)
	} else if r.Method == "GET" {
		showLoginPage(w, r)
	} else {
		// error unknown request
	}
}

func checkPassword(w http.ResponseWriter, r *http.Request) {
	password, err := getValueByName(r, "password")
	if err != nil || password == "" {
		http.Redirect(w, r, "/?err=ERR00001", 302)
		return
	}
	bsInstance := bykovstorage.GetInstance()
	err = bsInstance.SetPassword(password)
	if err != nil && err.Error() == bykovstorage.BSERR00010EncWrongPassword {
		http.Redirect(w, r, "/?err=ERR00005", 302)
		return
	}

	if err != nil {
		log.Print(err.Error())
		http.Redirect(w, r, "/?err=ERR00005", 302)
		return
	}

	http.Redirect(w, r, "/", 302)

}

func showLoginPage(w http.ResponseWriter, r *http.Request) {
	var lp LoginPage
	t, _ := template.ParseFiles("templates/login.html")
	errMsg, err := getValueByName(r, "err")
	if err == nil {
		errText, isFound := bsErrors[errMsg]
		if isFound == false {
			errText = "Unknown error"
		}
		lp = LoginPage{ErrorText: errText}
	}
	//bsInstance := bykovstorage.GetInstance()
	t.Execute(w, lp)
}
