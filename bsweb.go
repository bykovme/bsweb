package main

import (
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if !checkLockAndStandardActions(w, r) {
		log.Println("Show main screen")
		processMain(w, r)
	}
}

func addItemPage(w http.ResponseWriter, r *http.Request) {
	if !checkLockAndStandardActions(w, r) {
		log.Println("Show add item screen")
		processAddItem(w, r)
	}
}

func handlerFavIcon(w http.ResponseWriter, r *http.Request) {
	// empty handler for favicon just to remove calling any other router for favicon
}

func main() {

	log.Println("BSWEB has started")

	bsInstance := bykovstorage.GetInstance()
	log.Print("BSWEB is initiating ")
	err := bsInstance.Open("bs.sqlite")
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}

	log.Println("BSWEB is setting handlers")
	http.Handle("/images/", http.FileServer(http.Dir("path/to/file")))
	http.HandleFunc("/favicon.ico", handlerFavIcon)
	http.HandleFunc("/additem", addItemPage) // adding item router
	http.HandleFunc("/", mainPage)           // set router
	//
	err = http.ListenAndServe(":29568", nil) // set listen port "bykov" :)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
