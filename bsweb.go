package main

import (
	"log"
	"net/http"

	"github.com/bykovme/bslib"
)

func addItemPage(w http.ResponseWriter, r *http.Request) {
	if !checkLockAndStandardActions(w, r) {
		log.Println("Show add item screen")
		processAddItem(w, r)
	}
}

func handlerFavIcon(_ http.ResponseWriter, _ *http.Request) {
	// empty handler for favicon just to remove calling any other router for favicon
}

func main() {

	err := initApp()
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}

	log.Println("BSWEB has started")

	bsInstance := bslib.GetInstance()
	log.Print("BSWEB is initiating ")
	err = bsInstance.Open(loadedConfig.DBPath)
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}

	log.Println("BSWEB is setting handlers")

	http.HandleFunc("/assets/", ServeAssetsHTTP)
	http.HandleFunc("/favicon.ico", handlerFavIcon)
	//http.HandleFunc("/additem", addItemPage) // adding item router
	http.HandleFunc("/", ServeHTTP) // set router
	//
	err = http.ListenAndServe(":"+loadedConfig.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func ServeAssetsHTTP(w http.ResponseWriter, r *http.Request) {
	shortPath := r.URL.Path
	fullPath := loadedConfig.AssetsPath + "/" + shortPath[1:]
	http.ServeFile(w, r, fullPath)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	if checkLockAndStandardActions(w, r) {
		return
	}
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {

	case "additem":
		addItemPage(w, r)
	case "items":
		viewItemPage(w, r)
	default:
		processMain(w, r)
	}
}

func checkLockAndStandardActions(w http.ResponseWriter, r *http.Request) bool {

	action := getAction(r)
	bsInstance := bslib.GetInstance()
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
