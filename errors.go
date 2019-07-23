package main

var bsErrors map[string]string

func init() {
	bsErrors = make(map[string]string)
	bsErrors["ERR00001"] = "Password cannot be empty"
	bsErrors["ERR00002"] = "Password confirmation is not set"
	bsErrors["ERR00003"] = "Password and password confirmation are not equal"
	bsErrors["ERR00004"] = "Cypher is not set, registration is not possible"
	bsErrors["ERR00005"] = "Password is wrong, try again"
	bsErrors["ERR00006"] = "Cannot login, internal bykovstorage error"
	bsErrors["ERR00007"] = "Parameters are not set or empty"
}
