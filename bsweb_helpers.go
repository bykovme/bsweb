package main

import (
	"net/http"
)

func getAction(r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		return ""
	}
	action, err := getValueByName(r, "action")
	if err != nil {
		return ""
	}
	return action
}
