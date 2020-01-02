package main

import (
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

// RegisterPage - registration page structure
type RegisterPage struct {
	ErrorText string
	Cyphers   []string
}

func processRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("Register page started")

	if r.Method == "POST" {
		log.Println("RegisterPage: POST, setting new password")
		setMasterPassword(w, r)
	} else if r.Method == "GET" {
		log.Println("RegisterPage: GET, show password page")
		showRegisterPage(w, r)
	} else {
		// error unknown request
		log.Println("RegisterPage: Unknown request type")
	}
}

func setMasterPassword(w http.ResponseWriter, r *http.Request) {
	cypher, err := getValueByName(r, "cypher")
	if err != nil {
		http.Redirect(w, r, "/?err=ERR00004", 302)
		return
	}

	password, err := getValueByName(r, "password")
	if err != nil || password == "" {
		http.Redirect(w, r, "/?err=ERR00001", 302)
		return
	}
	passwordConf, err := getValueByName(r, "password_confirm")
	if err != nil || passwordConf == "" {
		http.Redirect(w, r, "/?err=ERR00002", 302)
		return
	}
	if password != passwordConf {
		http.Redirect(w, r, "/?err=ERR00003", 302)
		return
	}

	bsInstance := bslib.GetInstance()
	log.Printf("Sign up, password: %s, cypher: %s", password, cypher)
	err = bsInstance.SetNewPassword(password, cypher)
	if err != nil {
		log.Fatalln("Error initializing with master password:" + err.Error())
		return
	}

	log.Print("Set new password and start working")
	err = bsInstance.Unlock(password)
	if err != nil {
		log.Fatalln("Error to open with new password:" + err.Error())
		return
	}

	http.Redirect(w, r, "/", 302)
}

func showRegisterPage(w http.ResponseWriter, r *http.Request) {
	var rf RegisterPage

	log.Println("Get errors, if any")
	errMsg, err := getValueByName(r, "err")
	if err == nil {
		errText, isFound := bsErrors[errMsg]
		if isFound == false {
			errText = "Unknown error"
		}
		rf = RegisterPage{ErrorText: errText}
	}
	bsInstance := bslib.GetInstance()

	rf.Cyphers = bsInstance.GetAvailableCyphers()
	err = renderHTMLTemplate(w, "register", rf)
	if err != nil {
		ErrorPage(w, "Error", err.Error())
	}

}
