package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

func processAddItem(w http.ResponseWriter, r *http.Request) {
	action := getAction(r)
	switch action {
	case "additem":
		addItem(w, r)
		break
	default:
		t, _ := template.ParseFiles("templates/additem.html")
		t.Execute(w, nil)
		break
	}

}

func addItem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	itemName, err01 := getValueByName(r, "item_name")
	if err01 != nil || itemName == "" {
		log.Println(err01.Error())
		http.Redirect(w, r, "/?err=ERR00007", 302)
		return
	}
	itemIcon, err02 := getValueByName(r, "item_icon")
	if err02 != nil {
		log.Println(err02.Error())
		http.Redirect(w, r, "/?err=ERR00007", 302)
		return
	}

	if itemIcon == "" {
		itemIcon = "default"
	}

	bsInstance := bslib.GetInstance()
	if bsInstance.IsLocked() == true {
		log.Println("Encryption is locked")
		http.Redirect(w, r, "/?err=ERR00007", 302)
		return
	}

	_, err := bsInstance.AddNewItem(bslib.JSONInputUpdateItem{ItemName: itemName, ItemIcon: itemIcon})
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/?err=ERR00007", 302)
		return
	}

	log.Println("Item successfully added")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
