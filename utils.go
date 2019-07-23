package main

import (
	"errors"
	"net/http"
	"strings"
)

func getValueByName(r *http.Request, key string) (string, error) {

	for k, v := range r.Form {
		if k == key {
			return strings.Join(v, ""), nil
		}
	}
	return "", errors.New("Value not found")
}
